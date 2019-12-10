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
		fmt.Println(thisLoc, thisAsteroid)


	}

	return 0
}
