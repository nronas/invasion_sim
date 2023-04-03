package services

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nronas/invasion_sim/internal/models"
)

const (
	AliensNeededToCauseAFight = 2
	aliensTrappedStat         = "trapped-aliens"
	aliensKilledStat          = "killed-aliens"
	citiesDestroyedStat       = "destroyed-cities"
)

type SimulationService struct {
	ctx             context.Context
	name            string
	worldService    *worldService
	alienService    *alienService
	reportService   *reportService
	alienCount      uint
	alienMaxStamina uint

	aliens       map[uint64]*models.Alien
	aliensByCity map[string][]*models.Alien
	cityByAlien  map[uint64]*models.City
	stats        map[string]uint
}

func NewSimulationService() *SimulationService {
	return &SimulationService{
		aliens:       make(map[uint64]*models.Alien),
		aliensByCity: make(map[string][]*models.Alien),
		cityByAlien:  make(map[uint64]*models.City),
		stats:        make(map[string]uint),
	}
}

// WithName updates the name of the current simulation
func (s *SimulationService) WithName(name string) *SimulationService {
	s.name = name
	return s
}

func (s *SimulationService) WithAlienCount(alienCount uint) *SimulationService {
	s.alienCount = alienCount
	return s
}

func (s *SimulationService) WithAlienMaxStamina(alienStamina uint) *SimulationService {
	s.alienMaxStamina = alienStamina
	return s
}

func (s *SimulationService) WithReportService(reportService *reportService) *SimulationService {
	s.reportService = reportService
	return s
}

// WithAlienService updates the aliens of the current simulation
func (s *SimulationService) WithAlienService(alienService *alienService) *SimulationService {
	s.alienService = alienService
	return s
}

// WithWorldService updates the current simulation with a worldService object
func (s *SimulationService) WithWorldService(worldService *worldService) *SimulationService {
	s.worldService = worldService
	return s
}

// Run executes an alien invasion simulation.
// High level algo implemented:
//  1. Randomly place aliens onto the worldService
//     Note: At this step we can have more than AliensNeededToCauseAFight placed on a given city, because
//     the placement strategy is random.
//  2. Check whether a city needs to be destroyed
//     Note: I choose to perform collision check prior to alien movement, in order to clear up the worldService from
//     cities that had more than AliensNeededToCauseAFight after the initial placement.
//     2.1 When a city is destroyed, the killedAlien count will be increased with the number of alien casualties.
//  3. If alien is not dead after the collision check, move the alien into a random neighbor of its current city.
//     Note: When alien cannot move, cause there are no neighbors available, in cases whether a city is a dead-end from
//     the beginning or ended up with no neighbors after it's neighbors were destroyed, a message will be displayed that
//     the alien is trapped.
//  4. Deplete 1 unit from Alien's stamina
//     4.1 Check if Alien's stamina is zero and increase the exhausted counter.
//
// All the above are in a loop that is governed from the killed alien count and the exhausted alien count.
func (s *SimulationService) Run(ctx context.Context) error {
	s.ctx = ctx

	if err := s.createAliens(); err != nil {
		return err
	}

	s.reportHeader()

	if len(s.aliens) == 0 {
		_ = s.reportService.Report(s.ctx, "😅 False alarm, no aliens are attacking...")
		s.reportWorldState()
		return nil
	}

	s.randomizeAliens()
	shouldRun := true

	start := time.Now()
	day := 1

	exhaustedAliens := 0
	killedAliens := 0
	stopReason := ""

	for shouldRun {
		_ = s.reportService.Report(s.ctx, fmt.Sprintf("== DAY %d ==", day))
		for _, alien := range s.aliens {
			// Check if alien is dead or does not have remaining moves to make
			// Trapped aliens will continue to try to move and output the trapped message
			if alien.IsDead() || !alien.HasStamina() {
				continue
			}

			currentAlienCity := s.cityByAlien[alien.ID]
			if s.shouldDestroyCity(currentAlienCity) {
				casualties := s.performCityFight(currentAlienCity)
				s.reportFallOfACity(currentAlienCity, casualties)
				killedAliens += len(casualties)
				continue
			}

			if err := s.moveAlien(alien); err != nil {
				return fmt.Errorf("there has been a glitch in space-time %w", err)
			}

			alien.DepleteStamina(1)
			if !alien.HasStamina() {
				exhaustedAliens += 1
			}
		}
		if killedAliens == len(s.aliens) {
			stopReason = "all aliens are killed. Time to repair our cities!"
			shouldRun = false
		} else if exhaustedAliens+killedAliens == len(s.aliens) {
			stopReason = "all alive aliens are exhausted"
			shouldRun = false
		}
		day += 1
	}

	_ = s.reportService.Report(s.ctx, fmt.Sprintf("=== Simulation %s 👽 Finished. Took: %s. Termination Reason: %s ===", s.name, time.Since(start), stopReason))
	s.reportStats()
	_ = s.reportService.Report(s.ctx, "")
	s.reportWorldState()

	return nil
}

