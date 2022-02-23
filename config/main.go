package config

func RegisterConfig() (AppConfig, DatabaseConfig) {
	appConfig := AppConfig{}
	databaseConfig := DatabaseConfig{}

	return appConfig, databaseConfig
}
