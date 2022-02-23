package config

type AppConfig struct {
	Name        string `mapstructure:"APP_NAME" default:"Artifact"`
	Environment string `env:"APP_ENV" default:"local"`
	Debug       bool   `env:"APP_DEBUG" default:"true"`
	Url         string `env:"APP_URL"  default:"http://localhost"`
	Port        int    `env:"APP_PORT" default:"8080"`
	TimeZone    string `env:"APP_TIMEZONE"  default:"UTC"`
	Locale      string `env:"APP_LOCALE"  default:"en"`
	GinMode     string `mapstructure:"GIN_MODE" default:"debug"`
}
