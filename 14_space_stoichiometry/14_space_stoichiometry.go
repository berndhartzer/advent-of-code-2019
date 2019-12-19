package fourteen_space_stoichiometry

import (
	"fmt"
	"strings"
)

func SpaceStoichiometryPartOne(reactions []string) int {
	for _, reaction := range reactions {
		split := strings.Split(reaction, "=>")
		fmt.Println(split)

	}

	return 0
}
