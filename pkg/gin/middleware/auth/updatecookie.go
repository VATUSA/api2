package auth

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vatusa/api2/pkg/vatlog"
)

func UpdateCookie(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("t", time.Now().String())

	err := session.Save()
	if err != nil {
		vatlog.Logger.WithField("component", "middleware/UpdateCookie").Errorf("Error saving cookie: %s", err.Error())
	}

	c.Next()
}
