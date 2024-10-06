-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `offender_offence` (
    offenderId INTEGER  NOT NULL,
    offenceId INTEGER  NOT NULL,
    
    PRIMARY KEY (offenderId, offenceId),
    FOREIGN KEY (offenderId) REFERENCES offender(id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (offenceId) REFERENCES offence(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `offender_offence`;
-- +goose StatementEnd
