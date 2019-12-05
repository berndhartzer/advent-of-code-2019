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
