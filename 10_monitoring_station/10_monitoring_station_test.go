package ten_monitoring_station

import (
	"bufio"
	"os"
	"testing"
)

func TestMonitoringStationPartOneExampleInput(t *testing.T) {
	exampleInputs := [][]string{
		{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##",
		},

		{
			"......#.#.",
			"#..#.#....",
			"..#######.",
			".#.#.###..",
			".#..#.....",
			"..#....#.#",
			"#..#....#.",
			".##.#..###",
			"##...#..#.",
			".#....####",
		},

		{
			"#.#...#.#.",
			".###....#.",
			".#....#...",
			"##.#.#.#.#",
			"....#.#.#.",
			".##..###.#",
			"..#...##..",
			"..##....##",
			"......#...",
			".####.###.",
		},

		{
			".#..#..###",
			"####.###.#",
			"....###.#.",
			"..###.##.#",
			"##.##.#.#.",
			"....###..#",
			"..#.#..#.#",
			"#..#.#.###",
			".##...##.#",
			".....#.#..",
		},

		{
			".#..##.###...#######",
			"##.############..##.",
			".#.######.########.#",
			".###.#######.####.#.",
			"#####.##.#.##.###.##",
			"..#####..#.#########",
			"####################",
			"#.####....###.#.#.##",
			"##.#################",
			"#####.##.###..####..",
			"..######..##.#######",
			"####.##.####...##..#",
			".#####..#.######.###",
			"##...#.##########...",
			"#.##########.#######",
			".####.#.###.###.#.##",
			"....##.##.###..#####",
			".#.#.###########.###",
			"#.#.#.#####.####.###",
			"###.##.####.##.#..##",
		},
	}

	expectedOutputs := []int{8, 33, 35, 41, 210}

	for i, input := range exampleInputs {
		output := MonitoringStationPartOne(input)

		if output != expectedOutputs[i] {
			t.Fail()
		}
	}
}

func TestMonitoringStationPartOneActualInput(t *testing.T) {
	file, _ := os.Open("./10_input.txt")

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	output := MonitoringStationPartOne(input)
	t.Log(output)
}

func TestMonitoringStationPartTwoExampleInput(t *testing.T) {
	exampleInput := []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}
	expectedOutput := 802

	output := MonitoringStationPartTwo(exampleInput, 200)

	if output != expectedOutput {
		t.Fail()
	}
}

func TestMonitoringStationPartTwoActualInput(t *testing.T) {
	file, _ := os.Open("./10_input.txt")

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	output := MonitoringStationPartTwo(input, 200)
	t.Log(output)
}
