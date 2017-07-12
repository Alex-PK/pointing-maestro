package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"sync"
	"path/filepath"
	"html/template"
	"log"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("tpl", t.filename)))
	})
	t.templ.Execute(res, nil)
}

type room struct {
	name string
}

type roomHandler struct {
	rooms *map[string]*room
}

func (self *roomHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	roomdata, ok := (*self.rooms)[id]
	if !ok {
		roomdata = &room{ name: id }
		(*self.rooms)[id] = roomdata
	}

	res.Write([]byte(roomdata.name))
}

func main() {
	rooms := make(map[string]*room)
	router := mux.NewRouter()
	router.Handle("/", &templateHandler{filename: "home.html"});
	router.Handle("/room/{id:[a-zA-z0-9_-]+}", &roomHandler{rooms: &rooms});

	http.Handle("/", router)

	if err:= http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Cannot run server:", err)
	}
}
