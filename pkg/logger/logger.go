package logger

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

func Init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})
	logrus.SetLevel(logrus.InfoLevel)
}
