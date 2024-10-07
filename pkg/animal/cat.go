package animal

import (
	"fmt"
	"math/rand"
)

type Cat struct {
	Animal
}

func NewCat(name string, age uint, shape string) Cat {
	return Cat{Animal{age, shape, name, mew}}

}

func mew() {
	if rand.Intn(200) > 100 {
		fmt.Println("MEW")
	}
}
