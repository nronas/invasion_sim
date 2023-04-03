package repositories

import (
	"context"
	"strings"
	"testing"

	"github.com/nronas/invasion_sim/internal/models"
	"github.com/stretchr/testify/suite"
)

type CitiesIOImplTestSuite struct {
	ctx context.Context

	suite.Suite
}

func (s *CitiesIOImplTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func (s *CitiesIOImplTestSuite) TestParseReader() {
	scenarios := []struct {
		scenario string
		input    string
		err      error
	}{
		{
			"single valid line",
			"city north=otherCity",
			nil,
		},
		{
			"single valid line no neighbors",
			"city",
			nil,
		},
		{
			"single city with invalid direction",
			"city n0rth=otherCity",
			ErrInvalidSource,
		},
		{
			"single line invalid format",
			"city     north=>otherCity",
			ErrInvalidSource,
		},
		{
			"multiline valid",
			"city north=otherCity\notherCity south=city",
			nil,
		},
	}

	for _, scenario := range scenarios {
		s.T().Run(scenario.scenario, func(t *testing.T) {
			reader := strings.NewReader(scenario.input)
			_, err := NewCitiesIOImpl(reader)
			if scenario.err == nil {
				s.Assert().NoError(err)
			} else {
				s.Assert().ErrorIs(err, scenario.err)
			}
		})
	}
}

func (s *CitiesIOImplTestSuite) TestGetAll() {
	reader := strings.NewReader("city north=otherCity\notherCity south=city")
	impl, err := NewCitiesIOImpl(reader)
	s.Require().NoError(err, "should not error when initializing impl from valid source")

	actualCities, err := impl.GetAll(s.ctx)
	s.Require().NoError(err, "should not error when fetching all cities")

	expectedCities := []*models.City{
		models.NewCity(
			"city",
			map[models.Direction]string{
				"north": "otherCity",
			},
		),
		models.NewCity(
			"otherCity",
			map[models.Direction]string{
				"south": "city",
			},
		),
	}

	s.Assert().Equal(expectedCities, actualCities)
}

func TestCitiesIOImplTestSuite(t *testing.T) {
	suite.Run(t, &CitiesIOImplTestSuite{})
}
