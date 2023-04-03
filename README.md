# Alien Invasion Simulation

The risk is imminent we must prepare. Alien invasion is a tool that simulates what will happen to our beloved cities 
during an alien invasion.

# Installation

## Runtime

You will need go 1.19 to be installed in your machine. Please follow the instructions [here](https://go.dev/doc/install) for your environment

## Development

Dev tools required for working with this project.

Please run `make install-tools`

## Testing

You can run the test via `make test`

## Build

In order to build the project run `make build`

## Executing

After building the project a binary will be available under `./bin/` called `cli`.

## Usage

You can check the usage of the tool via `./bin/cli`

# Design choices

- For implementing the CLI the cobra library is used which is industry standard for building CLI tools.
- An adjacency list(WorldService#citiesGraph) is used to model the cities and their neighbors.
- An "index" is used on simulation for quick data retrieval of city to aliens(aliensByCity) and also alien to city(cityByAlien)
- The main simulation loop is not a universal one, cause the termination clauses are delegated to aliens. The reason behind this is that this allows to introduce
new alien types with potentially different stamina levels with ease.
- Modularity, maintainability and separation of concerns were at the forefront of designing this thus we have the following flows:
  - Service: Is the entry point with the data layer and implements business logic around the given domain i.e aliens, world, reporting, simulation etc
  - Repository: A repository is the single entry that abstracts away access to and from the data layer.
    - Enables easy support of different data layers, file to db to network.
  - Models: Model is a domain entry representation

NOTE: Simulation output is not deterministic for two reasons:
1) Random generators
2) Iterating over dictionaries
