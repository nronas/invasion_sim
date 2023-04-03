package cli

import (
	"context"

	"github.com/nronas/invasion_sim/internal/services"
	"github.com/spf13/cobra"
)

func NewSimulation(ctx context.Context, simulationService *services.SimulationService) *cobra.Command {
	c := &cobra.Command{
		Use:     "simulation [command]",
		Short:   "Manages simulation capabilities of the invasion",
		Aliases: []string{"s"},
		Args:    cobra.ExactArgs(1),
	}

	c.AddCommand(NewSimulationRun(ctx, simulationService))

	return c
}
