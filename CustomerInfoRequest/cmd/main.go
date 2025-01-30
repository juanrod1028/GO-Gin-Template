package main

import (
	"hexagonalExample/cmd/api/configurations"
	"hexagonalExample/cmd/api/core"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	config, err := configurations.Load("../app-config.yml")
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}
	logger.Info("Staritng server", "Name", config.App.Name, "Environment", config.App.Env)
	srv := core.New(*config, logger)

	if err := srv.Run(); err != nil {
		logger.Error("server failed", "error", err)
	}
}
