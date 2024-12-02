package animal

type Animal struct {
	age   uint
	shape string
	name  string
}

func NewAnimal(age uint, shape string, name string) Animal {
	return Animal{age, shape, name}
}

func (a *Animal) GetName() string {
	return a.name
}

func (a *Animal) GetAge() uint {
	return a.age
}

func (a *Animal) SetName(name string) {
	a.name = name
}

func (a *Animal) MakeSound() {

}
