package seed

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vatusa/api2/internal/config"
	"github.com/vatusa/api2/internal/seed"
	"github.com/vatusa/api2/pkg/database"
	"github.com/vatusa/api2/pkg/log"
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

func Run(c *cli.Context) error {
	if !seed.IsValidSeedType(c.String("seed-type")) {
		return errors.New("invalid seed type: " + c.String("seed-type"))
	}

	log.Logger.Info("Loading configuration")
	cfg, err := config.Load(c.String("config"))
	if err != nil {
		log.Logger.Error("Error loading configuration: " + err.Error())
		return err
	}
	Config = cfg

	log.Logger.Info("Connecting to Database")
	opts := database.DBOptions{
		Host:     Config.Database.Host,
		Port:     Config.Database.Port,
		User:     Config.Database.User,
		Password: Config.Database.Password,
		Database: Config.Database.Database,
		Driver:   Config.Database.Driver,
		Options:  "sslmode=disable TimeZone=UTC",
		Logger:   log.Logger,
	}
	err = database.Connect(opts)
	if err != nil {
		log.Logger.Error("Error connecting to database: " + err.Error())
		return err
	}
	log.Logger.Info("Connected to database")

	log.Logger.Info("Loading seed")
	file, err := os.Open(c.String("file"))
	if err != nil {
		log.Logger.Error("Error opening file: " + err.Error())
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Logger.Error("Error reading file: " + err.Error())
		return err
	}

	s, err := seed.BuildSeed(c.String("seed-type"), data)
	if err != nil {
		log.Logger.Error("Error building seed: " + err.Error())
		return err
	}

	log.Logger.Info("Seeding database")

	switch s.Kind {
	case "rating":
		ratings := seed.BuildRatings(s.Values)
		if database.DB.Create(ratings).Error != nil {
			//			log.Logger.Error("Error seeding ratings: " + err.Error())
			return err
		}
	case "facility":
		facilities := seed.BuildFacilities(s.Values)
		if database.DB.Create(facilities).Error != nil {
			//			log.Logger.Error("Error seeding facilities: " + err.Error())
			return err
		}
	}

	log.Logger.Info("Seeding complete")

	return nil
}
