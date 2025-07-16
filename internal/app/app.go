package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/savanyv/e-commerce/config"
	"github.com/savanyv/e-commerce/internal/database"
	"github.com/savanyv/e-commerce/internal/delivery/routes"
)

type Server struct {
	server *echo.Echo
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		server: echo.New(),
		config: cfg,
	}
}

func (s *Server) RunServer() error {
	// Serup Database Connection
	_, err := database.ConnectPostgres(s.config)
	if err != nil {
		log.Println("error connecting to postgres", err)
		return err
	}

	// Setup Routes
	if err := routes.SetupRoutes(s.server); err != nil {
		log.Println("error setting up routes", err)
		return err
	}

	// Start Server
	if err := s.server.Start(":" + s.config.PortServer); err != nil {
		log.Fatal("error starting server", err)
		return err
	}

	log.Println("Server started on port", s.config.PortServer)
	return nil
}
