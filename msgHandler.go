package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

type msgHandler struct {
	rooms *rooms
}

func (self *msgHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	room := self.rooms.get(id)

	socket, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   room,
		name:   "dummy",
	}

	room.join <- client

	defer func() { room.leave <- client }()

	go client.write()
	client.read()
}
