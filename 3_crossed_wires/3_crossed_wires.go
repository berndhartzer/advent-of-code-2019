package three_crossed_wires

import (
	"fmt"
	"strings"
	"strconv"
)

type Coord struct {
	X, Y int
}

type GridSquare struct {
	WireOneSteps, WireTwoSteps int
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
	grid := make(map[string]*GridSquare)

	// Build up the grid with wires
	for i, wire := range wires {
		coord := Coord{0, 0}
		steps := 0

		for _, instruction := range wire {
			split := []rune(instruction)
			distance, _ := strconv.Atoi(string(split[1:]))

			for j := 0; j < distance; j++ {
				steps += 1
				coord.Move(string(split[0]))

				location := fmt.Sprintf("%d,%d", coord.X, coord.Y)

				square, ok := grid[location]
				if !ok {
					grid[location] = &GridSquare{0, 0}
					if i == 1 {
						grid[location].WireOneSteps = steps
					} else {
						grid[location].WireTwoSteps = steps
					}
				} else {
					if i == 1 {
						if square.WireOneSteps == 0 {
							square.WireOneSteps = steps
						}
					} else {
						if square.WireTwoSteps == 0 {
							square.WireTwoSteps = steps
						}
					}

					grid[location] = square
				}
			}
		}
	}

	leastSteps := 0
	for _, point := range grid {
		if point.WireOneSteps != 0 && point.WireTwoSteps != 0 {
			steps := point.WireOneSteps + point.WireTwoSteps
			if steps < leastSteps || leastSteps == 0 {
				leastSteps = steps
			}
		}
	}

	return leastSteps
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
