package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vatusa/api2/pkg/gin/response"
)

func CreateRoutes(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		response.RespondMessage(c, http.StatusOK, "PONG")
	})
}
