package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	albumPanini2026 = "albums/panini2026_completo.json"
)

func (app application) buildAlbum() error {
	_, err := readAlbumJSON()
	if err != nil {
		return err
	}

	return nil
}

type section struct {
	Nombre      string `json:"nombre"`
	Codigo      string `json:"codigo_seccion"`
	Descripcion string `json:"descripcion,omitempty"`
	Orden       int    `json:"orden"`
	Stickers    []sticker
}

type sticker struct {
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
	Orden  int    `json:"orden"`
}

func readAlbumJSON() ([]section, error) {
	file, err := os.Open(albumPanini2026)
	if err != nil {
		return []section{}, fmt.Errorf("couldn't open album .json file")
	}
	defer file.Close()

	var newAlbum []section

	if err := json.NewDecoder(file).Decode(&newAlbum); err != nil {
		return []section{}, fmt.Errorf("couldn't build album: %w", err)
	}

	return newAlbum, nil
}
