package server

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/vatusa/api2/internal/config"
	"github.com/vatusa/api2/pkg/gin/middleware/auth"
	"github.com/vatusa/api2/pkg/gin/middleware/logger"
	"github.com/vatusa/api2/pkg/vatlog"
)

type Server struct {
	engine *gin.Engine
}

var log = vatlog.Logger

func Run() {
	// Build notify context, we'll use this to handle graceful shutdowns
	log.WithField("component", "internal/server").Info("Building notify context")
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.WithField("component", "internal/server").Info("Configuring Gin webserver")
	server := NewServer()
}

func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)

	server := Server{}
	engine := gin.New()

	log.WithField("component", "internal/server").Debug("Loading recovery middleware")
	engine.Use(gin.Recovery())

	log.WithField("component", "internal/server").Debug("Loading logging middleware")
	engine.Use(logger.Logger)

	log.WithField("component", "internal/server").Debug("Loading CORS middleware")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
		"X-Requested-With",
	}
	corsConfig.AllowCredentials = true
	corsConfig.AllowWildcard = true
	engine.Use(cors.New(corsConfig))

	log.WithField("component", "internal/server").Debug("Loading session and auth middlewares")
	store := buildCookieStore(config.Cfg.Session.Cookie)
	engine.Use(sessions.Sessions(config.Cfg.Session.Cookie.Name, store))
	engine.Use(auth.UpdateCookie)
}

func buildCookieStore(c config.ConfigSessionCookie) cookie.Store {
	store := cookie.NewStore([]byte(c.Secret))
	store.Options(sessions.Options{
		Path:     c.Path,
		Domain:   c.Domain,
		MaxAge:   c.MaxAge,
		HttpOnly: true,
	})

	return store
}
