package eleven_space_police

import (
	"fmt"
	"strconv"
	"strings"
)

func SpacePolicePartOne(program []int) int {
	hull := make(map[string]*Panel)

	robot := &Robot{}
	computer := NewIntCodeComputer(program)
	output := []int{}

	terminated := false
	for !terminated {
		coord := fmt.Sprintf("%d,%d", robot.X, robot.Y)
		_, ok := hull[coord]
		if !ok {
			hull[coord] = &Panel{".", false}
		}
		panel := hull[coord]

		input := 0
		if panel.Colour == "#" {
			input = 1
		}
		terminated = computer.Run(&output, [2]int{input, input}, true)
		thisOutput := output[len(output)-1]

		if thisOutput == 0 {
			panel.Colour = "."
		} else {
			panel.Colour = "#"
		}

		terminated = computer.Run(&output, [2]int{input, input}, true)
		thisOutput = output[len(output)-1]

		robot.Turn(thisOutput)
		robot.Move()

		panel.Painted = true
		hull[coord] = panel
	}

	return len(hull)
}

func SpacePolicePartTwo(program []int) [][]string {
	hull := make(map[string]*Panel)
	hull["0,0"] = &Panel{"#", false}

	robot := &Robot{}
	computer := NewIntCodeComputer(program)
	output := []int{}

	terminated := false
	for !terminated {
		coord := fmt.Sprintf("%d,%d", robot.X, robot.Y)
		_, ok := hull[coord]
		if !ok {
			hull[coord] = &Panel{".", false}

		}
		panel := hull[coord]

		input := 0
		if panel.Colour == "#" {
			input = 1
		}

		terminated = computer.Run(&output, [2]int{input, input}, true)
		thisOutput := output[len(output)-1]

		if thisOutput == 0 {
			panel.Colour = "."
		} else {
			panel.Colour = "#"
		}

		terminated = computer.Run(&output, [2]int{input, input}, true)
		thisOutput = output[len(output)-1]

		robot.Turn(thisOutput)
		robot.Move()

		panel.Painted = true
		hull[coord] = panel
	}

	paintedHull := make([][]string, robot.maxY+1)
	for i := range paintedHull {
		paintedHull[i] = make([]string, robot.maxX+1)

		// Prefill so the output is a clean grid
		for j := range paintedHull[i] {
			paintedHull[i][j] = "."
		}
	}

	for coord, panel := range hull {
		split := strings.Split(coord, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		paintedHull[y][x] = panel.Colour
	}

	return paintedHull
}

type Panel struct {
	Colour  string
	Painted bool
}

type Robot struct {
	X, Y, Direction, maxX, maxY int
}

func (r *Robot) Turn(dir int) {
	if dir == 0 {
		r.Direction -= 1
	} else if dir == 1 {
		r.Direction += 1
	}

	if r.Direction < 0 {
		r.Direction = 3
	} else if r.Direction > 3 {
		r.Direction = 0
	}
}

func (r *Robot) Move() {
	switch r.Direction {
	case 0:
		r.Y -= 1
	case 1:
		r.X += 1
	case 2:
		r.Y += 1
	case 3:
		r.X -= 1
	}

	if r.X > r.maxX {
		r.maxX = r.X
	}

	if r.Y > r.maxY {
		r.maxY = r.Y
	}
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
