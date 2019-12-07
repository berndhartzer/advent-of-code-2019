package six_universal_orbit_map

import (
	"strings"
)

type Planet struct {
	name   string
	orbits []*Planet
}

type Transfer struct {
	you   int
	santa int
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

func UniversalOrbitMapPartTwo(orbits []string) int {
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

	transfers := make(map[string]*Transfer)

	for _, planet := range planets["YOU"].orbits {
		steps := 0
		for len(planet.orbits) > 0 {
			transfer, ok := transfers[planet.name]
			if !ok {
				transfer = &Transfer{
					you: steps,
				}
			}
			transfer.you = steps
			transfers[planet.name] = transfer
			planet = planet.orbits[0]

			steps += 1
		}
	}

	lowest := 0

	for _, planet := range planets["SAN"].orbits {
		steps := 0
		for len(planet.orbits) > 0 {
			transfer, ok := transfers[planet.name]
			if !ok {
				transfer = &Transfer{
					santa: steps,
				}
			}
			transfer.santa = steps
			transfers[planet.name] = transfer
			planet = planet.orbits[0]

			if transfer.you > 0 && transfer.santa > 0 {
				sum := transfer.you + transfer.santa

				if sum < lowest || lowest == 0 {
					lowest = sum
				}
			}

			steps += 1
		}
	}

	return lowest
}
