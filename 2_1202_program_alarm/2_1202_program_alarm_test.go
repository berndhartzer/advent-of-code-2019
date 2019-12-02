package two_1202_program_alarm

import (
	"testing"
)


func TestTwelveZeroTwoProgramAlarmPartOneExampleInput(t *testing.T) {
	exampleInputOne := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	exampleInputTwo := [5]int{1, 0, 0, 0, 99}
	exampleInputThree := [5]int{2, 3, 0, 3, 99}
	exampleInputFour := [6]int{2, 4, 4, 5, 99, 0}
	exampleInputFive := [9]int{1, 1, 1, 4, 99, 5, 6, 0, 99}

	t.Log(exampleInputOne)
	exampleOuputOne := TwelveZeroTwoProgramAlarmPartOne(exampleInputOne)
	t.Log(exampleOuputOne)
}
