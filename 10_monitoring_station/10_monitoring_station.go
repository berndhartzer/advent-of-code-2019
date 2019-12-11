package ten_monitoring_station

import (
	"math"
	"fmt"
	"sort"
)

type Coord struct {
	X, Y int
}

func MonitoringStationPartOne(asteroids []string) int {
	asteroidMap := make(map[Coord]bool)

	for y, row := range asteroids {
		for x, loc := range row {
			asteroidMap[Coord{x, y}] = (loc == '#')
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

type CoordTwo struct {
	X, Y, Distance int
}

func MonitoringStationPartTwo(asteroids []string) int {
	asteroidMap := make(map[CoordTwo]bool)

	for y, row := range asteroids {
		for x, loc := range row {
			asteroidMap[CoordTwo{X: x, Y: y}] = (loc == '#')
		}
	}

	highest := 0
	asteroidBase := CoordTwo{}
	targets := make(map[float64][]CoordTwo)
	for thisLoc, thisAsteroid := range asteroidMap {
		if !thisAsteroid {
			continue
		}
		// fmt.Println("asteroid", thisLoc)

		visible := make(map[float64][]CoordTwo)
		// visibleCount := 0

		for otherLoc, otherAsteroid := range asteroidMap {
			if thisLoc == otherLoc || !otherAsteroid {
				continue
			}

			diff := CoordTwo{
				X: otherLoc.X - thisLoc.X,
				Y: otherLoc.Y - thisLoc.Y,
			}
			distance := Abs(diff.X) + Abs(diff.Y)
			angle := ((math.Atan2(float64(diff.Y), float64(diff.X)) * 180) / math.Pi) + 180
			// fmt.Println(diff, distance, angle)

			visible[angle] = append(visible[angle], CoordTwo{otherLoc.X, otherLoc.Y, distance})
		}
		// fmt.Println(visible)

		if len(visible) > highest {
			highest = len(visible)
			asteroidBase = thisLoc
			targets = make(map[float64][]CoordTwo)
			for key, value := range visible {
				targets[key] = value
			}
		}
	}

	// fmt.Println(targets)
	sorted := []float64{}
	for angle := range targets {
		sorted = append(sorted, angle)
	}
	sort.Float64s(sorted)
	fmt.Println(sorted, len(sorted))

	fmt.Println(sorted[0], targets[sorted[0]])
	fmt.Println(sorted[209], targets[sorted[209]])

	fmt.Println(asteroidBase)
	return 0
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
