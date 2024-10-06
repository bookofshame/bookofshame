-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `offence_witness` (
    offenceId INTEGER  NOT NULL,
    witnessId INTEGER  NOT NULL,
    
    PRIMARY KEY (offenceId, witnessId),
    FOREIGN KEY (offenceId) REFERENCES offence(id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (witnessId) REFERENCES witness(id) ON DELETE CASCADE ON UPDATE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `offence_witness`;
-- +goose StatementEnd
