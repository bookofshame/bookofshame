-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `evidence` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT NOT NULL,
    metadata TEXT,
    isCollected BOOLEAN DEFAULT 0 CHECK (isCollected IN (0,1)),
    isQuestionable BOOLEAN DEFAULT 0 CHECK (isQuestionable IN (0,1)),
    questionableCount INTEGER NOT NULL DEFAULT 0,
    offenceId INTEGER NOT NULL,
    createdBy INTEGER,
    createdAt TEXT NOT NULL DEFAULT current_timestamp,

    FOREIGN KEY (offenceId) REFERENCES offence(id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (createdBy) REFERENCES user(id) ON DELETE SET NULL ON UPDATE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `evidence`;
-- +goose StatementEnd
