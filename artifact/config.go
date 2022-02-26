package artifact

import (
	"github.com/goldeneggg/structil"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var Config *Configuration

type Configuration struct {
	RegisteredConfigStruct map[string]interface{}
	LoadedConfig           map[string]interface{}
}

func NewConfig() *Configuration {
	return &Configuration{
		RegisteredConfigStruct: make(map[string]interface{}),
	}
}

func (configuration *Configuration) Load() map[string]interface{} {
	newConfig := make(map[string]interface{})
	for name, value := range configuration.RegisteredConfigStruct {
		viper.SetConfigFile(".env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("cannot read configuration")
		}

		err = viper.Unmarshal(&value)
		if err != nil {
			log.Fatal("environment cant be loaded: ", err)
		}

		defaults.SetDefaults(value)

		newConfig[name] = value

	}

	configuration.LoadedConfig = newConfig

	return newConfig
}

func (configuration *Configuration) AddConfig(name string, userConfig interface{}) *Configuration {
	configuration.RegisteredConfigStruct[name] = userConfig

	return configuration
}

func (configuration Configuration) prepareConfigKey(key string) (string, string) {
	var rootKey string

	splitKey := strings.Split(key, ".")

	rootKey, splitKey = splitKey[0], splitKey[1:]

	key = strings.Join(splitKey, ".")

	return rootKey, key
}

func (configuration *Configuration) GetString(key string) string {

	rootKey, key := configuration.prepareConfigKey(key)

	newGetter, _ := structil.NewGetter(configuration.LoadedConfig[rootKey])

	value, _ := newGetter.String(key)

	return value
}

func (configuration Configuration) GetInt(key string) (int, bool) {

	rootKey, key := configuration.prepareConfigKey(key)

	newGetter, _ := structil.NewGetter(configuration.LoadedConfig[rootKey])

	return newGetter.Int(key)
}

func (configuration Configuration) Get(key string) (interface{}, bool) {

	rootKey, key := configuration.prepareConfigKey(key)

	newGetter, _ := structil.NewGetter(configuration.LoadedConfig[rootKey])

	return newGetter.Get(key)
}
