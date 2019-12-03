package three_crossed_wires

import (
	"bufio"
	"log"
	"os"
	"testing"
	"strings"
)

func TestCrossedWiresPartOneExampleInput(t *testing.T) {
	exampleInputOne := [][]string{
		{"R8", "U5", "L5", "D3"},
		{"U7", "R6", "D4", "L4"},
	}

	exampleInputTwo := [][]string{
		{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
	}

	exampleInputThree := [][]string{
		{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
		{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
	}

	expectedOutputOne := 6
	expectedOutputTwo := 159
	expectedOutputThree := 135

	exampleOutputOne := CrossedWiresPartOne(exampleInputOne)
	exampleOutputTwo := CrossedWiresPartOne(exampleInputTwo)
	exampleOutputThree := CrossedWiresPartOne(exampleInputThree)

	if exampleOutputOne != expectedOutputOne {
		t.Fail()
	}

	if exampleOutputTwo != expectedOutputTwo {
		t.Fail()
	}

	if exampleOutputThree != expectedOutputThree {
		t.Fail()
	}
}

func TestCrossedWiresPartOneActualInput(t *testing.T) {
	file, err := os.Open("./3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Split(line, ","))
	}

	output := CrossedWiresPartOne(input)
	t.Log(output)
}

func TestCrossedWiresPartTwoExampleInput(t *testing.T) {
	exampleInputOne := [][]string{
		{"R8", "U5", "L5", "D3"},
		{"U7", "R6", "D4", "L4"},
	}

	exampleInputTwo := [][]string{
		{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
	}

	exampleInputThree := [][]string{
		{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
		{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
	}

	expectedOutputOne := 30
	expectedOutputTwo := 610
	expectedOutputThree := 410

	exampleOutputOne := CrossedWiresPartTwo(exampleInputOne)
	exampleOutputTwo := CrossedWiresPartTwo(exampleInputTwo)
	exampleOutputThree := CrossedWiresPartTwo(exampleInputThree)

	if exampleOutputOne != expectedOutputOne {
		t.Fail()
	}

	if exampleOutputTwo != expectedOutputTwo {
		t.Fail()
	}

	if exampleOutputThree != expectedOutputThree {
		t.Fail()
	}
}
