package four_secure_container

import (
	"strconv"
)

func SecureContainerPartOne(lower, upper int) int {
	total := 0

	for i := lower; i <= upper; i++ {
		matchingAdjacent := false
		digitsDontDecrease := true

		password := []byte(strconv.Itoa(i))

		for j := 0; j < len(password)-1; j++ {
			if password[j] == password[j+1] {
				matchingAdjacent = true
			}

			if password[j] > password[j+1] {
				digitsDontDecrease = false
			}
		}

		if matchingAdjacent && digitsDontDecrease {
			total += 1
		}
	}

	return total
}

func SecureContainerPartTwo(lower, upper int) int {
	return 0
}
