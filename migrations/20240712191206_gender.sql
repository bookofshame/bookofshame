-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `gender` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    bnName TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `gender`;
-- +goose StatementEnd
