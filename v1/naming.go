package naming

import (
	"math/rand"
	"time"
)

var names = []string{"Mike", "Tim", "Elisabeth", "Jenifer", "Jianguo"}

func CreateBabyName() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(names))
	return names[index]
}