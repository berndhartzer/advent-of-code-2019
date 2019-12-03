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

	// Build up the grid with wires
	for i, wire := range wires {
		coord := Coord{0, 0}

		for _, instruction := range wire {
			split := []rune(instruction)
			distance, _ := strconv.Atoi(string(split[1:]))

			for j := 0; j < distance; j++ {
				coord.Move(string(split[0]))
				grid[fmt.Sprintf("%d,%d", coord.X, coord.Y)] = append(grid[fmt.Sprintf("%d,%d", coord.X, coord.Y)], i)
			}
		}
	}

	// Find the overlaps
	overlaps := []string{}
	for coord, hits := range grid {
		if len(hits) > 1 {

			// Make sure we're not counting when one wire overlaps itself
			uniqueWires := make(map[int]bool)
			for _, hit := range hits {
				uniqueWires[hit] = true
			}

			if len(uniqueWires) > 1 {
				overlaps = append(overlaps, coord)
			}
		}
	}

	closest := 0
	for _, coord := range overlaps {
		split := strings.Split(coord, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		distance := Abs(x) + Abs(y)
		if distance < closest || closest == 0 {
			closest = distance
		}
	}

	return closest
}

func CrossedWiresPartTwo(wires [][]string) int {
	return 0
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
