package ten_monitoring_station

import (
	"testing"
)

func TestMonitoringStationPartOneExampleInput(t *testing.T) {
	exampleInput := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}
	expectedOutput := 8

	output := MonitoringStationPartOne(exampleInput)
	if output != expectedOutput {
		t.Fail()
	}
}
