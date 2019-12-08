package seven_amplification_circuit

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestAmplificationCircuitPartOneExampleInput(t *testing.T) {
	exampleInputs := [][]int{
		{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
	}
	exampleInputSequences := [][]int{
		{4, 3, 2, 1, 0},
		{0, 1, 2, 3, 4},
		{1, 0, 4, 3, 2},
	}

	expectedExampleOutputs := []int{
		43210,
		54321,
		65210,
	}

	for i := range exampleInputs {
		output := AmplificationCircuitPartOne(exampleInputs[i], exampleInputSequences[i])

		if output != expectedExampleOutputs[i] {
			t.Fail()
		}
	}
}

func TestAmplificationCircuitPartOneActualInput(t *testing.T) {
	var inputSplit []string
	file, _ := os.Open("./7_input.txt")
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

	highest := 0
	basePermutation := []int{0, 1, 2, 3, 4}

	all := permutations(basePermutation)

	for _, permutation := range all {
		inputCopy := make([]int, len(input))
		copy(inputCopy, input)

		output := AmplificationCircuitPartOne(inputCopy, permutation)
		if output > highest {
			highest = output
		}
	}

	t.Log(highest)
}

func TestAmplificationCircuitPartTwoExampleInput(t *testing.T) {
	exampleInputs := [][]int{
		{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
	}
	exampleInputSequences := [][]int{
		{9, 8, 7, 6, 5},
		{9, 7, 8, 5, 6},
	}

	expectedExampleOutputs := []int{
		139629729,
		18216,
	}

	for i := range exampleInputs {
		output := AmplificationCircuitPartTwo(exampleInputs[i], exampleInputSequences[i])

		if output != expectedExampleOutputs[i] {
			t.Fail()
		}
	}
}

func TestAmplificationCircuitPartTwoActualInput(t *testing.T) {
	var inputSplit []string
	file, _ := os.Open("./7_input.txt")
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

	highest := 0
	basePermutation := []int{5, 6, 7, 8, 9}

	all := permutations(basePermutation)

	for _, permutation := range all {
		inputCopy := make([]int, len(input))
		copy(inputCopy, input)

		output := AmplificationCircuitPartTwo(inputCopy, permutation)
		if output > highest {
			highest = output
		}
	}

	t.Log(highest)
}

// https://stackoverflow.com/a/30226442/5374111
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
