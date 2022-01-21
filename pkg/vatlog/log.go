package vatlog

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger
var Format string

func New(format string) {
	Logger = logrus.New()
	if format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			TimestampFormat: "2006-01-02T15:04:05Z07:00",
			FieldsOrder:     []string{"component", "category"},
			ShowFullLevel:   true,
		})
	}
	Format = format
}

func IsValidFormat(format string) bool {
	return format == "text" || format == "json"
}

func IsValidLogLevel(level string) bool {
	_, err := ParseLogLevel(level)
	return err == nil
}

func ParseLogLevel(level string) (logrus.Level, error) {
	return logrus.ParseLevel(level)
}
