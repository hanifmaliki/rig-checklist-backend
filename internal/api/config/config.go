package config

type Config struct {
	Port           string `envconfig:"APP_PORT" default:"80"`
	PrivateKeyPath string `envconfig:"PRIVATE_KEY_PATH" default:"./private-key.pem"`
	DevMode        bool   `envconfig:"DEV_MODE" default:"false"`
	Log            *ConfigLog
}

type ConfigLog struct {
	Level int `envconfig:"LOG_LEVEL" default:"0"`
}
