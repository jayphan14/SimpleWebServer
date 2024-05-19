package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Engine *gin.Engine
}

func NewServer(port string) Server {
	s := Server{
		Port:   port,
		Engine: gin.Default(),
	}
	return s
}

func (s *Server) StartServer() {
	s.Engine.Run(fmt.Sprintf(":%s", s.Port))
}

func (s *Server) AddEndpoint(httpMethod string, endpoint string, handler func(c *gin.Context)) {
	switch httpMethod {
	case "GET":
		s.Engine.GET(endpoint, handler)
	case "POST":
		s.Engine.POST(endpoint, handler)
	case "PUT":
		s.Engine.PUT(endpoint, handler)
	case "PATCH":
		s.Engine.PATCH(endpoint, handler)
	case "DELETE":
		s.Engine.DELETE(endpoint, handler)
	case "OPTIONS":
		s.Engine.OPTIONS(endpoint, handler)
	default:
		fmt.Printf("Unsupported HTTP method: %s\n", httpMethod)
	}
}
