package seven_amplification_circuit

func AmplificationCircuitPartOne(program []int, sequence []int) int {
	input := [2]int{}
	outputSignal := 0

	for _, phase := range sequence {
		input[0] = phase
		input[1] = outputSignal
		outputSignal = IntCodeComputerPartOne(program, input)
	}

	return outputSignal
}

func AmplificationCircuitPartTwo(program []int, sequence []int) int {
	amplifiers := make(map[int][]int)
	amplifierIndexes := make(map[int]int)
	for i := 0; i < 5; i += 1 {
		programCopy := make([]int, len(program))
		copy(programCopy, program)
		amplifiers[i] = programCopy
		amplifierIndexes[i] = 0
	}

	input := [2]int{}
	amplifierIndex := 0
	phaseIndex := 0
	outputSignal := 0
	instructionPointer := 0
	terminated := false

	for !terminated {
		input[0] = outputSignal
		input[1] = 0
		if phaseIndex < len(sequence) {
			input[0] = sequence[phaseIndex]
			input[1] = outputSignal
		}

		outputSignal, instructionPointer, terminated = IntCodeComputerPartTwo(
			amplifiers[amplifierIndex],
			input,
			amplifierIndexes[amplifierIndex],
		)

		amplifierIndexes[amplifierIndex] = instructionPointer

		amplifierIndex += 1
		if amplifierIndex >= len(amplifiers) {
			amplifierIndex = 0
		}
		phaseIndex += 1
	}

	return outputSignal
}

func IntCodeComputerPartOne(program []int, input [2]int) int {
	output, _, _ := IntCodeComputer(program, input, 0, false)
	return output
}

func IntCodeComputerPartTwo(program []int, input [2]int, index int) (int, int, bool) {
	return IntCodeComputer(program, input, index, true)
}

func IntCodeComputer(program []int, input [2]int, i int, returnOutput bool) (int, int, bool) {
	inputIndex := 0
	for {
		if program[i] == 99 {
			if returnOutput {
				return input[0], i, true
			}
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
			if returnOutput {
				return program[len(program)-1], i, false
			}

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

	return program[len(program)-1], i, false
}
