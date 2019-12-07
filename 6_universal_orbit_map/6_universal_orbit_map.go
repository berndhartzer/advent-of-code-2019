package six_universal_orbit_map

import (
	"strings"
)

type Planet struct {
	name   string
	orbits []*Planet
}

func UniversalOrbitMapPartOne(orbits []string) int {
	planets := make(map[string]*Planet)

	for _, orbit := range orbits {
		split := strings.Split(orbit, ")")
		planetName := split[0]
		orbiterName := split[1]

		planet, ok := planets[planetName]
		if !ok {
			planet = &Planet{
				name: planetName,
			}
		}

		orbiter, ok := planets[orbiterName]
		if !ok {
			orbiter = &Planet{
				name: orbiterName,
			}
		}

		orbiter.orbits = append(orbiter.orbits, planet)

		planets[planetName] = planet
		planets[orbiterName] = orbiter
	}

	total := 0

	for _, planet := range planets {
		for len(planet.orbits) > 0 {
			total += 1
			planet = planet.orbits[0]
		}
	}

	return total
}
