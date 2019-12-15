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

func (m *Moon) String() string {
	return fmt.Sprintf("%d,%d,%d,%d,%d,%d", m.Position.X, m.Position.Y, m.Position.Z, m.Velocity.X, m.Velocity.Y, m.Velocity.Z)
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

	states := make(map[string]bool)

	repeated := false
	for !repeated {
		for i, moon := range moons {
			for j := i + 1; j < len(moons); j++ {
				moon.applyGravity(moons[j])
			}
		}

		for _, moon := range moons {
			moon.applyVelocity()
		}

		moonString := []string{}
		for _, moon := range moons {
			moonString = append(moonString, moon.String())
		}
		fullMoonString := strings.Join(moonString, ",")
		// fmt.Println(fullMoonString)

		_, ok := states[fullMoonString]
		if !ok {
			states[fullMoonString] = true
		} else {
			repeated = true
		}

		// repeated = true
	}

	return len(states)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
