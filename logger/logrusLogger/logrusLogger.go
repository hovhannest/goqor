package logrusLogger

import (
	"github.com/sirupsen/logrus"
	config "goqor1.0/config/Interface"
	"goqor1.0/logger/Interface"
	"os"
)

type LogConfig interface {
	GetLogLevel() Interface.LogLevel
	SetLogLevel(level Interface.LogLevel)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
type LogGoQorConfig struct {
	 Config *config.Config
}

func (lc LogGoQorConfig) GetLogLevel() Interface.LogLevel {
	return Interface.LogLevel((*lc.Config).GetInt(keyForPackage("level")))
}

func (lc LogGoQorConfig) SetLogLevel(level Interface.LogLevel) {
	(*lc.Config).Set(keyForPackage("level"), int(level))
}
////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	_config LogConfig
	_packageName = "log"
)

func keyForPackage(key string) string {
	return  _packageName + "." + key
}

// logger wraps logrus.Logger so that it can log messages sharing a common set of fields.
type logrusHolder struct {
	logger *logrus.Logger
}

var _holder logrusHolder = logrusHolder{logger: logrus.New()}

type logger struct {
	holder *logrusHolder
	fields logrus.Fields
}

func init() {
	_holder.logger.Level = logrus.DebugLevel // Log level control is done throw config package
	_holder.logger.SetOutput(os.Stderr)
	//// You could set this to any `io.Writer` such as a file
	//file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//	_holder.logger.Out = file
	//} else {
	//	_holder.logger.Info("Failed to log to file, using default stderr")
	//}
}

// NewLogger creates a logger object with the specified logrus.Logger and the fields that should be added to every message.
func NewLogger(config LogConfig, _package string, _fields map[string]interface{}) Interface.Logger {
	if(config != nil) {
		_config = config
	} else {
		panic("Logger needs to have a LogConfig!")
	}
	fields := make(map[string]interface{})
	if(_fields != nil) {
		fields = _fields
	}
	if(_package != "") {
		fields["package"] = _package
	}
	return &logger{
		holder: &_holder,
		fields: fields,
	}
}

func (l *logger) SetField(name, value string) {
	l.fields[name] = value
}

func (l *logger) Debugf(format string, args ...interface{}) {
	if(_config.GetLogLevel() < Interface.Debug) {
		return
	}
	l.tagged().Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	if(_config.GetLogLevel() < Interface.Info) {
		return
	}
	l.tagged().Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	if(_config.GetLogLevel() < Interface.Warn) {
		return
	}
	l.tagged().Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	if(_config.GetLogLevel() < Interface.Error) {
		return
	}
	l.tagged().Errorf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	if(_config.GetLogLevel() < Interface.Debug) {
		return
	}
	l.tagged().Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	if(_config.GetLogLevel() < Interface.Info) {
		return
	}
	l.tagged().Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	if(_config.GetLogLevel() < Interface.Warn) {
		return
	}
	l.tagged().Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	if(_config.GetLogLevel() < Interface.Error) {
		return
	}
	l.tagged().Error(args...)
}

func (l *logger) tagged() *logrus.Entry {
	return l.holder.logger.WithFields(l.fields)
}