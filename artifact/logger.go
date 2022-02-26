package artifact

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLog LoggerBuilder

// LoggerBuilder structure
type LoggerBuilder struct {
	*zap.SugaredLogger
}

// GetLogger gets the global instance of the logger
func GetLogger() LoggerBuilder {
	return globalLog
}

// NewLogger sets up logger
func NewLogger() LoggerBuilder {

	isLocal := Config.GetString("App.Environment")

	config := zap.NewDevelopmentConfig()

	if isLocal == "local" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.Level.SetLevel(zap.PanicLevel)
	}

	logger, _ := config.Build()

	globalLog := logger.Sugar()

	return LoggerBuilder{
		globalLog,
	}

}
