package ten_monitoring_station

import (
	"math"
	"sort"
)

type Coord struct {
	X, Y, Distance int
}

func MonitoringStationPartOne(asteroids []string) int {
	asteroidMap := make(map[Coord]bool)

	for y, row := range asteroids {
		for x, loc := range row {
			asteroidMap[Coord{X: x, Y: y}] = (loc == '#')
		}
	}

	highest := 0
	for thisLoc, thisAsteroid := range asteroidMap {
		if !thisAsteroid {
			continue
		}

		visible := make(map[float64]bool)
		visibleCount := 0

		for otherLoc, otherAsteroid := range asteroidMap {
			if thisLoc == otherLoc || !otherAsteroid {
				continue
			}

			diff := Coord{
				X: otherLoc.X - thisLoc.X,
				Y: otherLoc.Y - thisLoc.Y,
			}
			angle := ((math.Atan2(float64(diff.Y), float64(diff.X)) * 180) / math.Pi) + 180

			_, ok := visible[angle]
			if !ok {
				visibleCount += 1
				visible[angle] = true
			}
		}

		if visibleCount > highest {
			highest = visibleCount
		}
	}

	return highest
}

func MonitoringStationPartTwo(asteroids []string, vaporiseLimit int) int {
	asteroidMap := make(map[Coord]bool)

	for y, row := range asteroids {
		for x, loc := range row {
			asteroidMap[Coord{X: x, Y: y}] = (loc == '#')
		}
	}

	highest := 0
	targets := make(map[float64][]Coord)

	for thisLoc, thisAsteroid := range asteroidMap {
		if !thisAsteroid {
			continue
		}

		visible := make(map[float64][]Coord)

		for otherLoc, otherAsteroid := range asteroidMap {
			if thisLoc == otherLoc || !otherAsteroid {
				continue
			}

			diff := Coord{
				X: otherLoc.X - thisLoc.X,
				Y: otherLoc.Y - thisLoc.Y,
			}
			distance := Abs(diff.X) + Abs(diff.Y)
			angle := math.Atan2(float64(diff.Y), float64(diff.X)) * (180 / math.Pi)

			// rudimentary angle correction for easier sorting
			if angle < -90 {
				angle += 450
			} else {
				angle += 90
			}

			visible[angle] = append(visible[angle], Coord{otherLoc.X, otherLoc.Y, distance})
		}

		if len(visible) > highest {
			highest = len(visible)
			targets = make(map[float64][]Coord)
			for key, value := range visible {
				targets[key] = value
			}
		}
	}

	sorted := []float64{}
	for angle := range targets {
		sorted = append(sorted, angle)
	}
	sort.Float64s(sorted)

	twoHundredth := 0
	vaporised := 0
	i := 0
	for vaporised < vaporiseLimit {
		angle := sorted[i]

		if len(targets[angle]) > 0 {
			target := 0

			if len(targets[angle]) != 1 {
				closest := 0
				for j, loc := range targets[angle] {
					if loc.Distance < closest || closest == 0 {
						closest = loc.Distance
						target = j
					}
				}
			}

			twoHundredth = (targets[angle][target].X * 100) + targets[angle][target].Y
			vaporised += 1
		}

		i += 1
		if i >= len(targets) {
			i = 0
		}
	}

	return twoHundredth
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
