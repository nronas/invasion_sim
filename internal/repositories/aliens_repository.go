package repositories

import (
	"context"

	"github.com/nronas/invasion_sim/internal/models"
)

//go:generate mockery --inpackage --name AliensRepository --with-expecter
type AliensRepository interface {
	Create(ctx context.Context, alien *models.Alien) (*models.Alien, error)
}
