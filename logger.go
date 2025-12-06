package logger

import (
	"log/slog"
	"os"
)

// Config définit les options de configuration du logger.
type Config struct {
	EnableJSON bool 
	Debug      bool 
	ServiceName string 
}

// Init configure le logger global.
func Init(cfg Config) {
	logLevel := slog.LevelInfo
	if cfg.Debug {
		logLevel = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
		AddSource: cfg.Debug, 
	}

	var handler slog.Handler
	
	// Utilisation de os.Stderr, le standard pour les logs.
	if cfg.EnableJSON {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	} else {
		handler = slog.NewTextHandler(os.Stderr, opts)
	}

	logger := slog.New(handler)
	
	if cfg.ServiceName != "" {
		logger = logger.With("service", cfg.ServiceName)
	}

	slog.SetDefault(logger)
}

// Get retourne l'instance du logger par défaut
func Get() *slog.Logger {
	return slog.Default()
}