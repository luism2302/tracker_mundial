package main

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := application{
		logger: logger,
	}

	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		app.logger.Error("couldn't load .env")
	}

	addr := os.Getenv("PORT")
	if addr == "" {
		app.logger.Error("couldn't find PORT in .env")
	}

	app.logger.Info("starting api", slog.Any("addr", addr))

	http.ListenAndServe(addr, app.routes())
}
