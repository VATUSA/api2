package log

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func New(format string) {
	Logger = logrus.New()
	if format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:          true,
			PadLevelText:           true,
			DisableLevelTruncation: true,
		})
	}
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
