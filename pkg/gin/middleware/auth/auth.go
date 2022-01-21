package auth

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vatusa/api2/pkg/database/models"
	"github.com/vatusa/api2/pkg/gin/response"
	"github.com/vatusa/api2/pkg/jwt"
)

func Auth(c *gin.Context) {
	// JWT Check
	authHeader := c.GetHeader("authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString := authHeader[7:]
		token, err := jwt.ParseToken(tokenString)
		if err == nil {
			cid := token.Subject()
			user, err := models.FindUserByCID(cid)
			if err == nil {
				c.Set("x-guest", false)
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

	// API Key (X-API-Key header)
	if c.GetHeader("x-api-key") != "" {
		facility, err := models.FindFacilityByAPIKey(c.GetHeader("x-api-key"))
		if err == nil {
			c.Set("x-guest", false)
			c.Set("x-facility-d", facility.IATA)
			c.Set("x-facility", facility)
			c.Set("x-auth-type", "api-key")
			c.Next()
			return
		} else {
			response.RespondError(c, http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}
	}

	// Cookie check
	session := sessions.Default(c)
	cid := session.Get("cid")
	if cid == nil {
		c.Set("x-guest", true)
		c.Next()
		return
	}

	user, err := models.FindUserByCID(cid.(string))
	if err == nil {
		c.Set("x-guest", false)
		c.Set("x-cid", cid)
		c.Set("x-user", user)
		c.Set("x-auth-type", "cookie")
		c.Next()
		return
	}

	// We should only get here if they have a cookie with a cid but no user
	// So reset cookie and leave them as a guest
	session.Delete("cid")
	session.Save()
	c.Set("x-guest", true)
	c.Next()
}

func NotGuest(c *gin.Context) {
	if c.GetBool("x-guest") {
		response.RespondError(c, http.StatusForbidden, "Forbidden")
		c.Abort()
		return
	}
	c.Next()
}

func NotGuestOrAPIKey(c *gin.Context) {
	if c.GetBool("x-guest") || c.GetString("x-auth-type") == "api-key" {
		response.RespondError(c, http.StatusForbidden, "Forbidden")
		c.Abort()
		return
	}
	c.Next()
}
