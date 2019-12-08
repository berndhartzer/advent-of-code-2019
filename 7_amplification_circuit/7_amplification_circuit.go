package seven_amplification_circuit

func AmplificationCircuitPartOne(program []int, sequence []int) int {
	input := [2]int{}
	outputSignal := 0

	for _, phase := range sequence {
		input[0] = phase
		input[1] = outputSignal
		outputSignal = IntCodeComputer(program, input)
	}

	return outputSignal
}

// From SunnyWithAChanceOfAsteroidsPartTwo
func IntCodeComputer(program []int, input [2]int) int {
	i := 0
	inputIndex := 0
	for {
		if program[i] == 99 {
			break
		}

		opCodeRaw := program[i]
		opCodeParts := [5]int{}
		for j := 4; j >= 0; j -= 1 {
			opCodeParts[j] = opCodeRaw % 10
			opCodeRaw = opCodeRaw / 10
		}

		opCode := opCodeParts[4]
		paramOneMode := opCodeParts[2]
		paramTwoMode := opCodeParts[1]

		paramOne := program[i+1]
		if paramOneMode == 0 {
			paramOne = program[paramOne]
		}

		switch opCode {
		case 1, 2:
			paramTwo := program[i+2]
			if paramTwoMode == 0 {
				paramTwo = program[paramTwo]
			}

			outputPos := program[i+3]

			if opCode == 1 {
				program[outputPos] = paramOne + paramTwo
			} else {
				program[outputPos] = paramOne * paramTwo
			}
			i += 4

		case 3:
			posOne := program[i+1]
			program[posOne] = input[inputIndex]
			inputIndex += 1
			i += 2

		case 4:
			program = append(program, paramOne)
			i += 2

		case 5, 6:
			paramTwo := program[i+2]
			if paramTwoMode == 0 {
				paramTwo = program[paramTwo]
			}

			if (opCode == 5 && paramOne != 0) || (opCode == 6 && paramOne == 0) {
				i = paramTwo
			} else {
				i += 3
			}

		case 7, 8:
			paramTwo := program[i+2]
			if paramTwoMode == 0 {
				paramTwo = program[paramTwo]
			}

			outputPos := program[i+3]

			if (opCode == 7 && paramOne < paramTwo) || (opCode == 8 && paramOne == paramTwo) {
				program[outputPos] = 1
			} else {
				program[outputPos] = 0
			}
			i += 4

		}
	}

	return program[len(program)-1]
}
