package factory

import "fmt"

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{name: "AK47 gun", power: 4},
	}
}

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{name: "Maverick gun", power: 5},
	}
}

// gun factory
func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "maverick" {
		return newMaverick(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
