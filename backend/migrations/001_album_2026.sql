-- +goose Up
CREATE TABLE album_mundial_2026 (
    codigo VARCHAR(8) PRIMARY KEY,
    nombre TEXT,
    seccion TEXT,
    orden_seccion INTEGER
);

-- +goose Down
DROP TABLE album_mundial_2026;

