package one_the_tyranny_of_the_rocket_equation

import (
	"math"
)

func TheTyrannyOfTheRocketEquationPartOne(modules []float64) float64 {
	total := 0.0

	for _, module := range modules {
		total += math.Floor((module / 3) - 2)
	}

	return total
}

func TheTyrannyOfTheRocketEquationPartTwo(modules []float64) float64 {
	total := 0.0

	for _, module := range modules {
		more := true
		fuel := module

		for more {
			fuel = math.Floor((fuel / 3) - 2)

			if fuel <= 0 {
				more = false
				break
			}
			total += fuel
		}
	}

	return total
}
