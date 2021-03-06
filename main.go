package main

import (
	"errors"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vatusa/api2/cmd/migrate"
	"github.com/vatusa/api2/cmd/seed"
	"github.com/vatusa/api2/cmd/server"
	"github.com/vatusa/api2/pkg/vatlog"
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
			if !vatlog.IsValidFormat(format) {
				return errors.New("invalid log format")
			}
			vatlog.New(format)

			if vatlog.IsValidLogLevel(c.String("log-level")) {
				l, _ := vatlog.ParseLogLevel(c.String("log-level"))
				vatlog.Logger.SetLevel(l)
			} else {
				return errors.New("invalid log level")
			}

			vatlog.Logger.Info("Starting VATUSA API")

			return nil
		},
	}

	app.Run(os.Args)
}

// @title VATUSA API
// @version 3.0
// @description VATUSA APIv2

// @contact.name Daniel Hawton
// @contact.email daniel@hawton.org

// @license.name BSD
// @license.url https://github.com/VATUSA/api2/blob/main/LICENSE

// @host api.vatusa.net
// @BasePath /v3
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth
// @description JWT (header: Authorization: Bearer (token)), APIKey (header: X-API-Key: (apikey)), or Session Cookie
