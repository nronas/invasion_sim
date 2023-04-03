package models

type Alien struct {
	name    string
	ID      uint64
	stamina uint
	power   uint
	health  uint
	trapped bool
}

func NewAlien(name string, stamina uint, power uint) *Alien {
	return &Alien{name: name, stamina: stamina, power: power, health: 100, trapped: false}
}

func (a *Alien) Name() string {
	return a.name
}

func (a *Alien) HasStamina() bool {
	return a.stamina > 0
}

func (a *Alien) DepleteStamina(step uint) {
	a.stamina -= step
}

func (a *Alien) IsDead() bool {
	return a.health == 0
}

func (a *Alien) IsTrapped() bool {
	return a.trapped
}

func (a *Alien) Trapped(trapped bool) {
	a.trapped = trapped
}

func (a *Alien) Dead() {
	a.health = 0
}
