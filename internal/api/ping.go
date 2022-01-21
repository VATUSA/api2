package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vatusa/api2/pkg/gin/response"
)

// Healthcheck endpoint
// @Summary Ping, healthcheck endpoint
// @Description Ping, healthcheck endpoint
// @Tags misc
// @Accept  json,xml,application/x-yaml
// @Produce json,xml,application/x-yaml
// @Success 200 {object} response.R
// @Router /ping [get]
func GetPing(c *gin.Context) {
	response.RespondMessage(c, http.StatusOK, "PONG")
}
