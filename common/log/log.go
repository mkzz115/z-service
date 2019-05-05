package log

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
}

var l *logger

type logger struct {
	*logrus.Logger
}

func Init(path string) error {
	l = &logger{logrus.New()}
	l.Formatter = new(logrus.TextFormatter)
	l.Formatter.(*logrus.TextFormatter).DisableColors = true
	l.Level = logrus.InfoLevel
	isExist, err := PathExists(path)
	if err != nil {
		return err
	}
	if !isExist {
		err := os.Mkdir(filepath.Dir(path), 0755)
		if err != nil {
			return err
		}
	}
	logPath := path + ".log"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	l.Out = file
	return nil
}

func Info(format string, a ...interface{}) {
	//fmt.Printf(format, a...)
	l.Infof(format, a...)
}

func Warn(format string, a ...interface{}) {
	//fmt.Printf(format, a...)
	l.Warnf(format, a...)
}

func Error(format string, a ...interface{}) {
	//fmt.Printf(format, a...)
	l.Errorf(format, a...)
}
