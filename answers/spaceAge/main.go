package main

import (
	"fmt"
)

type Planet string

const (
	earth   Planet = "Earth"
	mercury Planet = "Mercury"
	venus   Planet = "Venus"
	mars    Planet = "Mars"
	jupiter Planet = "Jupiter"
	saturn  Planet = "Saturn"
	uranus  Planet = "Uranus"
	neptune Planet = "Neptune"
)

func secondsToAge(p Planet, age float64) float64 {

	if p == earth {
		age *= 60. * 60. * 24. * 365.25
	} else if p == mercury {
		age *= 60. * 60. * 24. * 88
	} else if p == venus {
		age *= 60. * 60. * 24. * 225
	} else if p == mars {
		age *= 60. * 60. * 24. * 687
	} else if p == jupiter {
		age *= 60. * 60. * 24. * 4300
	} else if p == saturn {
		age *= 60. * 60. * 24. * 11000
	} else if p == uranus {
		age *= 60. * 60. * 24. * 31000
	} else if p == neptune {
		age *= 60. * 60. * 24. * 60200
	}
	return age
}

func main() {

	fmt.Printf("This is how old I would be on Earth in seconds %f", secondsToAge(mercury, 24))
}
