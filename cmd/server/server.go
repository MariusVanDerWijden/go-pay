package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
)

type Server struct {
	channels map[[32]byte]*Channel
	database *DB
	lock     sync.RWMutex
}

func NewServer(database string) (*Server, error) {
	db, err := NewDB(database)
	if err != nil {
		return nil, err
	}
	channels, err := db.Load()
	if err != nil {
		return nil, err
	}
	return &Server{
		database: db,
		channels: channels,
	}, nil
}

func (s *Server) handleChannels(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Serving Channels request")
	s.lock.RLock()
	defer s.lock.RUnlock()

	ids := Channels{}
	for _, c := range s.channels {
		ids.ID = append(ids.ID, c.ID) // TODO cache this
	}
	resp, err := json.Marshal(ids)
	if err != nil {
		rw.WriteHeader(500)
		return
	}
	rw.Write(resp)
}

func (s *Server) getFullChannel(rw http.ResponseWriter, req *http.Request) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	reqID := mux.Vars(req)["id"]
	var id [32]byte
	copy(id[:], common.FromHex(reqID))
	channel, ok := s.channels[id]
	if !ok {
		rw.WriteHeader(404)
		return
	}
	fmt.Printf("Serving Channel request for channel %x\n", id)
	resp, err := json.Marshal(channel)
	if err != nil {
		rw.WriteHeader(500)
		return
	}
	rw.Write(resp)
}

func (s *Server) update(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Serving update channel request")
	s.lock.Lock()
	defer s.lock.Unlock()

	decoder := json.NewDecoder(req.Body)
	var request UpdateChannel
	if err := decoder.Decode(&request); err != nil {
		rw.WriteHeader(500)
		return
	}
	if _, ok := s.channels[request.ID]; !ok {
		currentRound := State{
			ValueA: request.ValueA,
			ValueB: request.ValueB,
		}
		channel := &Channel{
			ID:     request.ID,
			AddrA:  request.AddrA,
			AddrB:  request.AddrB,
			States: make(map[*big.Int]State),
		}
		channel.States[request.Round] = currentRound
		channel.CurrentRound = request.Round
		s.channels[request.ID] = channel
		s.database.Store(channel)
		fmt.Printf("Created channel with ID: %x", request.ID)
		rw.WriteHeader(200)
		return
	}
	fmt.Printf("Channel with ID %x already exists, updating\n", request.ID)
	if err := s.updateChannel(request); err != nil {
		rw.WriteHeader(500)
		return
	}
	rw.WriteHeader(200)
}

func (s *Server) updateChannel(request UpdateChannel) error {
	channel := s.channels[request.ID]
	if channel.AddrA != request.AddrA || channel.AddrB != request.AddrB {
		return errors.New("invalid update request, mismatching addresses")
	}
	currentState := channel.States[channel.CurrentRound]
	if new(big.Int).Add(currentState.ValueA, currentState.ValueB).Cmp(new(big.Int).Add(request.ValueA, request.ValueB)) != 0 {
		return errors.New("invalid update request, mismatching values")
	}

	if _, ok := channel.States[request.Round]; ok {
		return errors.New("state already stored")
	}
	channel.States[request.Round] = State{
		ValueA: request.ValueA,
		ValueB: request.ValueB,
	}
	if channel.CurrentRound.Cmp(request.Round) == -1 {
		channel.CurrentRound = request.Round
	}
	s.database.Store(channel)
	return nil
}

func (s *Server) getLightChannel(rw http.ResponseWriter, req *http.Request) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	reqID := mux.Vars(req)["id"]
	var id [32]byte
	copy(id[:], common.FromHex(reqID))
	channel, ok := s.channels[id]
	if !ok {
		rw.WriteHeader(404)
		return
	}

	round := channel.States[channel.CurrentRound]

	lightChannel := LightChannel{
		ID:           channel.ID,
		AddrA:        channel.AddrA,
		AddrB:        channel.AddrB,
		ValueA:       round.ValueA,
		ValueB:       round.ValueB,
		CurrentState: channel.CurrentRound,
	}
	fmt.Printf("Serving light channel request for channel %x\n", id)
	resp, err := json.Marshal(lightChannel)
	if err != nil {
		rw.WriteHeader(500)
		return
	}
	rw.Write(resp)
}

func (s *Server) getChannelState(rw http.ResponseWriter, req *http.Request) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	reqID := mux.Vars(req)["id"]
	var id [32]byte
	copy(id[:], common.FromHex(reqID))
	reqRound := mux.Vars(req)["round"]
	round, ok := new(big.Int).SetString(reqRound, 0)
	if !ok {
		rw.WriteHeader(400)
		return
	}

	channel, ok := s.channels[id]
	if !ok {
		rw.WriteHeader(404)
		return
	}
	fmt.Printf("Serving channel state request for channel %x\n", id)
	state, ok := channel.States[round]
	if !ok {
		rw.WriteHeader(404)
		return
	}

	resp, err := json.Marshal(state)
	if err != nil {
		rw.WriteHeader(500)
		return
	}
	rw.Write(resp)
}
