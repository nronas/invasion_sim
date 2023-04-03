package services

import (
	"context"
	"testing"

	"github.com/nronas/invasion_sim/internal/models"
	"github.com/nronas/invasion_sim/internal/repositories"
	"github.com/stretchr/testify/suite"
)

type AlienServiceTestSuite struct {
	ctx context.Context

	alienRepository *repositories.MockAliensRepository
	service         *alienService
	alien           *models.Alien

	suite.Suite
}

func (s *AlienServiceTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.alienRepository = &repositories.MockAliensRepository{}
	s.service = NewAlienService(s.alienRepository)
	s.alien = models.NewAlien("ET", 100, 100)
}

func (s *AlienServiceTestSuite) AfterTest() {
	s.alienRepository.AssertExpectations(s.T())
}

func (s *AlienServiceTestSuite) TestCreateAlienSucceeds() {
	s.alienRepository.EXPECT().Create(s.ctx, s.alien).Return(s.alien, nil)
	_, err := s.service.CreateAlien(s.ctx, s.alien)
	s.Assert().NoError(err, "should not error when creating new alien")
}

func (s *AlienServiceTestSuite) TestCreateAlienFails() {
	s.alienRepository.EXPECT().Create(s.ctx, s.alien).Return(nil, ErrRepo)
	_, err := s.service.CreateAlien(s.ctx, s.alien)
	s.Assert().ErrorIs(err, ErrRepo, "should bubble up the error when creating new alien fails")
}

func TestAlienServiceTestSuite(t *testing.T) {
	suite.Run(t, &AlienServiceTestSuite{})
}
