package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vatusa/api2/pkg/database/models"
	"github.com/vatusa/api2/pkg/gin/response"
	"github.com/vatusa/api2/pkg/jwt"
)

func Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString := authHeader[7:]
		token, err := jwt.ParseToken(tokenString)
		if err == nil {
			cid := token.Subject()
			user, err := models.FindUserByCID(cid)
			if err == nil {
				c.Set("x-cid", cid)
				c.Set("x-user", user)
				c.Set("x-auth-type", "jwt")
				c.Next()
				return
			} else {
				response.RespondError(c, http.StatusForbidden, "Forbidden")
				c.Abort()
				return
			}
		}
	}

	if c.GetHeader("x-api-key") != "" {

	}

	c.Next()
}
