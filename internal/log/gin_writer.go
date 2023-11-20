package log

import (
	"strings"

	"github.com/charmbracelet/log"
)

type CustomGinWriter struct {
	logger *log.Logger
}

func (lw *CustomGinWriter) Write(p []byte) (n int, err error) {
	str := string(p)
	prefix := deleteGinPrefix(&str)
	if prefix == "GIN" {
		return len(p), nil
	}

	start := strings.Index(str, "[")
	end := strings.Index(str, "]")

	var strInBrackets string

	if start != -1 && end > start {
		strInBrackets = str[start+1 : end]
		str = str[end+2:]
	}

	str = strings.TrimRight(str, " \t\n\r")

	methods := []string{"get", "post", "put", "delete", "patch"}

	for _, method := range methods {
		if strings.HasPrefix(strings.ToLower(str), method) {
			startRoute := strings.Index(str, "/")
			endRoute := strings.Index(str[startRoute:], " ")
			uri := str[startRoute : startRoute+endRoute]
			lw.logger.Debug("Route matched:", "method", method, "uri", uri)
			return 0, nil
		}
	}

	switch strInBrackets {
	case "debug":
		lw.logger.Debug(str)
	case "info":
		lw.logger.Info(str)
	case "warn":
		lw.logger.Warn(str)
	case "error":
		lw.logger.Error(str)
	default:
		lw.logger.Debug(str)
	}

	return len(p), nil
}

func deleteGinPrefix(str *string) string {
	start := strings.Index(*str, "[")
	end := strings.Index(*str, "]")
	prefix := (*str)[start+1 : end]
	if start != -1 && end != -1 && end > start {
		*str = strings.Replace(*str, (*str)[start:end+2], "", 1)
	}
	return prefix
}

func GetCustomGinWriter() *CustomGinWriter {
	return &CustomGinWriter{
		logger: AppLogger.WithPrefix("[gin]"),
	}
}
