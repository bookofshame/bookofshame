package offender

import (
	"context"
	"fmt"

	"github.com/bookofshame/bookofshame/pkg/database"
	"github.com/bookofshame/bookofshame/pkg/logging"
	"go.uber.org/zap"
)

type Repository struct {
	db     *database.Sql
	logger *zap.SugaredLogger
}

func NewRepository(ctx context.Context, sql *database.Sql) Repository {
	return Repository{
		db:     sql,
		logger: logging.FromContext(ctx),
	}
}

func (r *Repository) GetAll() ([]Offender, error) {
	offenders := []Offender{}
	err := r.db.Select(&offenders, "SELECT * FROM offender")

	if err != nil {
		r.logger.Errorf("query failed to fetch offenders. error: %w", err)
		return nil, fmt.Errorf("failed to fetch offenders")
	}

	return offenders, nil
}

func (r *Repository) Get(id int) (*Offender, error) {
	offender := []Offender{}
	err := r.db.Select(&offender, "SELECT * FROM offender WHERE id=?", id)

	if err != nil {
		r.logger.Errorf("query failed to fetch offender (id: %d). error: %w", id, err)
		return nil, fmt.Errorf("failed to fetch offender")
	}

	return &offender[0], nil
}

func (r *Repository) Create(offender Offender) (int64, error) {
	res, err := r.db.Exec(`
        INSERT INTO offender (fullName, address, divisionId, districtId, upazilaId, unionId, metadata, photo) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `, offender.FullName, offender.Address, offender.DivisionId, offender.DistrictId, offender.UpazilaId, offender.UnionId, offender.Metadata, offender.Photo)

	if err != nil {
		r.logger.Errorf("query failed to create offenders. error: %w", err)
		return 0, fmt.Errorf("failed to create offenders")
	}

	id, _ := res.LastInsertId()

	return id, nil
}

func (r *Repository) AlreadyExists(username string, districtId int) (bool, error) {
	id := []int{}
	if err := r.db.Select(&id, "SELECT id FROM offender WHERE fullName=? AND districtId=? LIMIT 1", username, districtId); err != nil {
		r.logger.Errorf("query failed to fetch existing offender. error: %w", err)
		return false, fmt.Errorf("failed to fetch existing offender")
	}

	return len(id) > 0, nil
}

func (r *Repository) Delete(id int) error {
	if _, err := r.db.Exec("DELETE FROM offender WHERE id=?", id); err != nil {
		r.logger.Errorf("query failed to delete offender. error: %w", err)
		return fmt.Errorf("failed to delete offender")
	}

	return nil
}
