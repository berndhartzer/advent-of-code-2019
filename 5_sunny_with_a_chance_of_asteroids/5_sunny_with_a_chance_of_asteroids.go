package five_sunny_with_a_chance_of_asteroids

func SunnyWithAChanceOfAsteroidsPartOne(program []int, input int) int {
	i := 0
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

		switch opCode {
		case 1, 2:
			inputOne := 0
			posOne := program[i+1]
			if paramOneMode == 0 {
				inputOne = program[posOne]
			} else {
				inputOne = posOne
			}

			inputTwo := 0
			posTwo := program[i+2]
			if paramTwoMode == 0 {
				inputTwo = program[posTwo]
			} else {
				inputTwo = posTwo
			}

			outputPos := program[i+3]

			if opCode == 1 {
				program[outputPos] = inputOne + inputTwo
			} else {
				program[outputPos] = inputOne * inputTwo
			}
			i += 4
		case 3:
			posOne := program[i+1]
			program[posOne] = input
			i += 2
		case 4:
			inputOne := 0
			posOne := program[i+1]
			if paramOneMode == 0 {
				inputOne = program[posOne]
			} else {
				inputOne = posOne
			}
			program = append(program, inputOne)
			i += 2
		}
	}

	return program[len(program)-1]
}

func SunnyWithAChanceOfAsteroidsPartTwo(program []int, input int) int {
	i := 0
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
			program[posOne] = input
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
