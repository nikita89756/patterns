package main

import (
	. "misis/pkg/animal"
	. "misis/pkg/person"
)

func main() {
	cat := NewCat("Sema", 14, "Long,Fat")
	p1 := NewPerson("Anton", &cat)
	p1.GetCompanionInfo()

	dog := NewDog("Rider", 10, "Big")
	p2 := NewPerson("Lena", &dog)
	p2.GetCompanionInfo()
}
