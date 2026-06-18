package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/luism2302/tracker_mundial/internal/models"
)

type application struct {
	logger   *slog.Logger
	stickers *models.StickerModel
}

func main() {
	fresh := flag.Bool("fresh", false, "specifies whether to build everything from scratch")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("couldn't load .env")
	}

	addr := os.Getenv("PORT")
	if addr == "" {
		log.Fatal("couldn't find PORT in .env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("couldn't find DATABASE_URL in .env")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("couldn't connect to database: %s", err.Error())
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("ping to database failed: %s", err.Error())
	}

	app := application{
		logger:   logger,
		stickers: &models.StickerModel{DB: conn},
	}

	if *fresh {
		if err := app.buildAlbum(); err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("API up and running at port %s", addr)
	http.ListenAndServe(addr, app.routes())
}
