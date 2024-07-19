package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func GetLogger() *logrus.Logger {
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	// Configurar a saída do log
	logger.SetOutput(os.Stdout)

	// Definir o nível de log
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
