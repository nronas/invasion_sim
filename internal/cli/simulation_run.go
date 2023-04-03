package cli

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/nronas/invasion_sim/internal/repositories"
	"github.com/nronas/invasion_sim/internal/services"
	"github.com/spf13/cobra"
)

const (
	worldConfigFlag            = "world-config"
	alienCountFlag             = "aliens"
	alienMaxIterationCountFlag = "max-alien-moves"

	maxAlienIterationDefaultValue = 10000
	defaultAlienCount             = 2
)

var ErrSimulation = errors.New("simulation error occurred")

func NewSimulationRun(ctx context.Context, simulation *services.SimulationService) *cobra.Command {
	c := &cobra.Command{
		Use:   "run [name]",
		Short: "Create a new account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return simulationRunHandler(ctx, simulation, cmd, args)
		},
	}

	c.Flags().String(worldConfigFlag, "", "Absolute location of the world file, that follows <city_name><space>(<direction>=<city_name>)*")
	_ = c.MarkFlagRequired(worldConfigFlag)

	c.Flags().Uint(alienCountFlag, defaultAlienCount, "Number of aliens in the simulation")
	c.Flags().Uint(alienMaxIterationCountFlag, maxAlienIterationDefaultValue, "Max number of moves that each alien can make")

	return c
}

func simulationRunHandler(ctx context.Context, simulationService *services.SimulationService, cmd *cobra.Command, args []string) error {
	if simulationService == nil {
		return fmt.Errorf("simulation cannot be initialized %w", ErrSimulation)
	}

	var (
		worldConfig, _            = cmd.Flags().GetString(worldConfigFlag)
		alienCount, _             = cmd.Flags().GetUint(alienCountFlag)
		alienMaxIterationCount, _ = cmd.Flags().GetUint(alienMaxIterationCountFlag)
	)

	file, err := os.Open(worldConfig)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	// Set up the worldService
	citiesRepository, err := repositories.NewCitiesIOImpl(file)
	if err != nil {
		return err
	}

	worldService, err := services.NewWorldService(ctx, citiesRepository)
	if err != nil {
		return err
	}

	alienRepository := repositories.NewAlienMemImpl()
	alienService := services.NewAlienService(alienRepository)

	reportService := services.NewReportService(repositories.NewStdoutReportRepository())

	// Configure the SimulationService
	simulationService.
		WithName(args[0]).
		WithAlienCount(alienCount).
		WithAlienMaxStamina(alienMaxIterationCount).
		WithAlienService(alienService).
		WithWorldService(worldService).
		WithReportService(reportService)

	// Fight...
	return simulationService.Run(ctx)
}
