package twelve_the_n_body_problem

import (
	"bufio"
	"os"
	"testing"
)

func TestTheNBodyProblemPartOneExampleInput(t *testing.T) {
	exampleInputs := [][]string{
		{
			"<x=-1, y=0, z=2>",
			"<x=2, y=-10, z=-7>",
			"<x=4, y=-8, z=8>",
			"<x=3, y=5, z=-1>",
		},
		{
			"<x=-8, y=-10, z=0>",
			"<x=5, y=5, z=10>",
			"<x=2, y=-7, z=3>",
			"<x=9, y=-8, z=-3>",
		},
	}
	exampleSimulations := []int{10, 100}
	expectedOutputs := []int{179, 1940}

	for i, input := range exampleInputs {
		output := TheNBodyProblemPartOne(input, exampleSimulations[i])
		if output != expectedOutputs[i] {
			t.Fail()
		}
	}
}

func TestTheNBodyProblemPartOneActualInput(t *testing.T) {
	var input []string
	file, _ := os.Open("./12_input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	output := TheNBodyProblemPartOne(input, 1000)
	t.Log(output)
}

func TestTheNBodyProblemPartTwoExampleInput(t *testing.T) {
	exampleInputs := [][]string{
		{
			"<x=-1, y=0, z=2>",
			"<x=2, y=-10, z=-7>",
			"<x=4, y=-8, z=8>",
			"<x=3, y=5, z=-1>",
		},
		{
			"<x=-8, y=-10, z=0>",
			"<x=5, y=5, z=10>",
			"<x=2, y=-7, z=3>",
			"<x=9, y=-8, z=-3>",
		},
	}
	expectedOutputs := []int{2772, 4686774924}

	for i, input := range exampleInputs {
		output := TheNBodyProblemPartTwo(input)
		if output != expectedOutputs[i] {
			t.Fail()
		}
	}

}
