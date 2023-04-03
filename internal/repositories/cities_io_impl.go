package repositories

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/nronas/invasion_sim/internal/models"
)

const (
	configTokenSeparator           = " "
	configDirectionToCitySeparator = "="
)

var ErrInvalidSource = errors.New("source is invalid error")
var _ CitiesRepository = (*citiesIOImpl)(nil)

type citiesIOImpl struct {
	cities []*models.City

	reader io.Reader
}

func NewCitiesIOImpl(reader io.Reader) (*citiesIOImpl, error) {
	cities, err := parse(reader)
	if err != nil {
		return nil, err
	}

	return &citiesIOImpl{cities: cities, reader: reader}, nil
}

func (cfi *citiesIOImpl) GetAll(_ context.Context) ([]*models.City, error) {
	return cfi.cities, nil
}

func parse(reader io.Reader) ([]*models.City, error) {
	scanner := bufio.NewScanner(reader)

	var cities []*models.City
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, configTokenSeparator)
		name := parts[0]

		neighbors := make(map[models.Direction]string)
		for _, directionToCityPart := range parts[1:] {
			neighborParts := strings.Split(directionToCityPart, configDirectionToCitySeparator)
			if len(neighborParts) != 2 {
				return nil, fmt.Errorf("invalid neighbors format %w", ErrInvalidSource)
			}
			direction := models.Direction(neighborParts[0])
			neighbor := neighborParts[1]
			if !direction.Valid() {
				return nil, fmt.Errorf("invalid direction format %s: %w", direction, ErrInvalidSource)
			}
			neighbors[direction] = neighbor
		}

		cities = append(cities, models.NewCity(name, neighbors))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cities, nil
}
