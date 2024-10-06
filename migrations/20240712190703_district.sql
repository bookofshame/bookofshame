-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `district` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    bnName TEXT NOT NULL,
    divisionId INTEGER NOT NULL,
    lat REAL DEFAULT NULL,
    long REAL DEFAULT NULL,
    FOREIGN KEY (divisionId) REFERENCES division(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `district`;
-- +goose StatementEnd
