// This file should provide common functions used by multiple subcommands, but
// may not necessarily be all subcommands (so it wouldn't be part of main.go)

package cmd

import (
	"github.com/vatusa/api2/internal/config"
	"github.com/vatusa/api2/pkg/database"
	"github.com/vatusa/api2/pkg/vatlog"
)

func LoadConfig(filename string) error {
	if config.Cfg != nil {
		return nil
	}

	vatlog.Logger.WithField("component", "cmd").Info("Loading configuration")
	cfg, err := config.Load(filename)
	if err != nil {
		vatlog.Logger.WithField("component", "cmd").Error("Error loading configuration: " + err.Error())
		return err
	}

	config.Cfg = cfg

	return nil
}

func BuildDatabase() error {
	if database.DB != nil {
		return nil
	}

	vatlog.Logger.WithField("component", "cmd").Info("Connecting to Database")
	opts := database.DBOptions{
		Host:     config.Cfg.Database.Host,
		Port:     config.Cfg.Database.Port,
		User:     config.Cfg.Database.User,
		Password: config.Cfg.Database.Password,
		Database: config.Cfg.Database.Database,
		Driver:   config.Cfg.Database.Driver,
		Options:  "sslmode=disable TimeZone=UTC",
		Logger:   vatlog.Logger,
	}
	err := database.Connect(opts)
	if err != nil {
		vatlog.Logger.WithField("component", "cmd").Error("Error connecting to database: " + err.Error())
		return err
	}

	return nil
}

func BuildRedis() {
	vatlog.Logger.WithField("component", "cmd").Info("Connecting to Redis")
	opts := database.RedisOptions{
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.DB,

		Sentinel:      config.Cfg.Redis.Sentinel,
		MasterName:    config.Cfg.Redis.MasterName,
		SentinelAddrs: config.Cfg.Redis.SentinelAddrs,

		Addr: config.Cfg.Redis.Address,
	}

	database.ConnectRedis(opts)
}
