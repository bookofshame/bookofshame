package gender

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

func (r *Repository) GetAll() ([]Gender, error) {
	var genders []Gender
	err := r.db.Select(&genders, "SELECT * FROM gender")

	if err != nil {
		r.logger.Errorf("query error: %w", err)
		return nil, fmt.Errorf("failed to fetch genders")
	}

	return genders, nil
}
