package tint_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/lmittmann/tint"
	"golang.org/x/exp/slog"
)

// TestSlog implements a simple example on how to use tint. This test performs no assertions but is still
// useful to manually inspect how tint handles various slog features.
func TestSlog(t *testing.T) {
	slog.SetDefault(slog.New(tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}.NewHandler(os.Stderr)))

	slog.Info("Starting server", "addr", ":8080", "env", "production")
	slog.Debug("Connected to DB", "db", "myapp", "host", "localhost:5432")
	slog.Warn("Slow request", "method", "GET", "path", "/users", "duration", 497*time.Millisecond)
	slog.Error("DB connection lost", tint.Err(errors.New("connection reset")), "db", "myapp")

	// grouped attributes
	attrs := slog.Group("req", slog.String("method", "GET"), slog.String("path", "/users"), slog.Any("duration", 497*time.Millisecond))
	slog.Debug("Slow request (grouped attributes)", attrs)
}
