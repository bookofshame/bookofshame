-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `offence` (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	address TEXT,
    divisionId INTEGER,
    districtId INTEGER,
    upazilaId INTEGER,
    unionId INTEGER,
    dateTime TEXT,
    metadata TEXT,
    createdBy INTEGER,
    createdAt TEXT NOT NULL DEFAULT current_timestamp,
    updatedAt TEXT NOT NULL DEFAULT current_timestamp,

    FOREIGN KEY (createdBy) REFERENCES user(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (divisionId) REFERENCES location_division(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (districtId) REFERENCES location_district(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (upazilaId) REFERENCES location_upazila(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (unionId) REFERENCES location_union(id) ON DELETE SET NULL ON UPDATE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `offence`;
-- +goose StatementEnd
