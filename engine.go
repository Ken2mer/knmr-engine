package engine

import (
	"html/template"
	"log"
	"net/http"
)

// Engine is the framework's instance
type Engine struct {
	FuncMap template.FuncMap
}

// New returns a new blank Engine instance
func New() *Engine {
	engine := &Engine{
		FuncMap: template.FuncMap{},
	}
	return engine
}

// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
func (engine *Engine) Run(addr string) error {
	log.Printf("Listening and serving HTTP on %s\n", addr)
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO
}
