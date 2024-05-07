package webserver

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WebServer struct {
	App 		*fiber.App
	Handlers      map[string]fiber.Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		App:        fiber.New(),
		Handlers:      make(map[string]fiber.Handler),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler fiber.Handler) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {
	for path, handler := range s.Handlers {
		s.App.Add(http.MethodGet, path, handler)
	}
	s.App.Listen(s.WebServerPort)
}
