package logger

import (
	"io"
	"log"
	"os"
)

//TODO: REFACTOR AND REORGANIZE

//TODO: add in concept of ALERT logging
type LoggingLevel int

const (
	LoggingOff LoggingLevel = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

const (
	MicroSecondsOnly int = log.Lmicroseconds
	DateOnly             = log.Ldate
	TimeOnly             = log.Ltime
	FullSourceFile       = log.Llongfile
	SourceFile           = log.Lshortfile
	DefaultFormat        = log.LstdFlags
)

var defaultLoggingLevel LoggingLevel = WarnLevel
var defaultOutput io.Writer = os.Stdout

type LevelLogger struct {
	Logger *log.Logger
	level  LoggingLevel
}

func Logger(out io.Writer, prefix string, flag int) *LevelLogger {
	logger := new(LevelLogger)
	logger.level = defaultLoggingLevel
	logger.Logger = log.New(out, prefix, flag)
	return logger
}
func DefaultLogger() *LevelLogger {
	return Logger(defaultOutput, "", DefaultFormat)
}

func DefaultLoggerWithName(name string) *LevelLogger {
	return Logger(defaultOutput, name, DefaultFormat)
}

func (logger *LevelLogger) LoggingFlags(flags int) *LevelLogger {
	logger.Logger.SetFlags(flags)
	return logger
}
func (logger *LevelLogger) LoggingPrefix(prefix string) *LevelLogger {
	logger.Logger.SetPrefix(prefix)
	return logger
}

func (logger *LevelLogger) IncreaseVerbosity() *LevelLogger {
	logger.level++
	return logger
}
func (logger *LevelLogger) DecreaseVerbosity() *LevelLogger {
	if logger.IsOff() {
		return logger
	}
	logger.level--
	return logger
}
func (logger *LevelLogger) TurnOff() *LevelLogger {
	logger.level = LoggingOff
	return logger
}
func (logger *LevelLogger) IsOn() bool {
	return logger.level > LoggingOff
}
func (logger *LevelLogger) IsOff() bool {
	return logger.level <= LoggingOff
}
func (logger *LevelLogger) IsError() bool {
	return logger.level >= ErrorLevel
}
func (logger *LevelLogger) IsWarn() bool {
	return logger.level >= WarnLevel
}
func (logger *LevelLogger) IsInfo() bool {
	return logger.level >= InfoLevel
}
func (logger *LevelLogger) IsDebug() bool {
	return logger.level >= DebugLevel
}
func (logger *LevelLogger) IsTrace() bool {
	return logger.level >= TraceLevel
}
func (logger *LevelLogger) SetError() {
	logger.level = ErrorLevel
}
func (logger *LevelLogger) SetWarn() {
	logger.level = WarnLevel
}
func (logger *LevelLogger) SetInfo() {
	logger.level = InfoLevel
}
func (logger *LevelLogger) SetDebug() {
	logger.level = DebugLevel
}
func (logger *LevelLogger) SetTrace() {
	logger.level = TraceLevel
}

func (logger *LevelLogger) Error(msg ...interface{}) *LevelLogger {
	if logger.IsError() {
		logger.Logger.Println(msg)
	}

	return logger
}
func (logger *LevelLogger) Warn(msg ...interface{}) *LevelLogger {
	if logger.IsWarn() {
		logger.Logger.Println(msg)
	}

	return logger
}
func (logger *LevelLogger) Info(msg ...interface{}) *LevelLogger {
	if logger.IsInfo() {
		logger.Logger.Println(msg)
	}

	return logger
}
func (logger *LevelLogger) Debug(msg ...interface{}) *LevelLogger {
	if logger.IsDebug() {
		logger.Logger.Println(msg)
	}

	return logger
}
func (logger *LevelLogger) Trace(msg ...interface{}) *LevelLogger {
	if logger.IsTrace() {
		logger.Logger.Println(msg)
	}

	return logger
}
