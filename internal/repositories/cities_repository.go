package repositories

import (
	"context"

	"github.com/nronas/invasion_sim/internal/models"
)

//go:generate mockery --inpackage --name CitiesRepository --with-expecter
type CitiesRepository interface {
	GetAll(ctx context.Context) ([]*models.City, error)
}
