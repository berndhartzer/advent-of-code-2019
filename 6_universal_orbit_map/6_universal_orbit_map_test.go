package six_universal_orbit_map

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestUniversalOrbitMapPartOneExampleInput(t *testing.T) {
	exampleInput := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	expectedOutput := 42

	output := UniversalOrbitMapPartOne(exampleInput)
	if output != expectedOutput {
		t.Fail()
	}
}

func TestUniversalOrbitMapPartOneActualInput(t *testing.T) {
	file, err := os.Open("./6_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	output := UniversalOrbitMapPartOne(input)
	t.Log(output)
}
