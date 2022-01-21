package migrate

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/vatusa/api2/cmd"
	"github.com/vatusa/api2/internal/config"
	"github.com/vatusa/api2/pkg/database"
	"github.com/vatusa/api2/pkg/database/models"
	"github.com/vatusa/api2/pkg/vatlog"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "Run the Migrations",
		Flags: []cli.Flag{
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
var log *logrus.Entry

func Run(c *cli.Context) error {
	log = vatlog.Logger.WithField("component", "cmd/migrate")
	err := cmd.LoadConfig(c.String("config"))
	if err != nil {
		return err
	}

	err = cmd.BuildDatabase()
	if err != nil {
		return err
	}

	log.Info("Migrating database")
	database.DB.AutoMigrate(
		&models.Facility{},
		&models.Rating{},
		&models.User{},
	)
	log.Info("Done")

	return nil
}
