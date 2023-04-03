package repositories

import (
	"context"
	"testing"

	"github.com/nronas/invasion_sim/internal/models"
	"github.com/stretchr/testify/suite"
)

type AliensMemImplTestSuite struct {
	ctx context.Context

	aliensMemImpl *aliensMemImpl
	alien         *models.Alien
	otherAlien    *models.Alien

	suite.Suite
}

func (s *AliensMemImplTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.aliensMemImpl = NewAlienMemImpl()

	s.alien = models.NewAlien("ET", 100, 100)
	s.otherAlien = models.NewAlien("Predator", 100, 100)
}

func (s *AliensMemImplTestSuite) TestCreateAlien() {
	alien, err := s.aliensMemImpl.Create(s.ctx, s.alien)
	s.Require().NoError(err, "should not error when creating an alien")

	s.Assert().Equal(uint64(1), alien.ID, "should assign the correct ID upon creation")
}

func (s *AliensMemImplTestSuite) TestCreateConsecutiveAliens() {
	alien, err := s.aliensMemImpl.Create(s.ctx, s.alien)
	s.Require().NoError(err, "should not error when creating an alien")

	otherAlien, err := s.aliensMemImpl.Create(s.ctx, s.otherAlien)
	s.Require().NoError(err, "should not error when creating an alien")

	s.Assert().Equal(uint64(1), alien.ID, "should assign the correct ID upon creation")
	s.Assert().Equal(uint64(2), otherAlien.ID, "should assign the correct ID upon creation")
}

func TestAliensMemImplTestSuite(t *testing.T) {
	suite.Run(t, &AliensMemImplTestSuite{})
}
