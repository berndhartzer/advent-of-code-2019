package nine_sensor_boost

import (
	"fmt"
)

func SensorBoostPartOne(program []int, input int) []int {
	// output, _, _ := RunIntCodeComputer(program, &dst, [2]int{0, 0}, 0, false)
	output := []int{}
	RunIntCodeComputer(program, &output, [2]int{input, input}, 0, false)
	return output
}

type IntCodeComputer struct {
	Program []int
	OpCode int
	ParamMode map[int]int
	RelativeBase int
}

func (c *IntCodeComputer) ParseOpCode(i int) {
	opCodeRaw := c.Program[i]
	opCodeParts := [5]int{}
	for j := 4; j >= 0; j -= 1 {
		opCodeParts[j] = opCodeRaw % 10
		opCodeRaw = opCodeRaw / 10
	}

	c.OpCode = opCodeParts[4]
	c.ParamMode[1] = opCodeParts[2]
	c.ParamMode[2] = opCodeParts[1]
	c.ParamMode[3] = opCodeParts[0]
}

func (c *IntCodeComputer) getValue(i, param int) int {
	value := 0
	switch c.ParamMode[param] {
	case 0:
		value = c.Program[c.Program[i+param]]
	case 1:
		value = c.Program[i+param]
	case 2:
		value = c.Program[c.RelativeBase + c.Program[i+param]]
	}
	return value
}

func RunIntCodeComputer(originalProgram []int, dst *[]int, input [2]int, i int, returnOutput bool) {
	program := make([]int, 5000)
	copy(program, originalProgram)

	computer := &IntCodeComputer{
		Program: program,
		ParamMode: make(map[int]int),
	}

	inputIndex := 0
	relativeBase := 0
	for {
		if program[i] == 99 {
			if returnOutput {
				// return *(dst)[len(*dst)-1], i, true
				return
			}
			break
		}

		opCodeRaw := program[i]
		opCodeParts := [5]int{}
		for j := 4; j >= 0; j -= 1 {
			opCodeParts[j] = opCodeRaw % 10
			opCodeRaw = opCodeRaw / 10
		}
		// fmt.Println(opCodeParts)

		computer.ParseOpCode(i)

		opCode := opCodeParts[4]
		// paramOneMode := opCodeParts[2]
		// paramTwoMode := opCodeParts[1]
		// paramThreeMode := opCodeParts[0]
		paramMode := map[int]int{
			1: opCodeParts[2],
			2: opCodeParts[1],
			3: opCodeParts[0],
		}

		paramOne := 0
		if paramMode[1] == 0 {
			paramOne = program[program[i+1]]
		} else if paramMode[1] == 1 {
			paramOne = program[i+1]
		} else if paramMode[1] == 2 {
			paramOne = program[relativeBase + program[i+1]]
		}
		fmt.Println("===")
		fmt.Println("paramOne", paramMode[1], paramOne)
		fmt.Println("paramOne", computer.ParamMode[1], computer.getValue(i, 1))

		switch opCode {
		case 0:
			return // TODO: remove
		case 1, 2:
			paramTwo := 0
			if paramMode[2] == 0 {
				paramTwo = program[program[i+2]]
			} else if paramMode[2] == 1 {
				paramTwo = program[i+2]
			} else if paramMode[2] == 2 {
				paramTwo = program[relativeBase + program[i+2]]
			}

			// outputPos := program[i+3]
			paramThree := 0
			if paramMode[3] == 0 {
				paramThree = program[program[i+3]]
			} else if paramMode[3] == 1 {
				// 
			} else if paramMode[3] == 2 {
				paramThree = program[relativeBase + program[i+3]]
			}

			if opCode == 1 {
				// program[outputPos] = paramOne + paramTwo
				program[paramThree] = paramOne + paramTwo
			} else {
				// program[outputPos] = paramOne * paramTwo
				program[paramThree] = paramOne * paramTwo
			}
			i += 4

		case 3:
			// posOne := program[i+1]
			// if paramMode[1] == 0 {
			// 	program[posOne] = input[inputIndex]
			// } else if paramMode[1] == 2 {
			// 	// program[posOne] = input[inputIndex]
			// 	program[relativeBase + program[paramOne]] = input[inputIndex]
			// }


			// paramOne := 0
			if paramMode[1] == 0 {
				program[program[i+1]] = input[inputIndex]
			} else if paramMode[1] == 1 {
				// paramOne = program[i+1]
			} else if paramMode[1] == 2 {
				program[relativeBase + program[i+1]] = input[inputIndex]
			}


			inputIndex += 1
			i += 2

		case 4:
			// program = append(program, paramOne)
			*dst = append(*dst, paramOne)
			i += 2
			if returnOutput {
				// return program[len(program)-1], i, false
				return
			}

		case 5, 6:
			paramTwo := 0
			if paramMode[2] == 0 {
				paramTwo = program[program[i+2]]
			} else if paramMode[2] == 1 {
				paramTwo = program[i+2]
			} else if paramMode[2] == 2 {
				paramTwo = program[relativeBase + program[i+2]]
			}

			if (opCode == 5 && paramOne != 0) || (opCode == 6 && paramOne == 0) {
				i = paramTwo
			} else {
				i += 3
			}

		case 7, 8:
			paramTwo := 0
			if paramMode[2] == 0 {
				paramTwo = program[program[i+2]]
			} else if paramMode[2] == 1 {
				paramTwo = program[i+2]
			} else if paramMode[2] == 2 {
				paramTwo = program[relativeBase + program[i+2]]
			}

			// outputPos := program[i+3]
			paramThree := 0
			if paramMode[3] == 0 {
				paramThree = program[program[i+3]]
			} else if paramMode[3] == 1 {
				// 
			} else if paramMode[3] == 2 {
				paramThree = program[relativeBase + program[i+3]]
			}

			if (opCode == 7 && paramOne < paramTwo) || (opCode == 8 && paramOne == paramTwo) {
				// program[outputPos] = 1
				program[paramThree] = 1
			} else {
				// program[outputPos] = 0
				program[paramThree] = 0
			}
			i += 4

		case 9:
			relativeBase += paramOne
			i += 2

		}
	}

	// return dst[len(dst)-1], i, false
	return
}
