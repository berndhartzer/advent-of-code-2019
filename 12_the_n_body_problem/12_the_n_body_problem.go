package twelve_the_n_body_problem

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Axis struct {
	X, Y, Z int
}

type Moon struct {
	Position, Velocity Axis
}

func (m *Moon) applyGravity(moon *Moon) {
	if m.Position.X > moon.Position.X {
		m.Velocity.X -= 1
		moon.Velocity.X += 1
	} else if m.Position.X < moon.Position.X {
		m.Velocity.X += 1
		moon.Velocity.X -= 1
	}

	if m.Position.Y > moon.Position.Y {
		m.Velocity.Y -= 1
		moon.Velocity.Y += 1
	} else if m.Position.Y < moon.Position.Y {
		m.Velocity.Y += 1
		moon.Velocity.Y -= 1
	}

	if m.Position.Z > moon.Position.Z {
		m.Velocity.Z -= 1
		moon.Velocity.Z += 1
	} else if m.Position.Z < moon.Position.Z {
		m.Velocity.Z += 1
		moon.Velocity.Z -= 1
	}
}

func (m *Moon) applyVelocity() {
	m.Position.X += m.Velocity.X
	m.Position.Y += m.Velocity.Y
	m.Position.Z += m.Velocity.Z
}

func (m *Moon) getTotalEnergy() int {
	potentialEnergy := Abs(m.Position.X) + Abs(m.Position.Y) + Abs(m.Position.Z)
	kineticEnergy := Abs(m.Velocity.X) + Abs(m.Velocity.Y) + Abs(m.Velocity.Z)
	return potentialEnergy * kineticEnergy
}

func (m *Moon) stringAxis(axis string) string {
	output := ""
	switch axis {
	case "X":
		output = fmt.Sprintf("%d,%d", m.Position.X, m.Velocity.X)
	case "Y":
		output = fmt.Sprintf("%d,%d", m.Position.Y, m.Velocity.Y)
	case "Z":
		output = fmt.Sprintf("%d,%d", m.Position.Z, m.Velocity.Z)
	}
	return output
}

func TheNBodyProblemPartOne(moonPositions []string, simulations int) int {
	moons := []*Moon{}

	re := regexp.MustCompile(`[-]?\d[\d]*[\.]?[\d{2}]*`)
	for _, position := range moonPositions {
		nums := re.FindAllString(position, -1)
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])

		moon := &Moon{
			Position: Axis{x, y, z},
		}
		moons = append(moons, moon)
	}

	for k := 0; k < simulations; k++ {
		for i, moon := range moons {
			for j := i + 1; j < len(moons); j++ {
				moon.applyGravity(moons[j])
			}
		}

		for _, moon := range moons {
			moon.applyVelocity()
		}
	}

	total := 0
	for _, moon := range moons {
		total += moon.getTotalEnergy()
	}

	return total
}

func TheNBodyProblemPartTwo(moonPositions []string) int {
	moons := []*Moon{}

	re := regexp.MustCompile(`[-]?\d[\d]*[\.]?[\d{2}]*`)
	for _, position := range moonPositions {
		nums := re.FindAllString(position, -1)
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])

		moon := &Moon{
			Position: Axis{x, y, z},
		}
		moons = append(moons, moon)
	}

	xStates := make(map[string]bool)
	yStates := make(map[string]bool)
	zStates := make(map[string]bool)
	x := 0
	y := 0
	z := 0

	i := 0
	for x == 0 || y == 0 || z == 0 {
		xStateStringParts := []string{}
		yStateStringParts := []string{}
		zStateStringParts := []string{}
		for _, moon := range moons {
			xStateStringParts = append(xStateStringParts, moon.stringAxis("X"))
			yStateStringParts = append(yStateStringParts, moon.stringAxis("Y"))
			zStateStringParts = append(zStateStringParts, moon.stringAxis("Z"))
		}

		if x == 0 {
			xString := strings.Join(xStateStringParts, ",")
			_, ok := xStates[xString]
			if ok {
				x = i
			}
			xStates[xString] = true
		}

		if y == 0 {
			yString := strings.Join(yStateStringParts, ",")
			_, ok := yStates[yString]
			if ok {
				y = i
			}
			yStates[yString] = true
		}

		if z == 0 {
			zString := strings.Join(zStateStringParts, ",")
			_, ok := zStates[zString]
			if ok {
				z = i
			}
			zStates[zString] = true
		}

		for i, moon := range moons {
			for j := i + 1; j < len(moons); j++ {
				moon.applyGravity(moons[j])
			}
		}

		for _, moon := range moons {
			moon.applyVelocity()
		}

		i += 1
	}

	return LCM(x, y, z)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
