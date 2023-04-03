package models

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/exp/maps"
)

type City struct {
	name      string
	neighbors map[Direction]string
}

func NewCity(name string, neighbors map[Direction]string) *City {
	return &City{
		name:      name,
		neighbors: neighbors,
	}
}

func (c *City) Name() string {
	return c.name
}

func (c *City) Neighbors() map[Direction]string {
	return c.neighbors
}

func (c *City) GetRandomNeighbor() *string {
	if c.neighbors == nil || len(c.neighbors) == 0 {
		return nil
	}

	directions := maps.Keys(c.neighbors)
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(directions))))
	neighbor := c.neighbors[directions[n.Int64()]]
	return &neighbor
}
