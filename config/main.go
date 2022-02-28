package config

import (
	. "github.com/shipu/artifact"
)

func RegisterConfig() {
	Config.AddConfig("App", new(AppConfig))
	Config.AddConfig("DB", new(DatabaseConfig))

	Config.Load()
}
