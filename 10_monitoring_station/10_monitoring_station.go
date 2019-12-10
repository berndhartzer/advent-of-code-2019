package ten_monitoring_station

import (
	"fmt"
	// "math" // atan2
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

	for thisLoc, thisAsteroid := range asteroidMap {
		if !thisAsteroid {
			continue
		}
		fmt.Println("asteroid", thisLoc, thisAsteroid)

		for otherLoc, otherAsteroid := range asteroidMap {
			if thisLoc == otherLoc || !otherAsteroid {
				continue
			}
			fmt.Println(otherLoc, otherAsteroid)
		}



	}

	return 0
}
