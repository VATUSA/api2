package server

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/vatusa/api2/internal/config"
	"github.com/vatusa/api2/internal/server/router"
	"github.com/vatusa/api2/pkg/gin/middleware/auth"
	"github.com/vatusa/api2/pkg/gin/middleware/logger"
	"github.com/vatusa/api2/pkg/vatlog"
)

type Server struct {
	Engine *gin.Engine
}

var log *logrus.Logger

func Run() error {
	log = vatlog.Logger

	// Build notify context, we'll use this to handle graceful shutdowns
	log.WithField("component", "internal/server").Info("Building notify context")
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.WithField("component", "internal/server").Info("Configuring Gin webserver")
	server := NewServer()

	log.WithField("component", "internal/server").Info("Generating Routes")
	router.CreateRoutes(server.Engine)

	log.WithField("component", "internal/server").Infof("Starting Gin webserver on %s:%s", config.Cfg.Server.Host, config.Cfg.Server.Port)
	srv := &http.Server{
		Addr:    config.Cfg.Server.Host + ":" + config.Cfg.Server.Port,
		Handler: server.Engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithField("component", "internal/server").Errorf("Gin webserver failed: %s", err)
		}
	}()

	// Catch interrupt signals and gracefully shutdown
	<-ctx.Done()
	log.WithField("component", "internal/server").Info("Shutting down server")
	stop()                                                                   // Stop context
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Create context with timeout, to force shutdown if not completed in 15 seconds
	defer cancel()                                                           // Cancel context
	if err := srv.Shutdown(ctx); err != nil {
		log.WithField("component", "internal/server").Errorf("Gin webserver failed to shutdown: %s. Will force.", err)
	}
	log.WithField("component", "internal/server").Info("Gin webserver shutdown complete")

	return nil
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

	// Load up API Docs endpoints
	engine.LoadHTMLGlob("static/*.html")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Engine = engine
	return &server
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
