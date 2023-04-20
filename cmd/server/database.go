package main

import (
	"encoding/json"
	"sync"

	"github.com/cockroachdb/pebble"
)

type DB struct {
	db   *pebble.DB
	lock sync.Mutex
}

var (
	channelPrefix = []byte("c")
)

func NewDB(file string) (*DB, error) {
	db, err := pebble.Open("/tmp/foo.db", nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &DB{
		db: db,
	}, nil
}

func (db *DB) Load() (map[[32]byte]*Channel, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	iter := db.db.NewIter(nil)
	channels := make(map[[32]byte]*Channel, 0)
	for ok := iter.First(); ok; ok = iter.Next() {
		val := iter.Value()
		var ch Channel
		if err := json.Unmarshal(val, &ch); err != nil {
			return nil, err
		}
		channels[ch.ID] = &ch
	}
	return channels, nil
}

func (db *DB) Store(ch *Channel) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	key := append(channelPrefix, ch.ID[:]...)
	val, err := json.Marshal(ch)
	if err != nil {
		return err
	}
	return db.db.Set(key, val, pebble.Sync)
}
