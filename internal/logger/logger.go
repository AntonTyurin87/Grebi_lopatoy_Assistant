package logger

import (
	"log"
	"os"
	"strings"
)

// TODO файл и/или область для логирования

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	ErrorLevel
)

type Logger struct {
	*log.Logger
	level Level
}

func New(level string) *Logger {
	var l Level
	switch strings.ToLower(level) {
	case "debug":
		l = DebugLevel
	case "info":
		l = InfoLevel
	case "error":
		l = ErrorLevel
	default:
		l = InfoLevel
	}

	return &Logger{
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		level:  l,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level <= DebugLevel {
		l.Logger.SetPrefix("[DEBUG] ")
		l.Println(v...)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level <= InfoLevel {
		l.Logger.SetPrefix("[INFO] ")
		l.Println(v...)
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.level <= ErrorLevel {
		l.Logger.SetPrefix("[ERROR] ")
		l.Println(v...)
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Logger.SetPrefix("[FATAL] ")
	l.Fatal(v...)
}