func (s *SimulationService) reportWorldState() {
	_ = s.reportService.Report(s.ctx, "World Map 🌎 after the invasion.")
	_ = s.reportService.Report(s.ctx, s.worldService.WorldToString(s.ctx))
}

func (s *SimulationService) reportHeader() {
	header := fmt.Sprintf("=== Simulation %s 👽 Started. Let the hope die last. ===", s.name)
	_ = s.reportService.Report(s.ctx, strings.Repeat("=", len(header)))
	_ = s.reportService.Report(s.ctx, header)
	_ = s.reportService.Report(s.ctx, strings.Repeat("=", len(header)))
	_ = s.reportService.Report(s.ctx, "Simulation Parameters are as follows:")
	_ = s.reportService.Report(s.ctx, fmt.Sprintf("world Cities: %d", s.worldService.TotalCities()))
	_ = s.reportService.Report(s.ctx, fmt.Sprintf("Aliens: %d", len(s.aliens)))
	_ = s.reportService.Report(s.ctx, "")
}

func (s *SimulationService) randomizeAliens() {
	_ = s.reportService.Report(s.ctx, "Aliens are approaching...")
	for _, alien := range s.aliens {
		city := s.worldService.GetRandomCity()
		s.addAlienToCity(city, alien)

		_ = s.reportService.Report(s.ctx, fmt.Sprintf("%s has landed on %s", alien.Name(), city.Name()))
	}
}

func (s *SimulationService) moveAlien(alien *models.Alien) error {
	currentAlienCity := s.cityByAlien[alien.ID]
	nextCityName := currentAlienCity.GetRandomNeighbor()

	if nextCityName == nil {
		_ = s.reportService.Report(s.ctx, fmt.Sprintf("Alien(%d) is TRAPPED, there is hope!", alien.ID))
		if !alien.IsTrapped() {
			alien.Trapped(true)
			s.stats[aliensTrappedStat] += 1
		}
		return nil
	}

	_ = s.reportService.Report(s.ctx, fmt.Sprintf("Alien(%d) moving from %s -> %s", alien.ID, currentAlienCity.Name(), *nextCityName))

	nextCity, err := s.worldService.GetCity(*nextCityName)
	if err != nil {
		return err
	}

	s.addAlienToCity(nextCity, alien)

	return nil
}

func (s *SimulationService) addAlienToCity(city *models.City, alien *models.Alien) {
	if currentAlienCity, ok := s.cityByAlien[alien.ID]; ok {
		if s.aliensByCity[currentAlienCity.Name()] == nil {
			s.aliensByCity[currentAlienCity.Name()] = []*models.Alien{}
		}

		for i, alienInCity := range s.aliensByCity[currentAlienCity.Name()] {
			if alienInCity.ID == alien.ID {
				s.aliensByCity[currentAlienCity.Name()] = append(s.aliensByCity[currentAlienCity.Name()][:i], s.aliensByCity[currentAlienCity.Name()][i+1:]...)
				break
			}
		}
	}

	s.aliensByCity[city.Name()] = append(s.aliensByCity[city.Name()], alien)
	s.cityByAlien[alien.ID] = city
}

func (s *SimulationService) shouldDestroyCity(city *models.City) bool {
	if city == nil {
		return false
	}

	return len(s.aliensByCity[city.Name()]) == AliensNeededToCauseAFight
}

func (s *SimulationService) performCityFight(city *models.City) []uint64 {
	aliens := s.aliensByCity[city.Name()]
	s.worldService.DestroyCity(city)
	s.stats[citiesDestroyedStat] += 1
	delete(s.aliensByCity, city.Name())
	casualties := make([]uint64, len(aliens))
	for i, alien := range aliens {
		delete(s.cityByAlien, alien.ID)
		alien.Dead()
		s.stats[aliensKilledStat] += 1
		if alien.IsTrapped() {
			s.stats[aliensTrappedStat] -= 1
		}
		casualties[i] = alien.ID
	}

	return casualties
}

func (s *SimulationService) reportFallOfACity(city *models.City, casualties []uint64) {
	casualtiesStr := make([]string, len(casualties))
	for i, casualty := range casualties {
		casualtiesStr[i] = strconv.FormatUint(casualty, 10)
	}
	_ = s.reportService.Report(s.ctx, fmt.Sprintf("💥 city %s has been destroyed by aliens %s 💥", city.Name(), strings.Join(casualtiesStr, " and ")))
}

func (s *SimulationService) reportStats() {
	for statName, count := range s.stats {
		_ = s.reportService.Report(s.ctx, fmt.Sprintf("%s: %d", statName, count))
	}
}

func (s *SimulationService) createAliens() error {
	for i := 1; i <= int(s.alienCount); i++ {
		alien := models.NewAlien(fmt.Sprintf("Alien(%d)", i), s.alienMaxStamina, 100)
		createdAlien, err := s.alienService.CreateAlien(s.ctx, alien)
		if err != nil {
			return err
		}

		s.aliens[alien.ID] = createdAlien
	}

	return nil
}
