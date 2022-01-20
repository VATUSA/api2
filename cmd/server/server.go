package server

import (
	"github.com/urfave/cli/v2"
	"github.com/vatusa/api2/internal/config"
	"github.com/vatusa/api2/pkg/log"
)

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

var Config *config.Config

func Run(c *cli.Context) error {
	log.Logger.Info("Loading configuration")
	cfg, err := config.Load(c.String("config"))
	if err != nil {
		log.Logger.Error("Error loading configuration: %s", err)
		return err
	}
	Config = cfg

	log.Logger.Info("Connecting to Database")
	/* 	opts := database.DBOptions{
		Host:     Config.Database.Host,
		Port:     Config.Database.Port,
		User:     Config.Database.User,
		Password: Config.Database.Password,
		Database: Config.Database.Database,
		Driver:   Config.Database.Driver,
		Options:  "sslmode=disable TimeZone=UTC",
	} */

	return nil
}
