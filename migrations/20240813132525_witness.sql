-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `witness` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    genderId INTEGER,
    description TEXT NOT NULL,
    contact TEXT,
    offenceId INTEGER  NOT NULL,

    FOREIGN KEY (genderId) REFERENCES gender(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (offenceId) REFERENCES offence(id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE INDEX IF NOT EXISTS `witness_offenceId_idx` ON `witness` (`offenceId`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `witness`;
-- +goose StatementEnd
