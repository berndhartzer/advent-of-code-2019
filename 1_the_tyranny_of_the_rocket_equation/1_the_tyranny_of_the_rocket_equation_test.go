package one_the_tyranny_of_the_rocket_equation

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestTheTyrannyOfTheRocketEquationPartOneActualInput(t *testing.T) {
	file, err := os.Open("./1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.ParseFloat(line, 64)
		input = append(input, num)
	}

	output := TheTyrannyOfTheRocketEquationPartOne(input)
	t.Log(fmt.Sprintf("%f\n", output))
}

func TestTheTyrannyOfTheRocketEquationPartTwoActualInput(t *testing.T) {
	file, err := os.Open("./1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.ParseFloat(line, 64)
		input = append(input, num)
	}

	output := TheTyrannyOfTheRocketEquationPartTwo(input)
	t.Log(fmt.Sprintf("%f\n", output))
}
