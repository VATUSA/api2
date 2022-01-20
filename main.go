package main

import (
	"errors"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vatusa/api2/cmd/migrate"
	"github.com/vatusa/api2/cmd/seed"
	"github.com/vatusa/api2/cmd/server"
	"github.com/vatusa/api2/pkg/log"
)

func main() {
	app := &cli.App{
		Name:                 "api",
		Usage:                "VATUSA API",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			server.Command(),
			migrate.Command(),
			seed.Command(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "log-level",
				Value:   "info",
				Aliases: []string{"l"},
				Usage:   "Log level (accepted values: trace, debug, info, warn, error, fatal, panic)",
			},
			&cli.StringFlag{
				Name:  "log-format",
				Value: "text",
				Usage: "Log format (accepted values: text, json)",
			},
		},
		Before: func(c *cli.Context) error {
			format := c.String("log-format")
			if !log.IsValidFormat(format) {
				return errors.New("invalid log format")
			}
			log.New(format)

			if log.IsValidLogLevel(c.String("log-level")) {
				l, _ := log.ParseLogLevel(c.String("log-level"))
				log.Logger.SetLevel(l)
			} else {
				return errors.New("invalid log level")
			}

			log.Logger.Info("Starting VATUSA API")

			return nil
		},
	}

	app.Run(os.Args)
}
