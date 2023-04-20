package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if len(os.Args) < 2 {
		panic("invalid amount of args, need 2")
	}

	server := Server{}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/channels", server.handleChannels).
		Methods("GET")
	router.HandleFunc("/channels/{id}", server.getLightChannel).
		Methods("GET")
	router.HandleFunc("/channels/{id}/full", server.getFullChannel).
		Methods("GET")
	router.HandleFunc("/channels/{id}/states/{round}", server.getChannelState).
		Methods("GET")
	router.HandleFunc("/channels/new", server.update).
		Methods("POST")
	if err := http.ListenAndServe(":2016", router); err != nil {
		log.Fatal(err)
	}
}
