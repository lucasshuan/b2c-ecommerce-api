package log

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var AppLogger *log.Logger

func Init() {
	AppLogger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix:          "\x1b[94m[api]",
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
	AppLogger.SetLevel(log.DebugLevel)
	replaceGinWriter()
}
