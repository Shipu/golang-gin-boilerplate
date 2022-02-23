package config

type DatabaseConfig struct {
	Username   string `mapstructure:"DB_USER" default:""`
	Password   string `mapstructure:"DB_PASS" default:""`
	Host       string `mapstructure:"DB_HOST" default:""`
	Port       string `mapstructure:"DB_PORT" default:""`
	Database   string `mapstructure:"DB_DATABASE" default:""`
	Connection string `mapstructure:"DB_CONNECTION" default:""`
}
