package thirteen_care_package

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCarePackagePartOneActualInput(t *testing.T) {
	var inputSplit []string
	file, _ := os.Open("./13_input.txt")
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

	output := CarePackagePartOne(input)
	t.Log(output)
}
