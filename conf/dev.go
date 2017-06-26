package conf

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}
