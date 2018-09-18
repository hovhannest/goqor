package logger

import (
	config "goqor1.0/config/Interface"
	"goqor1.0/logger/Interface"
	"goqor1.0/logger/logrusLogger"
)

type Logger = Interface.Logger

type LogLevel = Interface.LogLevel

type LogConfig = logrusLogger.LogConfig

var _lastConfig LogConfig

// NewLogger creates a logger object with the specified logrus.Logger and the fields that should be added to every message.
func NewLogger(fields_optional ...interface{}) Logger {
	var (
		_config LogConfig
		_package string
		_fields map[string]interface{}
		)
	for i := 0; i < 3; i++ { // this function could hav maximum 3 arguments
		if len(fields_optional) > i {
			switch fields_optional[i].(type) {
			case string:
				_package = fields_optional[i].(string)
				break
			case map[string]interface{}:
				_fields = fields_optional[i].(map[string]interface{})
				break
			case LogConfig:
				_config = fields_optional[i].(LogConfig)
				break
			}
		}
	}
	if(_config == nil) {
		_config = _lastConfig
	} else {
		_lastConfig = _config
	}
	return logrusLogger.NewLogger(_config, _package, _fields)
}

func NewLogConfig(config *config.Config) logrusLogger.LogGoQorConfig {
	return  logrusLogger.LogGoQorConfig{config}
}