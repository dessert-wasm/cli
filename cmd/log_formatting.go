package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type myFormatter struct {
	logrus.TextFormatter
}

var cuteFormat = map[logrus.Level]string{
	logrus.PanicLevel: "🚨",
	logrus.FatalLevel: "!!!",
	logrus.ErrorLevel: "ERR",
	logrus.WarnLevel:  "WRN",
	logrus.InfoLevel:  "🍰 ",
	logrus.DebugLevel: "🔍",
	logrus.TraceLevel: "🐾",
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	prefix := cuteFormat[entry.Level]
	return []byte(fmt.Sprintf("%s %s\n", prefix, entry.Message)), nil
}
