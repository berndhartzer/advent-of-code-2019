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
	total := 0

	for i := lower; i <= upper; i++ {
		digitsDontDecrease := true
		groupedAdjacent := false

		password := []byte(strconv.Itoa(i))
		counts := map[byte]int{}

		for j := 0; j < len(password); j++ {
			counts[password[j]] += 1

			if j+1 < len(password) && password[j] > password[j+1] {
				digitsDontDecrease = false
			}
		}

		if !digitsDontDecrease {
			continue
		}

		for _, count := range counts {
			if count == 2 {
				groupedAdjacent = true
			}
		}

		if groupedAdjacent {
			total += 1
		}
	}

	return total
}
