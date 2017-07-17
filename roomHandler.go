package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type roomHandler struct {
	rooms *rooms
	tpls *templates
}

func (self *roomHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	room := self.rooms.get(id)

	data := struct {
		Room string
	}{
		Room: room.name,
	}

	self.tpls.render(res, "room.html", data)
}
