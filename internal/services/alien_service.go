package services

import (
	"context"

	"github.com/nronas/invasion_sim/internal/models"
	"github.com/nronas/invasion_sim/internal/repositories"
)

type alienService struct {
	aliensRepository repositories.AliensRepository
}

// NewAlienService creates a new service that acts as the entry point for all interaction with aliens.
func NewAlienService(aliensRepository repositories.AliensRepository) *alienService {
	return &alienService{aliensRepository: aliensRepository}
}

func (as *alienService) CreateAlien(ctx context.Context, alien *models.Alien) (*models.Alien, error) {
	return as.aliensRepository.Create(ctx, alien)
}
