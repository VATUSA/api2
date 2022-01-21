package server

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/vatusa/api2/cmd"
	"github.com/vatusa/api2/cmd/migrate"
	"github.com/vatusa/api2/internal/config"
	"github.com/vatusa/api2/pkg/vatlog"
)

var log *logrus.Entry

func Command() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Run the VATUSA API server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Value:   3000,
				Aliases: []string{"p"},
				Usage:   "Port to listen on",
			},
			&cli.StringFlag{
				Name:    "config",
				Usage:   "Path to the configuration file. Default: config.yaml",
				Value:   "config.yaml",
				Aliases: []string{"c"},
			},
		},
		Action: Run,
	}
}

func Run(c *cli.Context) error {
	log = vatlog.Logger.WithField("component", "cmd/server")
	err := cmd.LoadConfig(c.String("config"))
	if err != nil {
		return err
	}

	err = cmd.BuildDatabase()
	if err != nil {
		return err
	}

	log.Debug("Checking if we should automigrate")
	if config.Cfg.Database.AutoMigrate {
		log.Info("Migrating database")
		err := migrate.Run(c)
		if err != nil {
			return err
		}
	}

	return nil
}
