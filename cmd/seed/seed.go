package seed

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/vatusa/api2/cmd"
	"github.com/vatusa/api2/internal/config"
	"github.com/vatusa/api2/internal/seed"
	"github.com/vatusa/api2/pkg/database"
	"github.com/vatusa/api2/pkg/vatlog"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "seed",
		Usage: "Seed certain needed fields",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Usage:   "Path to the configuration file. Default: config.yaml",
				Value:   "config.yaml",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:     "file",
				Usage:    "Path to the file with the data.",
				Aliases:  []string{"f"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "seed-type",
				Value:   "yaml",
				Usage:   "Type of the seed file. Accepted values: yaml, json. Default: yaml",
				Aliases: []string{"t"},
			},
		},
		Action: Run,
	}
}

var Config *config.Config
var log *logrus.Entry

func Run(c *cli.Context) error {
	log = vatlog.Logger.WithField("component", "cmd/seed")

	if !seed.IsValidSeedType(c.String("seed-type")) {
		return errors.New("invalid seed type: " + c.String("seed-type"))
	}

	err := cmd.LoadConfig(c.String("config"))
	if err != nil {
		return err
	}

	err = cmd.BuildDatabase()
	if err != nil {
		return err
	}

	log.Info("Loading seed")
	file, err := os.Open(c.String("file"))
	if err != nil {
		log.Error("Error opening file: " + err.Error())
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error("Error reading file: " + err.Error())
		return err
	}

	s, err := seed.BuildSeed(c.String("seed-type"), data)
	if err != nil {
		log.Error("Error building seed: " + err.Error())
		return err
	}

	log.Info("Seeding database")

	switch s.Kind {
	case "rating":
		ratings := seed.BuildRatings(s.Values)
		if database.DB.Create(ratings).Error != nil {
			log.Error("Error seeding ratings: " + err.Error())
			return err
		}
	case "facility":
		facilities := seed.BuildFacilities(s.Values)
		if database.DB.Create(facilities).Error != nil {
			log.Error("Error seeding facilities: " + err.Error())
			return err
		}
	}

	log.Info("Seeding complete")

	return nil
}
