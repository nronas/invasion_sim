package main

import (
	"context"
	"log"

	"github.com/nronas/invasion_sim/internal/cli"
)

func main() {
	ctx := context.Background()
	cli, err := cli.NewCLI(ctx)
	if err != nil {
		log.Fatalf(" initializing command error: %+v", err)
	}

	if err := cli.Execute(); err != nil {
		log.Fatalf("Probably the aliens have killed us all. ERROR: %+v", err)
	}
}
