package person

import (
	"fmt"
)

type Alive interface {
	MakeSound()
	GetAge() uint
	GetName() string
}
type Person struct {
	name      string
	companion Alive
}

func NewPerson(name string, companion Alive) Person {
	return Person{name, companion}

}

func (p *Person) GetCompanionInfo() {
	fmt.Printf("My name is %s \nMy pet's name is %s \nAge : %d\n \n ", p.name, p.companion.GetName(), p.companion.GetAge())
	p.companion.MakeSound()
}
