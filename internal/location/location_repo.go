package location

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

func (r *Repository) GetDivisions() ([]Division, error) {
	divisions := []Division{}

	if err := r.db.Select(&divisions, "SELECT * FROM `division`"); err != nil {
		r.logger.Errorf("query error: %w", err)
		return nil, fmt.Errorf("failed to fetch divisions")
	}

	return divisions, nil
}

func (r *Repository) GetDistricts(divisionId int) ([]District, error) {
	districts := []District{}
	var err error

	if divisionId != 0 {
		err = r.db.Select(&districts, "SELECT * FROM `district` WHERE divisionId=?", divisionId)
	} else {
		err = r.db.Select(&districts, "SELECT * FROM `district`")
	}

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return nil, fmt.Errorf("failed to fetch districts")
	}

	return districts, nil
}

func (r *Repository) GetUpazilas(districtId int) ([]Upazila, error) {
	upazilas := []Upazila{}
	var err error

	if districtId != 0 {
		err = r.db.Select(&upazilas, "SELECT * FROM `upazila` WHERE districtId=?", districtId)
	} else {
		err = r.db.Select(&upazilas, "SELECT * FROM `upazila`")
	}

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return nil, fmt.Errorf("failed to fetch upazilas")
	}

	return upazilas, nil
}

func (r *Repository) GetUnions(upazilaId int) ([]Union, error) {
	unions := []Union{}
	var err error

	if upazilaId != 0 {
		err = r.db.Select(&unions, "SELECT * FROM `union` WHERE upazilaId=?", upazilaId)
	} else {
		err = r.db.Select(&unions, "SELECT * FROM `union`")
	}

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return nil, fmt.Errorf("failed to fetch unions")
	}

	return unions, nil
}
