package config

type Config struct {
	Server   ConfigServer   `json:"server"`
	Database ConfigDatabase `json:"database"`
}

type ConfigServer struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type ConfigDatabase struct {
	Driver      string `json:"driver"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	AutoMigrate bool   `json:"automigrate"`
}
