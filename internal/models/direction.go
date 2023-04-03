package models

type Direction string

const (
	DirectionEast  = Direction("east")
	DirectionWest  = Direction("west")
	DirectionNorth = Direction("north")
	DirectionSouth = Direction("south")
)

func (d Direction) Valid() bool {
	return DirectionEast == d || DirectionWest == d || DirectionNorth == d || DirectionSouth == d
}
