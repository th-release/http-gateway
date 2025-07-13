package api

import (
	"cth.release/http-gateway/api/gateway"
	"cth.release/http-gateway/utils"
	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	App    *fiber.App
	Config utils.Config
}

func InitServer(config *utils.Config) *ServerConfig {
	app := fiber.New()

	if config == nil {
		return nil
	}

	server := &ServerConfig{
		App:    app,
		Config: *config,
	}

	server.setupRoutes()
	return server
}

func (s *ServerConfig) setupRoutes() {
	apiGroup := s.App.Group("/api", ApiMiddleware)

	apiGroup.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(utils.BasicResponse{
			Success: true,
			Message: "",
			Data:    nil,
		})
	})

	apiGroup.Get("/", gateway.Get)
	apiGroup.Post("/", gateway.Post)
	apiGroup.Put("/", gateway.Put)
	apiGroup.Delete("/", gateway.Delete)
	apiGroup.Patch("/", gateway.Patch)
}
