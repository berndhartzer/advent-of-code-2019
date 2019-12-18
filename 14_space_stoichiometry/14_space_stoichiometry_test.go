package fourteen_space_stoichiometry

import (
	"testing"
)

func TestSpaceStoichiometryPartOneExampleInput(t *testing.T) {
	exampleInputs := [][]string{
		{
			"10 ORE => 10 A",
			"1 ORE => 1 B",
			"7 A, 1 B => 1 C",
			"7 A, 1 C => 1 D",
			"7 A, 1 D => 1 E",
			"7 A, 1 E => 1 FUEL",
		},
	}
	expectedOutputs := []int{31}

	for i, input := range exampleInputs {
		output := SpaceStoichiometryPartOne(input)

		if output != expectedOutputs[i] {
			t.Fail()
		}
	}

}
