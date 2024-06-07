package server

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/moura95/transferapi/config"
	"github.com/moura95/transferapi/internal/api"
	"github.com/moura95/transferapi/internal/middleware"

	"go.uber.org/zap"
)

type Server struct {
	store  *sqlx.DB
	router *gin.Engine
	config *config.Config
	logger *zap.SugaredLogger
}

func NewServer(cfg config.Config, store *sqlx.DB, log *zap.SugaredLogger) *Server {

	server := &Server{
		store:  store,
		config: &cfg,
		logger: log,
	}
	var router *gin.Engine

	router = gin.Default()

	// Middleware Rate Limiter
	router.Use(middleware.RateLimitMiddleware())
	// Gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Middleware Cors
	router.Use(middleware.CORSMiddleware())

	// Init all Routers
	api.CreateRoutesV1(store, server.config, router, log)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func RunGinServer(cfg config.Config, store *sqlx.DB, log *zap.SugaredLogger) {
	server := NewServer(cfg, store, log)

	_ = server.Start(cfg.HTTPServerAddress)
}
