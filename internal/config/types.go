package config

type Config struct {
	Server   ConfigServer   `json:"server"`
	Database ConfigDatabase `json:"database"`
	Redis    ConfigRedis    `json:"redis"`
	Session  ConfigSession  `json:"session"`
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

type ConfigRedis struct {
	// Sentinel config
	Sentinel      bool     `json:"sentinel"`
	MasterName    string   `json:"master_name"`
	SentinelAddrs []string `json:"sentinel_addrs"`

	// Standalone config
	Address string `json:"address"`

	// Common
	Password string `json:"password"`
	DB       int    `json:"database"`
}

type ConfigSession struct {
	Cookie ConfigSessionCookie `json:"cookie"`
	JWT    ConfigSessionJWT    `json:"jwt"`
}

type ConfigSessionCookie struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
	MaxAge int    `json:"max_age"`
}

type ConfigSessionJWT struct {
	JWKSPath string `json:"jwks_path"`
}
