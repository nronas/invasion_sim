package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/nronas/invasion_sim/internal/models"
	"github.com/nronas/invasion_sim/internal/repositories"
	"golang.org/x/exp/maps"
)

var ErrUnknownCity = errors.New("unknown city to this worldService")

type worldService struct {
	citiesRepository repositories.CitiesRepository
	citiesGraph      map[string]*models.City
}

func NewWorldService(ctx context.Context, citiesRepository repositories.CitiesRepository) (*worldService, error) {
	world := &worldService{citiesRepository: citiesRepository}
	if err := world.computeCitiesGraph(ctx); err != nil {
		return nil, err
	}

	return world, nil
}

func (w *worldService) TotalCities() int {
	return len(w.citiesGraph)
}

func (w *worldService) GetRandomCity() *models.City {
	cityNames := maps.Keys(w.citiesGraph)
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(cityNames))))
	randomCityName := cityNames[n.Int64()]
	return w.citiesGraph[randomCityName]
}

func (w *worldService) WorldToString(_ context.Context) string {
	line := "\n"
	for _, city := range w.citiesGraph {
		line += city.Name()
		for direction, neighborCityName := range city.Neighbors() {
			line += fmt.Sprintf(" %s=%s", direction, neighborCityName)
		}
		line += "\n"
	}

	return line
}

func (w *worldService) GetCity(cityName string) (*models.City, error) {
	if city, ok := w.citiesGraph[cityName]; ok {
		return city, nil
	}

	return nil, ErrUnknownCity
}

func (w *worldService) DestroyCity(city *models.City) {
	if city == nil {
		return
	}

	for _, otherCity := range w.citiesGraph {
		if otherCity.Neighbors() == nil {
			continue
		}

		for direction, neighborName := range otherCity.Neighbors() {
			if neighborName == city.Name() {
				delete(otherCity.Neighbors(), direction)
			}
		}
	}
	delete(w.citiesGraph, city.Name())
}

func (w *worldService) computeCitiesGraph(ctx context.Context) error {
	cities, err := w.citiesRepository.GetAll(ctx)
	if err != nil {
		return err
	}

	if w.citiesGraph == nil {
		w.citiesGraph = make(map[string]*models.City)
	}

	for _, city := range cities {
		w.citiesGraph[city.Name()] = city
	}

	return nil
}
