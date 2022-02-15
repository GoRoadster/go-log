package log

import (
	"fmt"
	"io"
	"os"
	osPath "path"
	"time"

	logging "github.com/sirupsen/logrus"
)

const (
	PANIC = "panic"
	FATAL = "fatal"
	ERROR = "error"
	WARN  = "warn"
	INFO  = "info"
	DEBUG = "debug"
	TRACE = "trace"
)

var (
	_ = Trace
	_ = Debug
	_ = Info
	_ = Warn
	_ = Error
	_ = Fatal
	_ = InitLogger

	_ = PANIC
	_ = FATAL
	_ = ERROR
	_ = WARN
	_ = INFO
	_ = DEBUG
	_ = TRACE
)

func InitLogger(path, logPrefix, logLevel string, shouldSave bool) error {
	if shouldSave {
		file, err := os.OpenFile(
			getLogFileDir(path, logPrefix),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
		)
		if err != nil {
			return err
		}

		mw := io.MultiWriter(file, os.Stdout)
		logging.SetOutput(mw)
	}

	logging.SetFormatter(&Formatter{})
	lvl, err := logging.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	logging.SetLevel(lvl)

	return nil
}

func Trace(message string, params ...interface{}) {
	if hasParseableFields(params...) {
		logging.WithFields(makeFields(params...)).Trace(message)
		return
	}

	logging.Trace(message)
}

func Debug(message string, params ...interface{}) {
	if hasParseableFields(params...) {
		logging.WithFields(makeFields(params...)).Debug(message)
		return
	}

	logging.Debug(message)
}

func Info(message string, params ...interface{}) {
	if hasParseableFields(params...) {
		logging.WithFields(makeFields(params...)).Info(message)
		return
	}

	logging.Info(message)
}

func Warn(message string, params ...interface{}) {
	if hasParseableFields(params...) {
		logging.WithFields(makeFields(params...)).Warn(message)
		return
	}

	logging.Warn(message)
}

func Error(message string, params ...interface{}) {
	if hasParseableFields(params...) {
		logging.WithFields(makeFields(params...)).Error(message)
		return
	}

	logging.Error(message)
}

func Fatal(message string, params ...interface{}) {
	if hasParseableFields(params...) {
		logging.WithFields(makeFields(params...)).Fatal(message)
		return
	}

	logging.Fatal(message)
}

func makeFields(params ...interface{}) logging.Fields {
	m := make(logging.Fields)
	for i := 0; i < len(params); i += 2 {
		k, ok := params[i].(string)
		if !ok {
			continue
		}
		m[k] = params[i+1]
	}
	return m
}

func hasParseableFields(params ...interface{}) bool {
	return len(params) != 0 && len(params)%2 == 0
}

func getLogFileDir(path, filePrefix string) string {
	return osPath.Join(path, fmt.Sprintf("%s-%s.log", filePrefix, time.Now().UTC().Format(time.RFC822)))
}
