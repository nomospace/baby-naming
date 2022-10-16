package naming

import (
	"math/rand"
	"time"
)

var boys = []string{"Mike", "Tim", "Jianguo"}
var girls = []string{"Cracie", "Monica",  "Elisabeth", "Jenifer"}


func CreateBabyName(isBoy bool) string {
	rand.Seed(time.Now().UnixNano())

	if isBoy {
		index := rand.Intn(len(boys))
		return boys[index]
	} else {
		index := rand.Intn(len(girls))
		return girls[index]
	}
}
