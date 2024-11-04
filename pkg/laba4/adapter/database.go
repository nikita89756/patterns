package database

import (
	"fmt"
	Mlogger "misis/pkg/laba3/singleton"
	"misis/pkg/laba4/user"
	"strconv"
)


type Database interface{
	Connect()
	GetUser(int) *user.User
}

type Users1 struct{
	logger *Mlogger.Mlogger
	data map[int]string
}

func NewUsers1(logger *Mlogger.Mlogger) *Users1{
	return &Users1{logger,nil}
}

type Users2 struct{
	logger *Mlogger.Mlogger
	data map[string]string
}

func NewUsers2(logger *Mlogger.Mlogger) *Users2{
	return &Users2{logger,nil}
}

func (p *Users1)Connect(){
	p.data = map[int]string{
		1 : "Andrey",
		2 : "Anton",
		3 : "Arina",
	}
	fmt.Println(p.logger.LogInfo("Подключение к Users1"))

}
func (p *Users1)GetUser(id int) *user.User{

	if name,ok := p.data[id] ; ok{
	return &user.User{name}
}

	return nil
}

func (p *Users2)Connect(){
	p.data = map[string]string{
		"1" : "Andrey",
		"2" : "Anton",
		"3" : "Arina",
	}
	fmt.Println(p.logger.LogInfo("Подключение к Users2"))

}
func (p *Users2)GetUser(id int) *user.User{

	if name,ok := p.data[strconv.Itoa(id)] ; ok{
	return &user.User{name}
}

	return nil
}