package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

type roomHandler struct {
	rooms *map[string]*room
	tpls *templates
}

func (self *roomHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	room, ok := (*self.rooms)[id]
	if !ok {
		room = newRoom(id)
		(*self.rooms)[id] = room
		go room.run()
		log.Printf("Created new room %s\n", id)
	}

	data := struct {
		Room string
	}{
		Room: room.name,
	}

	self.tpls.render(res, "room.html", data)
}
