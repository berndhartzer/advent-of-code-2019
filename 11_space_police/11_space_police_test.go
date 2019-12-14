package eleven_space_police

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSpacePolicePartOneActualInput(t *testing.T) {
	var inputSplit []string
	file, _ := os.Open("./11_input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputSplit = strings.Split(line, ",")
	}

	input := []int{}

	for _, str := range inputSplit {
		num, _ := strconv.Atoi(str)
		input = append(input, num)
	}

	output := SpacePolicePartOne(input)
	t.Log(output)
}

func TestSpacePolicePartTwoActualInput(t *testing.T) {
	var inputSplit []string
	file, _ := os.Open("./11_input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputSplit = strings.Split(line, ",")
	}

	input := []int{}

	for _, str := range inputSplit {
		num, _ := strconv.Atoi(str)
		input = append(input, num)
	}

	output := SpacePolicePartTwo(input)
	for _, row := range output {
		t.Log(row)
	}
}
