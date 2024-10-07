package animal

import (
	"fmt"
)

type Dog struct {
	Animal
}

func NewDog(name string, age uint, shape string) Dog {
	return Dog{Animal{age, shape, name, waf}}

}

func waf() {

	fmt.Println("ARF")

}
