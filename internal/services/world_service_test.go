package services

import (
	"context"
	"errors"
	"testing"

	"github.com/nronas/invasion_sim/internal/models"
	"github.com/nronas/invasion_sim/internal/repositories"
	"github.com/stretchr/testify/suite"
)

var ErrRepo = errors.New("repo failure")

type WorldServiceTestSuite struct {
	ctx context.Context

	citiesRepository *repositories.MockCitiesRepository
	service          *worldService
	cities           []*models.City

	suite.Suite
}

func (s *WorldServiceTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.cities = []*models.City{
		models.NewCity(
			"Athens",
			map[models.Direction]string{
				models.DirectionEast: "Ankara",
				models.DirectionWest: "Rome",
			},
		),
		models.NewCity(
			"Rome",
			map[models.Direction]string{
				models.DirectionWest: "Athens",
			},
		),
	}
	s.citiesRepository = &repositories.MockCitiesRepository{}
	s.citiesRepository.EXPECT().GetAll(s.ctx).Return(s.cities, nil)
	service, err := NewWorldService(s.ctx, s.citiesRepository)
	s.Require().NoError(err)
	s.service = service
}

func (s *WorldServiceTestSuite) AfterTest() {
	s.citiesRepository.AssertExpectations(s.T())
}

func (s *WorldServiceTestSuite) TestTotalCities() {
	s.Assert().Equal(len(s.cities), s.service.TotalCities())
}

func (s *WorldServiceTestSuite) TestGetCityExists() {
	cityToFetch := s.cities[0]
	cityFetched, err := s.service.GetCity(cityToFetch.Name())
	s.Require().NoError(err, "should not error fetching a city")
	s.Assert().Equal(cityFetched, cityToFetch)
}

func (s *WorldServiceTestSuite) TestGetCityDoesNotExists() {
	cityFetched, err := s.service.GetCity("Asgard")
	s.Require().ErrorIs(err, ErrUnknownCity, "should error fetching a city")
	s.Assert().Nil(cityFetched)
}

func (s *WorldServiceTestSuite) TestDestroyCity() {
	cityToDestroy := s.cities[0]
	s.service.DestroyCity(cityToDestroy)

	cityFetched, err := s.service.GetCity(cityToDestroy.Name())
	s.Require().ErrorIs(err, ErrUnknownCity, "should error fetching a city")
	s.Assert().Nil(cityFetched)
}

func TestWorldServiceTestSuite(t *testing.T) {
	suite.Run(t, &WorldServiceTestSuite{})
}
