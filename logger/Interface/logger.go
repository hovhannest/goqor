package Interface

type LogLevel int

const(
	NoLog LogLevel = 0
	Error LogLevel =1
	Warn LogLevel =2
	Info LogLevel = 3
	Debug LogLevel = 4
	All LogLevel = 5
)

type Logger interface {
	// adds a field that should be added to every message being logged
	SetField(name, value string)

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
