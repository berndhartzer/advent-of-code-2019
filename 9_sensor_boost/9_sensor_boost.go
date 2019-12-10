package nine_sensor_boost

func SensorBoost(program []int, input int) []int {
	output := []int{}
	RunIntCodeComputer(program, &output, [2]int{input, input}, 0, false)
	return output
}

type IntCodeComputer struct {
	Program            []int
	InstructionPointer int
	OpCode             int
	ParamMode          map[int]int
	RelativeBase       int
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
		value = c.Program[c.RelativeBase+c.Program[i+param]]
	}
	return value
}

func (c *IntCodeComputer) putValue(i, param, value int) {
	switch c.ParamMode[param] {
	case 0:
		c.Program[c.Program[i+param]] = value
	case 2:
		c.Program[c.RelativeBase+c.Program[i+param]] = value
	}
}

func (c *IntCodeComputer) AdjustRelativeBase(adj int) {
	c.RelativeBase += adj
}

func RunIntCodeComputer(originalProgram []int, dst *[]int, input [2]int, i int, returnOutput bool) {
	program := make([]int, 5000)
	copy(program, originalProgram)

	icc := &IntCodeComputer{
		Program:   program,
		ParamMode: make(map[int]int),
	}

	inputIndex := 0

	for {
		if program[i] == 99 {
			if returnOutput {
				return
			}
			break
		}

		icc.ParseOpCode(i)

		switch icc.OpCode {
		case 1:
			icc.putValue(i, 3, (icc.getValue(i, 1) + icc.getValue(i, 2)))
			i += 4

		case 2:
			icc.putValue(i, 3, (icc.getValue(i, 1) * icc.getValue(i, 2)))
			i += 4

		case 3:
			icc.putValue(i, 1, input[inputIndex])

			inputIndex += 1
			i += 2

		case 4:
			*dst = append(*dst, icc.getValue(i, 1))

			i += 2
			if returnOutput {
				return
			}

		case 5:
			if icc.getValue(i, 1) != 0 {
				i = icc.getValue(i, 2)
			} else {
				i += 3
			}
		case 6:
			if icc.getValue(i, 1) == 0 {
				i = icc.getValue(i, 2)
			} else {
				i += 3
			}

		case 7:
			if icc.getValue(i, 1) < icc.getValue(i, 2) {
				icc.putValue(i, 3, 1)
			} else {
				icc.putValue(i, 3, 0)
			}
			i += 4

		case 8:
			if icc.getValue(i, 1) == icc.getValue(i, 2) {
				icc.putValue(i, 3, 1)
			} else {
				icc.putValue(i, 3, 0)
			}
			i += 4

		case 9:
			icc.AdjustRelativeBase(icc.getValue(i, 1))
			i += 2
		}
	}

	return
}
