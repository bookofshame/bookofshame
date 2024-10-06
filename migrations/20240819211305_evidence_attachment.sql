-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `evidence_attachment` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    fileUrl TEXT NOT NULL,
    thumbnailUrl TEXT,
    contentType TEXT NOT NULL,
    evidenceId INTEGER NOT NULL,
    createdBy INTEGER,
    createdAt TEXT NOT NULL DEFAULT current_timestamp,

    FOREIGN KEY (evidenceId) REFERENCES evidence(id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (createdBy) REFERENCES user(id) ON DELETE SET NULL ON UPDATE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `evidence_attachment`;
-- +goose StatementEnd
