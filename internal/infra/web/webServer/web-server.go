package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type webHandlers struct {
	path    string
	handler http.HandlerFunc
	method  string
}

type WebServer struct {
	Router        chi.Router
	Handlers      []webHandlers
	webServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []webHandlers{},
		webServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc, method string) {
	fmt.Println("Adding handler", path, method)
	s.Handlers = append(s.Handlers, webHandlers{
		path:    path,
		handler: handler,
		method:  method,
	})
}

func (s *WebServer) Start() error {
	s.Router.Use(middleware.Logger)
	for _, h := range s.Handlers {
		s.Router.Method(h.method, h.path, h.handler)
	}

	return http.ListenAndServe(s.webServerPort, s.Router)
}
