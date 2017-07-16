package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"sync"
	"path/filepath"
	"html/template"
	"log"
)

type templates struct {
	lock sync.RWMutex
	tpls map[string]*template.Template
}

func newTemplates() *templates {
	return &templates{tpls: make(map[string]*template.Template)}
}

func (self *templates) render(dest http.ResponseWriter, name string, data interface{}) {
	self.lock.RLock();
	tpl, ok := self.tpls[name]
	if !ok 	{
		self.lock.RUnlock()
		self.lock.Lock()
		tpl = template.Must(template.ParseFiles(filepath.Join("tpl", name)))
		self.tpls[name] = tpl
		self.lock.Unlock()
	}
	tpl.Execute(dest, data)
}

type homeHandler struct {
	tpls *templates
}

func (self *homeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	self.tpls.render(res, "home.html", nil)
}

func main() {
	rooms := make(map[string]*room)
	tpls := newTemplates()

	router := mux.NewRouter()

	router.Handle("/", &homeHandler{tpls});
	router.Handle("/room/{id:[a-zA-z0-9_-]+}", &roomHandler{rooms: &rooms, tpls: tpls});
	router.Handle("/msg/{id:[a-zA-z0-9_-]+}", &msgHandler{rooms: &rooms});

	http.Handle("/", router)

	if err:= http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Cannot run server:", err)
	}
}
