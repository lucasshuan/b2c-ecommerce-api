package log

import (
	"os"

	"github.com/charmbracelet/log"
)

var AppLogger *log.Logger

func Init() {
	AppLogger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix:          "\x1b[94m[api]",
		ReportTimestamp: true,
		TimeFormat:      "01/02 15:04:05",
	})
	AppLogger.SetLevel(log.DebugLevel)
	replaceGinWriter()
}
