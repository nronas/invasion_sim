package repositories

import (
	"context"
	"sync/atomic"

	"github.com/nronas/invasion_sim/internal/models"
)

var _ AliensRepository = (*aliensMemImpl)(nil)

type aliensMemImpl struct {
	aliens    []*models.Alien
	idCounter uint64
}

func NewAlienMemImpl() *aliensMemImpl {
	return &aliensMemImpl{}
}

func (ami *aliensMemImpl) Create(_ context.Context, alien *models.Alien) (*models.Alien, error) {
	alien.ID = atomic.AddUint64(&ami.idCounter, 1)
	ami.aliens = append(ami.aliens, alien)

	return alien, nil
}
