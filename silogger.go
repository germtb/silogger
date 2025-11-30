package silogger

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/germtb/sidb"
	"github.com/google/uuid"
)

type Logger struct {
	db    *sidb.Database
	level LogLevel
}

func InitLogger(db *sidb.Database) *Logger {
	log.SetFlags(0) // Disable default timestamp logging
	color.NoColor = false

	return &Logger{
		db:    db,
		level: INFO,
	}
}

type LogLevel uint64

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Define colors for each log level
var (
	debugColor     = color.New(color.FgCyan).SprintFunc()
	infoColor      = color.New(color.FgGreen).SprintFunc()
	warnColor      = color.New(color.FgYellow).SprintFunc()
	errorColor     = color.New(color.FgRed).SprintFunc()
	fatalColor     = color.New(color.FgHiRed, color.Bold).SprintFunc()
	timestampColor = color.New(color.FgHiBlack, color.Faint).SprintFunc()
)

func (logger *Logger) Debug(args ...any) {
	if logger.level > DEBUG {
		return
	}

	message := fmt.Sprint(args...)

	if logger.db != nil {
		go logger.db.Upsert(sidb.EntryInput{
			Key:   uuid.NewString(),
			Value: []byte(message),
			Type:  "debug",
		})
	}

	logger.logWithLevel("DEBUG", debugColor, message)
}

func (logger *Logger) Info(args ...any) {
	if logger.level > INFO {
		return
	}

	message := fmt.Sprint(args...)

	if logger.db != nil {
		go logger.db.Upsert(sidb.EntryInput{
			Key:   uuid.NewString(),
			Value: []byte(message),
			Type:  "info",
		})
	}

	logger.logWithLevel("INFO", infoColor, message)
}

func (logger *Logger) Warn(args ...any) {
	if logger.level > WARN {
		return
	}

	message := fmt.Sprint(args...)

	if logger.db != nil {
		go logger.db.Upsert(sidb.EntryInput{
			Key:   uuid.NewString(),
			Value: []byte(message),
			Type:  "warn",
		})
	}

	logger.logWithLevel("WARN", warnColor, message)
}

func (logger *Logger) Error(args ...any) {
	if logger.level > ERROR {
		return
	}

	message := fmt.Sprint(args...)

	if logger.db != nil {
		go logger.db.Upsert(sidb.EntryInput{
			Key:   uuid.NewString(),
			Value: []byte(message),
			Type:  "error",
		})
	}

	logger.logWithLevel("ERROR", errorColor, message)
}

func (logger *Logger) Fatal(args ...any) {
	if logger.level > FATAL {
		return
	}

	message := fmt.Sprint(args...)

	if logger.db != nil {
		go logger.db.Upsert(sidb.EntryInput{
			Key:   uuid.NewString(),
			Value: []byte(message),
			Type:  "fatal",
		})
	}

	logger.logWithLevel("FATAL", fatalColor, message)
}

func (logger *Logger) SetLevel(level LogLevel) {
	logger.level = level
}

func (logger *Logger) GetLevel() LogLevel {
	return logger.level
}

func (logger *Logger) SetDatabase(db *sidb.Database) {
	logger.db = db
}

func (logger *Logger) logWithLevel(level string, colorFunc func(a ...any) string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := fmt.Sprintf("[%s]", level)
	paddedLevel := fmt.Sprintf("%-7s", levelStr)
	log.Println(colorFunc(paddedLevel), timestampColor(timestamp)+" ", message)
}
