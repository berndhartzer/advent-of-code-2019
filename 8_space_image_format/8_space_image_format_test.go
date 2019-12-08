package eight_space_image_format

import (
	"bufio"
	"os"
	"testing"
)

func TestSpaceImageFormatPartOneExampleInput(t *testing.T) {
	input := "123456789012"
	expectedOutput := 1

	output := SpaceImageFormatPartOne(input, 3, 2)
	if output != expectedOutput {
		t.Fail()
	}
}

func TestSpaceImageFormatPartOneActualInput(t *testing.T) {
	var input []string
	file, _ := os.Open("./8_input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	output := SpaceImageFormatPartOne(input[0], 25, 6)
	t.Log(output)
}
