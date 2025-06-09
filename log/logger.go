package log

const (
	LOG_DEBUG uint8 = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_FATAL
)

var CurrentLevel = LOG_INFO

type Logger interface {
	Debug(value any)
	Info(value any)
	Warn(value any)
	Error(value any)
	Fatal(value any)
	Debugf(format string, value ...any)
	Infof(format string, value ...any)
	Warnf(format string, value ...any)
	Errorf(format string, value ...any)
	Fatalf(format string, value ...any)
	Println(logLevel uint8, value any)
}
