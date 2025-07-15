package helpers

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func SetupLogger() {
	logg := logrus.New()

	logg.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logg.Info("Loger Initalized")
	Logger = logg
}
