package log

import (
	"fmt"
	"sync"

	"cloud.google.com/go/logging"
)

// Logger is the interface for loggers used in the application components.
type Logger interface {
	Log(e logging.Entry)

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Notice(args ...interface{})
	Noticef(format string, args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Critical(args ...interface{})
	Criticalf(format string, args ...interface{})

	Alert(args ...interface{})
	Alertf(format string, args ...interface{})

	Emergency(args ...interface{})
	Emergencyf(format string, args ...interface{})

	SetSeverity(severity string)
}

type logger struct {
	golog    *logging.Logger
	severity logging.Severity
}

func (l logger) log(s logging.Severity, e logging.Entry) {
	e.Severity = s
	l.golog.Log(e)
}

func (l logger) Log(e logging.Entry) {
	l.log(e.Severity, e)
}

func (l logger) Debug(args ...interface{}) {
	l.log(logging.Debug, logging.Entry{Payload: fmt.Sprint(args...)})
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l logger) Info(args ...interface{}) {
	l.log(logging.Info, logging.Entry{Payload: fmt.Sprint(args...)})
}

func (l logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l logger) Notice(args ...interface{}) {
	l.log(logging.Notice, logging.Entry{Payload: fmt.Sprint(args...)})
}

func (l logger) Noticef(format string, args ...interface{}) {
	l.Notice(fmt.Sprintf(format, args...))
}

func (l logger) Warning(args ...interface{}) {
	l.log(logging.Warning, logging.Entry{Payload: fmt.Sprint(args...)})
}

func (l logger) Warningf(format string, args ...interface{}) {
	l.Warning(fmt.Sprintf(format, args...))
}

func (l logger) Error(args ...interface{}) {
	l.log(logging.Error, logging.Entry{Payload: fmt.Sprint(args...)})
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l logger) Critical(args ...interface{}) {
	l.log(logging.Critical, logging.Entry{Payload: fmt.Sprint(args...)})
}

func (l logger) Criticalf(format string, args ...interface{}) {
	l.Critical(fmt.Sprintf(format, args...))
}

func (l logger) Alert(args ...interface{}) {
	l.log(logging.Alert, logging.Entry{Payload: fmt.Sprint(args...)})
}

func (l logger) Alertf(format string, args ...interface{}) {
	l.Alert(fmt.Sprintf(format, args...))
}

func (l logger) Emergency(args ...interface{}) {
	l.log(logging.Emergency, logging.Entry{Payload: fmt.Sprint(args...)})
}

func (l logger) Emergencyf(format string, args ...interface{}) {
	l.Emergency(fmt.Sprintf(format, args...))
}

func (l logger) SetSeverity(severity string) {
	l.severity = logging.ParseSeverity(severity)
}

// NewLogger returns a new Logger.
func NewLogger(client *logging.Client, name string) Logger {
	return logger{
		golog:    client.Logger(name),
		severity: logging.Info,
	}
}

var base logger
var once sync.Once

// InitLogger initializes base logger.
func InitLogger(client *logging.Client, name string) {
	once.Do(func() {
		base = logger{
			golog:    client.Logger(name),
			severity: logging.Info,
		}
	})
}

func Log(e logging.Entry) {
	base.Log(e)
}

func Debug(args ...interface{}) {
	base.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	base.Debugf(format, args)
}

func Info(args ...interface{}) {
	base.Info(args)
}

func Infof(format string, args ...interface{}) {
	base.Infof(format, args)
}

func Notice(args ...interface{}) {
	base.Notice(args)
}

func Noticef(format string, args ...interface{}) {
	base.Noticef(format, args)
}

func Warning(args ...interface{}) {
	base.Warning(args)
}

func Warningf(format string, args ...interface{}) {
	base.Warningf(format, args)
}

func Error(args ...interface{}) {
	base.Error(args)
}

func Errorf(format string, args ...interface{}) {
	base.Errorf(format, args)
}

func Critical(args ...interface{}) {
	base.Critical(args)
}

func Criticalf(format string, args ...interface{}) {
	base.Criticalf(format, args)
}

func Alert(args ...interface{}) {
	base.Alert(args)
}

func Alertf(format string, args ...interface{}) {
	base.Alertf(format, args)
}

func Emergency(args ...interface{}) {
	base.Emergency(args)
}

func Emergencyf(format string, args ...interface{}) {
	base.Emergencyf(format, args)
}

func SetSeverity(severity string) {
	base.SetSeverity(severity)
}
