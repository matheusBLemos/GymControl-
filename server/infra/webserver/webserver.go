package webserver

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WebServer struct {
	App           *fiber.App
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		App:           fiber.New(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {
	s.App.Handler()
	s.App.Listen(s.WebServerPort)
}
