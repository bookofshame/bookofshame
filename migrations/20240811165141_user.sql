-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    fullName TEXT NOT NULL,
    genderId INTEGER,
    address TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT,
    locale TEXT NOT NULL DEFAULT 'bn',
    activationCode TEXT,
    isActive BOOLEAN DEFAULT 0 CHECK (isActive IN (0,1)),
    isAdmin BOOLEAN DEFAULT 0 CHECK (isAdmin IN (0,1)),
    password TEXT NOT NULL,
    createdAt TEXT NOT NULL DEFAULT current_timestamp,

    FOREIGN KEY (genderId) REFERENCES gender(id) ON DELETE SET NULL ON UPDATE NO ACTION
);

CREATE UNIQUE INDEX IF NOT EXISTS `user_phone_idx` ON `user` (`email`);
CREATE UNIQUE INDEX IF NOT EXISTS `user_activation_code_idx` ON `user` (`activationCode`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user`;
-- +goose StatementEnd