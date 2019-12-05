package five_sunny_with_a_chance_of_asteroids

import (
	"io/ioutil"
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
		output := SunnyWithAChanceOfAsteroidsPartOne(input, 1)

		if output != expectedExampleOutputs[i][len(expectedExampleOutputs[i])-1] {
			t.Fail()
		}

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
