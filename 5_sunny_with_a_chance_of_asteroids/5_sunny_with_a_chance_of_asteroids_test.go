package five_sunny_with_a_chance_of_asteroids

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSunnyWithAChanceOfAsteroidsPartOneExampleInput(t *testing.T) {
	exampleInputs := [][]int{
		{3, 0, 4, 0, 99},
		{1002, 4, 3, 4, 33},
		{1101, 100, -1, 4, 0},
	}

	expectedExampleOutputs := [][]int{
		{1, 0, 4, 0, 99, 0},
		{1002, 4, 3, 4, 99},
		{1101, 100, -1, 4, 99},
	}

	for i, input := range exampleInputs {
		SunnyWithAChanceOfAsteroidsPartOne(input, 1)

		for j, num := range input {
			if num != expectedExampleOutputs[i][j] {
				t.Fail()
			}
		}
	}
}

func TestSunnyWithAChanceOfAsteroidsPartOneActualInput(t *testing.T) {
	inputBytes, _ := ioutil.ReadFile("5_input.txt")
	inputString := string(inputBytes)
	inputSplit := strings.Split(inputString, ",")
	input := []int{}
	for _, str := range inputSplit {
		num, _ := strconv.Atoi(str)
		input = append(input, num)
	}

	output := SunnyWithAChanceOfAsteroidsPartOne(input, 1)
	t.Log(output)
}

func TestSunnyWithAChanceOfAsteroidsPartTwoExampleInput(t *testing.T) {
	exampleInputs := [][]int{
		{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},

		{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		{3, 3, 1107, -1, 8, 3, 4, 3, 99},

		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},

		{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
		{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
		{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
	}
	exampleInputParameters := []int{
		1,
		8,
		1,
		9,

		8,
		100,
		7,
		9,

		0,
		2,
		0,
		5,

		6,
		8,
		10,
	}

	expectedOutputs := []int{
		0,
		1,
		1,
		0,

		1,
		0,
		1,
		0,

		0,
		1,
		0,
		1,

		999,
		1000,
		1001,
	}

	for i := range exampleInputs {
		output := SunnyWithAChanceOfAsteroidsPartTwo(exampleInputs[i], exampleInputParameters[i])

		if output != expectedOutputs[i] {
			t.Fail()
		}
	}
}

func TestSunnyWithAChanceOfAsteroidsPartTwoActualInput(t *testing.T) {
	var inputSplit []string
	file, _ := os.Open("./5_input.txt")
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

	output := SunnyWithAChanceOfAsteroidsPartTwo(input, 5)
	t.Log(output)
}
