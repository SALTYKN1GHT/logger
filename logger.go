package logger

import (
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

const (
	LightBlue = "\033[36m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Red       = "\033[31m"
	Reset     = "\033[0m"
)

type LoggerConfig struct {
	LogLevel     LogLevel
	LogToFile    bool
	LogToConsole bool
	FilePath     string
}

type Logger struct {
	level         LogLevel
	consoleLogger *log.Logger
	fileLogger    *log.Logger
}

func NewLogger(config LoggerConfig) (*Logger, error) {
	var consoleLogger, fileLogger *log.Logger

	if config.LogToConsole {
		consoleLogger = log.New(os.Stdout, "", log.LstdFlags)
	}

	if config.LogToFile {
		file, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		fileLogger = log.New(file, "", log.LstdFlags)
	}

	return &Logger{
		level:         config.LogLevel,
		consoleLogger: consoleLogger,
		fileLogger:    fileLogger,
	}, nil
}

func (l *Logger) log(level LogLevel, msg string) {
	if level >= l.level {
		color := levelToColor(level)
		levelStr := levelToString(level)

		if l.consoleLogger != nil {
			l.consoleLogger.Printf("%s[%s]%s %s", color, levelStr, Reset, msg)
		}

		if l.fileLogger != nil {
			l.fileLogger.Printf("[%s] %s", levelStr, msg)
		}
	}
}

func (l *Logger) Debug(msg string) {
	l.log(DEBUG, msg)
}

func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

func (l *Logger) Warning(msg string) {
	l.log(WARNING, msg)
}

func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}

func levelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func levelToColor(level LogLevel) string {
	switch level {
	case DEBUG:
		return LightBlue
	case INFO:
		return Green
	case WARNING:
		return Yellow
	case ERROR:
		return Red
	default:
		return Reset
	}
}
