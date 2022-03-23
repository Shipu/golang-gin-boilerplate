package config

type MongoConfig struct {
	Username   string `mapstructure:"MONGO_USER" default:""`
	Password   string `mapstructure:"MONGO_PASS" default:""`
	Host       string `mapstructure:"MONGO_HOST" default:""`
	Port       string `mapstructure:"MONGO_PORT" default:""`
	Database   string `mapstructure:"MONGO_DATABASE" default:""`
	Connection string `mapstructure:"MONGO_CONNECTION" default:"mongodb"`
}
