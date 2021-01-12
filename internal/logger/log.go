package logger

import (
	"fmt"
	"github.com/trinhdaiphuc/social-network/config"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type AppLog struct {
	*logrus.Logger
}

var (
	loggerOnce sync.Once
	appLog     *AppLog
)

func caller() (timeString, packageName, funcName, filename, line string) {
	pc, file, l, _ := runtime.Caller(3)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)

	funcName = parts[pl-1]
	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}
	filename = filepath.Base(file)
	line = fmt.Sprint(l)
	timeString = time.Now().Format(time.RFC3339)
	return
}

func NewAppLog() *AppLog {
	loggerOnce.Do(func() {
		appLog = &AppLog{
			logrus.New(),
		}
		// set log format
		if config.GetConfig().Env == "production" {
			logrus.SetFormatter(&logrus.JSONFormatter{})
		} else {
			logrus.SetFormatter(&logrus.TextFormatter{})
		}

		// set log level
		switch config.GetConfig().LogLevel {
		case "TRACE":
			appLog.SetLevel(logrus.TraceLevel)
		case "DEBUG":
			appLog.SetLevel(logrus.DebugLevel)
		case "INFO":
			appLog.SetLevel(logrus.InfoLevel)
		case "WARNING":
			appLog.SetLevel(logrus.WarnLevel)
		case "ERROR":
			appLog.SetLevel(logrus.ErrorLevel)
		case "CRITICAL":
			appLog.SetLevel(logrus.PanicLevel)
		case "FATAL":
			appLog.SetLevel(logrus.FatalLevel)
		default:
			appLog.SetLevel(logrus.WarnLevel)
		}

		// set log output
		if len(config.GetConfig().LogPath) > 0 {
			accessLogFileHandler, err := os.OpenFile(config.GetConfig().LogPath, os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				panic(err)
			}
			appLog.Out = accessLogFileHandler
		}

	})

	return appLog
}

func (appLog *AppLog) Info(args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Info(args...)
}

func (appLog *AppLog) Debug(args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Debug(args...)
}

func (appLog *AppLog) Warn(args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Warn(args...)
}

func (appLog *AppLog) Error(args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Error(args...)
}

func (appLog *AppLog) Fatal(args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Fatal(args...)
}

func (appLog *AppLog) Panic(args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Panic(args...)
}
