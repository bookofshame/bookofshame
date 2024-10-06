-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `upazila` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    bnName TEXT NOT NULL,
    districtId INTEGER NOT NULL,
    FOREIGN KEY (districtId) REFERENCES district(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `upazila`;
-- +goose StatementEnd
