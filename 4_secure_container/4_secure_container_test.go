package four_secure_container

import (
	"testing"
)

func TestSecureContainerPartOneActualInput(t *testing.T) {
	input := []int{183564, 657474}
	output := SecureContainerPartOne(input[0], input[1])
	t.Log(output)
}

func TestSecureContainerPartTwoExampleInput(t *testing.T) {
	exampleInputs := []int{112233, 123444, 111122}
	expectedOutputs := []int{1, 0, 1}

	exampleOutputOne := SecureContainerPartTwo(exampleInputs[0], exampleInputs[0])
	exampleOutputTwo := SecureContainerPartTwo(exampleInputs[1], exampleInputs[1])
	exampleOutputThree := SecureContainerPartTwo(exampleInputs[2], exampleInputs[2])

	if exampleOutputOne != expectedOutputs[0] {
		t.Log("one")
		t.Fail()
	}

	if exampleOutputTwo != expectedOutputs[1] {
		t.Log("two")
		t.Fail()
	}

	if exampleOutputThree != expectedOutputs[2] {
		t.Log("three")
		t.Fail()
	}
}

func TestSecureContainerPartTwoActualInput(t *testing.T) {
	input := []int{183564, 657474}
	output := SecureContainerPartTwo(input[0], input[1])
	t.Log(output)
}
