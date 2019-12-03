package three_crossed_wires

import (
	"fmt"
	"strings"
	"strconv"
)

type Coord struct {
	X, Y int
}

func (c *Coord) Move(direction string) {
	switch direction {
	case "U":
		c.Y += 1
	case "D":
		c.Y -= 1
	case "R":
		c.X += 1
	case "L":
		c.X -= 1
	}
}

func CrossedWiresPartOne(wires [][]string) int {
	grid := make(map[string][]int)

	for i, wire := range wires {

		coord := Coord{0, 0}
		grid[fmt.Sprintf("%d,%d", coord.X, coord.Y)] = append(grid[fmt.Sprintf("%d,%d", coord.X, coord.Y)], i)

		for _, instruction := range wire {

			split := []rune(instruction)
			distance, _ := strconv.Atoi(string(split[1:]))
			// fmt.Println(split, distance)

			for j := 0; j < distance; j++ {
				// coord.Move(split[0])
				coord.Move(string(split[0]))
				grid[fmt.Sprintf("%d,%d", coord.X, coord.Y)] = append(grid[fmt.Sprintf("%d,%d", coord.X, coord.Y)], i)
			}

		}
	}

	overlaps := []string{}
	for coord, hits := range grid {
		if len(hits) > 1 {
			if coord != "0,0" {
				uniqueWires := make(map[int]bool)
				for _, hit := range hits {
					uniqueWires[hit] = true
				}

				if len(uniqueWires) > 1 {
					overlaps = append(overlaps, coord)
				}
			}
		}
	}
	// fmt.Println(overlaps)

	closest := 0
	for _, coord := range overlaps {
		// fmt.Println(coord)
		split := strings.Split(coord, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		distance := Abs(x) + Abs(y)
		// fmt.Println(distance)
		if distance < closest || closest == 0 {
			closest = x + y
		}
	}


	return closest
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
