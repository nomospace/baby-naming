package naming

import (
	"fmt"
	"math/rand"
)
var boys = []string{"Mike", "Tim", "Jianguo"}
var girls = []string{"Cracie", "Monica",  "Elisabeth", "Jenifer"}


func CreateBabyName(isBoy bool, minLen uint) (string, error) {
	names := cookNames(isBoy)

	for i := 0; i < len(names); i++ {
		if uint(len(names[i])) >= minLen {
			return names[i], nil
		}
	}

	return "", fmt.Errorf("no baby found with min length %v", minLen)
}

func cookNames(isBoy bool) []string {
	var names []string

	if isBoy {
		names = boys
	} else {
		names = girls
	}

	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	return names
}