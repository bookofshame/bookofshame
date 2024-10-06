-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `division` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    bnName TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `division`;
-- +goose StatementEnd
