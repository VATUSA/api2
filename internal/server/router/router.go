package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vatusa/api2/internal/api"
)

func CreateRoutes(engine *gin.Engine) {
	engine.GET("/ping", api.GetPing)

	v3 := engine.Group("/v3")
	{
		v3.GET("/ping", api.GetPing)
	}
}
