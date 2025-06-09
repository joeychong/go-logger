package console

import (
	"ebizzone.com/logger/log"
	"fmt"
	"time"
)

type Style struct {
	label string
	color string
	full  bool
}

// var label = [5]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
var styles = [5]Style{
	{label: "DEBUG", color: "\x1b[32m", full: false},
	{label: "INFO", color: "\x1b[97m", full: false},
	{label: "WARN", color: "\x1b[93m", full: false},
	{label: "ERROR", color: "\x1b[91m", full: false},
	{label: "FATAL", color: "\x1b[93;101m", full: true},
}

const reset = "\x1b[0m"

type Console struct {
	tag string
}

func (c Console) Println(logLevel uint8, value any) {
	now := time.Now()
	if log.CurrentLevel <= logLevel {
		style := styles[logLevel]
		if style.full {
			fmt.Printf("%s%s [%s] %s - %v%s\n", style.color, now.Format("2006-01-02 15:04:05.000"), style.label, c.tag, value, reset)
		} else {
			fmt.Printf("%s [%s%s%s] %s - %v\n", now.Format("2006-01-02 15:04:05.000"), style.color, style.label, reset, c.tag, value)
		}
	}
}

func (c Console) Debug(value any) {
	c.Println(log.LOG_DEBUG, value)
}

func (c Console) Info(value any) {
	c.Println(log.LOG_INFO, value)
}

func (c Console) Warn(value any) {
	c.Println(log.LOG_WARN, value)
}

func (c Console) Error(value any) {
	c.Println(log.LOG_ERROR, value)
}

func (c Console) Fatal(value any) {
	c.Println(log.LOG_FATAL, value)
}

func (c Console) Debugf(format string, value ...any) {
	c.Println(log.LOG_DEBUG, fmt.Sprintf(format, value...))
}

func (c Console) Infof(format string, value ...any) {
	c.Println(log.LOG_INFO, fmt.Sprintf(format, value...))
}

func (c Console) Warnf(format string, value ...any) {
	c.Println(log.LOG_WARN, fmt.Sprintf(format, value...))
}

func (c Console) Errorf(format string, value ...any) {
	c.Println(log.LOG_ERROR, fmt.Sprintf(format, value...))
}

func (c Console) Fatalf(format string, value ...any) {
	c.Println(log.LOG_FATAL, fmt.Sprintf(format, value...))
}

func NewLogger(tag string) log.Logger {
	console := Console{}
	console.tag = tag
	return console
}
