package two_1202_program_alarm

func TwelveZeroTwoProgramAlarmPartOne(program []int) int {
	for i := 0; i < len(program); i += 4 {
		if program[i] == 99 {
			break
		}

		posOne := program[i+1]
		posTwo := program[i+2]
		outputPos := program[i+3]

		if program[i] == 1 {
			program[outputPos] = program[posOne] + program[posTwo]
		} else if program[i] == 2 {
			program[outputPos] = program[posOne] * program[posTwo]
		}
	}

	return program[0]
}
