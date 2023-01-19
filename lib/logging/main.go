package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

// Setup sets up the logging format for the project
func init() {
	environment := os.Getenv("ENVIRONMENT")

	if strings.ToUpper(environment) == "DEVELOPMENT" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
}
