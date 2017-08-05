package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type homeHandler struct {
	tpls *templates
}

func (self *homeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	self.tpls.render(res, "home.html", nil)
}

func main() {
	rooms := newRooms()
	tpls := newTemplates()

	router := mux.NewRouter()

	router.Handle("/", &homeHandler{tpls})
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	router.Handle("/room/{id:[a-zA-z0-9_-]+}", &roomHandler{rooms: rooms, tpls: tpls})
	router.Handle("/msg/{id:[a-zA-z0-9_-]+}", &msgHandler{rooms: rooms})

	http.Handle("/", router)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Cannot run server:", err)
	}
}
