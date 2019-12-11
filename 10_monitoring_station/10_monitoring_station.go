package ten_monitoring_station

import (
	"math"
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
