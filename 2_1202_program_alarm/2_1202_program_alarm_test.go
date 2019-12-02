package two_1202_program_alarm

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestTwelveZeroTwoProgramAlarmPartOneExampleInput(t *testing.T) {
	exampleInputs := [][]int{
		{1, 0, 0, 0, 99},
		{2, 3, 0, 3, 99},
		{2, 4, 4, 5, 99, 0},
		{1, 1, 1, 4, 99, 5, 6, 0, 99},
	}

	expectedExampleOutputs := [][]int{
		{2, 0, 0, 0, 99},
		{2, 3, 0, 6, 99},
		{2, 4, 4, 5, 99, 9801},
		{30, 1, 1, 4, 2, 5, 6, 0, 99},
	}

	for i, input := range exampleInputs {
		output := TwelveZeroTwoProgramAlarmPartOne(input)

		if output != expectedExampleOutputs[i][0] {
			t.Fail()
		}

		for j, num := range input {
			if num != expectedExampleOutputs[i][j] {
				t.Fail()
			}
		}
	}
}

func TestTwelveZeroTwoProgramAlarmPartOneActualInput(t *testing.T) {
	inputBytes, _ := ioutil.ReadFile("2_input.txt")
	inputString := string(inputBytes)
	inputSplit := strings.Split(inputString, ",")
	input := []int{}
	for _, str := range inputSplit {
		num, _ := strconv.Atoi(str)
		input = append(input, num)
	}

	// Perform the correction
	input[1] = 12
	input[2] = 2

	output := TwelveZeroTwoProgramAlarmPartOne(input)
	t.Log(output)
}

func TestTwelveZeroTwoProgramAlarmPartTwoActualInput(t *testing.T) {
	inputBytes, _ := ioutil.ReadFile("2_input.txt")
	inputString := string(inputBytes)
	inputSplit := strings.Split(inputString, ",")
	input := []int{}
	for _, str := range inputSplit {
		num, _ := strconv.Atoi(str)
		input = append(input, num)
	}

	output := TwelveZeroTwoProgramAlarmPartTwo(input, 19690720)
	t.Log(output)
}
