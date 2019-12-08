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

func TestSpaceImageFormatPartTwoExampleInput(t *testing.T) {
	input := "0222112222120000"

	output := SpaceImageFormatPartTwo(input, 2, 2)
	for _, row := range output {
		t.Log(row)
	}
}

func TestSpaceImageFormatPartTwoActualInput(t *testing.T) {
	var input []string
	file, _ := os.Open("./8_input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	output := SpaceImageFormatPartTwo(input[0], 25, 6)
	for _, row := range output {
		t.Log(row)
	}
}
