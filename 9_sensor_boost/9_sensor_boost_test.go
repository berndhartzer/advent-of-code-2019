package nine_sensor_boost

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSensorBoostPartOneExampleInput(t *testing.T) {
	exampleOneInput := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	exampleOneOutput := SensorBoost(exampleOneInput, 0)

	for i := 0; i < len(exampleOneInput); i += 1 {
		if exampleOneInput[i] != exampleOneOutput[i] {
			t.Fail()
		}
	}

	exampleInputs := [][]int{
		{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
		{104, 1125899906842624, 99},
	}
	expectedExampleOutputs := []int{
		1219070632396864,
		1125899906842624,
	}

	for i := range exampleInputs {
		output := SensorBoost(exampleInputs[i], 0)

		if output[len(output)-1] != expectedExampleOutputs[i] {
			t.Fail()
		}
	}
}

func TestSensorBoostPartOneActualInput(t *testing.T) {
	var inputSplit []string
	file, _ := os.Open("./9_input.txt")
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

	output := SensorBoost(input, 1)
	t.Log(output)
}

func TestSensorBoostPartTwoActualInput(t *testing.T) {
	var inputSplit []string
	file, _ := os.Open("./9_input.txt")
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

	output := SensorBoost(input, 2)
	t.Log(output)
}
