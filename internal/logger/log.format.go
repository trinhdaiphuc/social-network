package logger

import (
	"github.com/sirupsen/logrus"
)

func (appLog *AppLog) Infof(format string, args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Infof(format, args...)
}

func (appLog *AppLog) Debugf(format string, args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Debugf(format, args...)
}

func (appLog *AppLog) Warnf(format string, args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	go entry.Warnf(format, args...)
}

func (appLog *AppLog) Errorf(format string, args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Errorf(format, args...)
}

func (appLog *AppLog) Fatalf(format string, args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Fatalf(format, args...)
}

func (appLog *AppLog) Panicf(format string, args ...interface{}) {
	_, packageName, funcName, filename, line := caller()
	entry := appLog.WithFields(logrus.Fields{
		"PACKAGE": packageName,
		"FILE":    filename,
		"LINE":    line,
		"FUNC":    funcName,
	})
	entry.Panicf(format, args...)
}
