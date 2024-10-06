-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `union` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    bnName TEXT NOT NULL,
    upazilaId INTEGER NOT NULL,
    FOREIGN KEY (upazilaId) REFERENCES upazila(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `union`;
-- +goose StatementEnd
