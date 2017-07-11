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

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("tpl", t.filename)))
	})
	t.templ.Execute(w, nil)
}

type roomHandler struct {

}

func (room *roomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte(id))
}

func main() {
	router := mux.NewRouter()
	router.Handle("/", &templateHandler{filename: "home.html"});
	router.Handle("/{id:[0-9]+}", &roomHandler{});

	http.Handle("/", router)

	if err:= http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Cannot run server:", err)
	}
}
