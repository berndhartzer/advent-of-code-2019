package four_secure_container

import (
	"testing"
)

func TestSecureContainerPartOneActualInput(t *testing.T) {
	input := []int{183564, 657474}
	output := SecureContainerPartOne(input[0], input[1])
	t.Log(output)
}

func TestSecureContainerPartTwoActualInput(t *testing.T) {
	input := []int{183564, 657474}
	output := SecureContainerPartTwo(input[0], input[1])
	t.Log(output)
}
