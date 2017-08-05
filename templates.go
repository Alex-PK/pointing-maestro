package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type templates struct {
	lock sync.RWMutex
	tpls map[string]*template.Template
}

func newTemplates() *templates {
	return &templates{tpls: make(map[string]*template.Template)}
}

func (self *templates) render(dest http.ResponseWriter, name string, data interface{}) {
	self.lock.RLock()
	tpl, ok := self.tpls[name]
	if !ok {
		self.lock.RUnlock()
		self.lock.Lock()
		tpl = template.Must(template.ParseFiles(filepath.Join("tpl", name)))
		self.tpls[name] = tpl
		self.lock.Unlock()
	}
	tpl.Execute(dest, data)
}
