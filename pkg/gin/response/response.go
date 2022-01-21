package response

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type R struct {
	XMLName xml.Name    `xml:"response" json:"-" yaml:"-"`
	Status  string      `xml:"status" json:"status" yaml:"status"`
	Data    interface{} `xml:"data" json:"data" yaml:"data"`
}

func RespondMessage(c *gin.Context, status int, message string) {
	Respond(c, status, struct {
		Message string `json:"message" yaml:"message" xml:"message"`
	}{message})
}

func RespondBlank(c *gin.Context, status int) {
	c.Status(status)
	c.Abort()
}

func RespondError(c *gin.Context, status int, message string) {
	RespondMessage(c, status, message)
}

func Respond(c *gin.Context, status int, data interface{}) {
	ret := R{}
	ret.Status = http.StatusText(status)
	ret.Data = data

	if acceptYaml(c.GetHeader("Accept")) {
		c.YAML(status, ret)
	} else if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, ret)
	} else {
		c.JSON(status, ret)
	}
}

func acceptYaml(accept string) bool {
	return accept == "application/yaml" || accept == "text/yaml" || accept == "application/x-yaml"
}
