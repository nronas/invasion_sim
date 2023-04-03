package cli

import (
	"context"

	"github.com/nronas/invasion_sim/internal/services"
	"github.com/spf13/cobra"
)

// NewCLI creates a new root command for `Invasion Simulator`.
func NewCLI(ctx context.Context) (*cobra.Command, error) {
	cobra.EnableCommandSorting = false

	c := &cobra.Command{
		Use:           "invasion",
		Short:         "An invasion simulator tool",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	c.AddCommand(NewSimulation(ctx, services.NewSimulationService()))

	return c, nil
}
