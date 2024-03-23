package middlewares

import (
	"log/slog"
	"net/http"
	"os"
)

// Logger for requests.
type Logger struct {
	params slog.Attr
	logger *slog.Logger
}

// NewLogger is a constructor for request logger.
func NewLogger() *Logger {
	return &Logger{logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))}
}

func (l *Logger) setParameters(params ...string) {
	l.params = slog.Group("request", "params", params)
}

// LogScheduleAction is a basic handler logging.
func (l *Logger) LogScheduleAction(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		l.setParameters("method:", r.Method, "url:", r.URL.Path)
		l.logger.Info("scheduler", l.params)
		next(w, r)
	}
}
