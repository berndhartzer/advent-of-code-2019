package twelve_the_n_body_problem

import (
	"testing"
)

func TestTheNBodyProblemPartOneExampleInput(t *testing.T) {
	exampleInput := []string{
		"<x=-1, y=0, z=2>",
		"<x=2, y=-10, z=-7>",
		"<x=4, y=-8, z=8>",
		"<x=3, y=5, z=-1>",
	}
	expectedOutput := 179

	output := TheNBodyProblemPartOne(exampleInput)
	if output != expectedOutput {
		t.Fail()
	}
}
