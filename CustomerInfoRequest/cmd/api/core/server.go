package core

import (
	"fmt"
	"log/slog"

	"hexagonalExample/cmd/api/configurations"
	"hexagonalExample/cmd/api/handlers/health"

	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg      configurations.Config
	logger   *slog.Logger
	httpAddr string
	engine   *gin.Engine
}

func New(cfg configurations.Config, logger *slog.Logger) Server {
	server := Server{
		cfg:      cfg,
		logger:   logger,
		httpAddr: fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		engine:   gin.Default(),
	}
	server.registerRoutes(logger)
	return server
}

func (s *Server) Run() error {
	log.Println("Server runing on port: ", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}
func (s *Server) registerRoutes(logger *slog.Logger) {

	s.engine.GET("/health", health.CheckHandler(logger))

	s.engine.NoMethod(func(c *gin.Context) {
		logger.Error("Method not supported",
			"method", c.Request.Method,
			"error", "method not allowed",
		)
		c.JSON(405, gin.H{
			"error":   "method not allowed",
			"message": fmt.Sprintf("The method %s is not supported for this route", c.Request.Method),
		})
	})

	s.engine.NoRoute(func(c *gin.Context) {
		logger.Error("Endpoint not found",
			"method", c.Request.Method,
			"error", "endpoint not found",
		)
		c.JSON(404, gin.H{
			"error":   "endpoint not found",
			"message": fmt.Sprintf("The route %s is not registered", c.Request.URL.Path),
		})
	})

}
