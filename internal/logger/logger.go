package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var Log *log.Logger

func Init() {
	Log = log.NewWithOptions(os.Stderr, log.Options{
		Prefix:          "\x1b[94m[api]",
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
	Log.SetLevel(log.DebugLevel)
	replaceGinWriter()
}
