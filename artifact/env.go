package artifact

import (
	"github.com/mcuadros/go-defaults"
	"golang-gin-boilerplate/config"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	App      config.AppConfig
	Database config.DatabaseConfig
}

func NewEnv() Env {
	appConfig, databaseConfig := config.RegisterConfig()
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read configuration")
	}

	appErr := viper.Unmarshal(&appConfig)
	databaseErr := viper.Unmarshal(&databaseConfig)
	if err != nil || appErr != nil || databaseErr != nil {
		log.Fatal("environment cant be loaded: ", err)
	}
	defaults.SetDefaults(&appConfig)
	defaults.SetDefaults(&databaseConfig)

	env := Env{appConfig, databaseConfig}

	log.Printf("%#v \n", &env)

	return env
}
