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

func TwelveZeroTwoProgramAlarmPartTwo(program []int, num int) int {
	noun := 0
	verb := 0

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			programCopy := make([]int, len(program))
			copy(programCopy, program)

			programCopy[1] = i
			programCopy[2] = j

			output := TwelveZeroTwoProgramAlarmPartOne(programCopy)

			if output == num {
				noun = i
				verb = j
				break
			}
		}
	}

	return (100 * noun) + verb
}
