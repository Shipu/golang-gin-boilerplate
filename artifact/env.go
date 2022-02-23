package artifact

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppName      string `mapstructure:"APP_Name"`
	Port         int    `mapstructure:"APP_PORT"`
	Debug        string `mapstructure:"APP_DEBUG"`
	URL          string `mapstructure:"APP_URL"`
	Environment  string `mapstructure:"APP_ENV"`
	DBUsername   string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASS"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBDatabase   string `mapstructure:"DB_DATABASE"`
	DBConnection string `mapstructure:"DB_CONNECTION"`
	GinMode      string `mapstructure:"GIN_MODE"`
}

func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}

	return env
}
