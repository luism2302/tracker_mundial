package models

import "github.com/jackc/pgx/v5"

type Sticker struct {
	ID     string
	Equipo string
	Nombre string
}

type StickerModel struct {
	DB *pgx.Conn
}

func (m *StickerModel) Insert(code, name string, order int) error {
	return nil
}
