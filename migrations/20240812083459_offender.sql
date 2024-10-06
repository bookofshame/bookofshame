-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `offender` (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    fullName TEXT NOT NULL,
    address TEXT,
    genderId INTEGER,
    divisionId INTEGER,
    districtId INTEGER,
    upazilaId INTEGER,
    unionId INTEGER,
    isOrganization BOOLEAN DEFAULT 0 CHECK (isOrganization IN (0,1)),
    isEnabler BOOLEAN DEFAULT 0 CHECK (isEnabler IN (0,1)),
    isPerpetrator BOOLEAN DEFAULT 0 CHECK (isPerpetrator IN (0,1)),
    photo TEXT,
    metadata TEXT,
    createdBy INTEGER,
    createdAt TEXT NOT NULL DEFAULT current_timestamp,

    FOREIGN KEY (createdBy) REFERENCES user(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (genderId) REFERENCES gender(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (divisionId) REFERENCES location_division(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (districtId) REFERENCES location_district(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (upazilaId) REFERENCES location_upazila(id) ON DELETE SET NULL ON UPDATE NO ACTION,
    FOREIGN KEY (unionId) REFERENCES location_union(id) ON DELETE SET NULL ON UPDATE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `offender`;
-- +goose StatementEnd
