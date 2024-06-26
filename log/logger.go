package logger

import "github.com/sirupsen/logrus"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

/*
Print error message in the console
*/
func Error(format string, args ...interface{}) {
	logrus.Errorf(formateMessage(Red, format), args)
}

func Info(format string, args ...interface{}) {
	message := formateMessage(Blue, format)
	if len(args) == 0 {
		logrus.Info(message)
	} else {
		logrus.Infof(formateMessage(Blue, format), args)
	}
}

/*
Formate a given message with the specified color
*/
func formateMessage(color string, message string) string {
	return color + message + Reset
}
