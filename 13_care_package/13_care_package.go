package thirteen_care_package

import (
	"fmt"
)

type Game struct {
	Grid    map[string]int
	PaddleX int
	BallX   int
	Score   int
}

func (g *Game) moveJoystick() int {
	position := 0
	if g.PaddleX > g.BallX {
		position = -1
	} else if g.PaddleX < g.BallX {
		position = 1
	}
	return position
}

func CarePackagePartOne(program []int) int {
	computer := NewIntCodeComputer(program)
	allOutput := []int{}

	i := 1
	tileCount := 0
	terminated := false
	for !terminated {
		terminated = computer.Run(&allOutput, [2]int{0, 0}, true)
		output := allOutput[len(allOutput)-1]

		if i == 3 && output == 2 {
			tileCount += 1
		}

		i += 1
		if i > 3 {
			i = 1
		}
	}

	return tileCount
}

func CarePackagePartTwo(program []int) int {
	// Insert quarters
	program[0] = 2

	game := &Game{}
	game.Grid = make(map[string]int)

	computer := NewIntCodeComputer(program)
	allOutput := []int{}

	i := 1
	x := 0
	y := 0
	input := 0
	terminated := false
	for !terminated {
		terminated = computer.Run(&allOutput, [2]int{input, input}, true)
		output := allOutput[len(allOutput)-1]

		switch i {
		case 1:
			x = output
		case 2:
			y = output
		case 3:
			if x == -1 && y == 0 {
				game.Score = output
			} else {
				game.Grid[fmt.Sprintf("%d,%d", x, y)] = output

				switch output {
				case 3:
					game.PaddleX = x
				case 4:
					game.BallX = x
				}
			}
		}

		input = game.moveJoystick()

		i += 1
		if i > 3 {
			i = 1
		}
	}

	return game.Score
}

type IntCodeComputer struct {
	Program            []int
	InstructionPointer int
	InputPointer       int
	OpCode             int
	ParamMode          map[int]int
	RelativeBase       int
}

func NewIntCodeComputer(program []int) *IntCodeComputer {
	c := new(IntCodeComputer)

	programBase := make([]int, 5000)
	copy(programBase, program)

	c.Program = programBase
	c.ParamMode = make(map[int]int)

	return c
}

func (c *IntCodeComputer) ParseOpCode() {
	opCodeRaw := c.Program[c.InstructionPointer]
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

func (c *IntCodeComputer) getValue(param int) int {
	value := 0
	switch c.ParamMode[param] {
	case 0:
		value = c.Program[c.Program[c.InstructionPointer+param]]
	case 1:
		value = c.Program[c.InstructionPointer+param]
	case 2:
		value = c.Program[c.RelativeBase+c.Program[c.InstructionPointer+param]]
	}
	return value
}

func (c *IntCodeComputer) putValue(param, value int) {
	switch c.ParamMode[param] {
	case 0:
		c.Program[c.Program[c.InstructionPointer+param]] = value
	case 2:
		c.Program[c.RelativeBase+c.Program[c.InstructionPointer+param]] = value
	}
}

func (c *IntCodeComputer) AdjustRelativeBase(adj int) {
	c.RelativeBase += adj
}

func (c *IntCodeComputer) Run(dst *[]int, input [2]int, returnOutput bool) bool {
	for {
		if c.Program[c.InstructionPointer] == 99 {
			if returnOutput {
				return true
			}
			break
		}

		c.ParseOpCode()

		switch c.OpCode {
		case 1:
			c.putValue(3, (c.getValue(1) + c.getValue(2)))
			c.InstructionPointer += 4

		case 2:
			c.putValue(3, (c.getValue(1) * c.getValue(2)))
			c.InstructionPointer += 4

		case 3:
			c.putValue(1, input[c.InputPointer])

			c.InputPointer += 1
			if c.InputPointer >= len(input) {
				c.InputPointer = 0
			}
			c.InstructionPointer += 2

		case 4:
			*dst = append(*dst, c.getValue(1))

			c.InstructionPointer += 2
			if returnOutput {
				return false
			}

		case 5:
			if c.getValue(1) != 0 {
				c.InstructionPointer = c.getValue(2)
			} else {
				c.InstructionPointer += 3
			}

		case 6:
			if c.getValue(1) == 0 {
				c.InstructionPointer = c.getValue(2)
			} else {
				c.InstructionPointer += 3
			}

		case 7:
			if c.getValue(1) < c.getValue(2) {
				c.putValue(3, 1)
			} else {
				c.putValue(3, 0)
			}
			c.InstructionPointer += 4

		case 8:
			if c.getValue(1) == c.getValue(2) {
				c.putValue(3, 1)
			} else {
				c.putValue(3, 0)
			}
			c.InstructionPointer += 4

		case 9:
			c.AdjustRelativeBase(c.getValue(1))
			c.InstructionPointer += 2
		}
	}

	return false
}
