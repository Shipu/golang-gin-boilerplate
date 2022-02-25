package config

type AppConfig struct {
	Name        string `mapstructure:"APP_NAME" default:"Artifact"`
	Environment string `mapstructure:"APP_ENV" default:"local"`
	Debug       bool   `mapstructure:"APP_DEBUG" default:"true"`
	Url         string `mapstructure:"APP_URL"  default:"http://localhost"`
	Port        int    `mapstructure:"APP_PORT" default:"8098"`
	TimeZone    string `mapstructure:"APP_TIMEZONE"  default:"UTC"`
	Locale      string `mapstructure:"APP_LOCALE"  default:"en"`
	GinMode     string `mapstructure:"GIN_MODE" default:"debug"`
}
