package main

import (
	"fmt"
	Mlogger "misis/pkg/laba3/singleton"
	"misis/pkg/laba4/adapter"
	messagehandler "misis/pkg/laba4/decorator"
	sender "misis/pkg/laba4/facade"
	userproxy "misis/pkg/laba4/proxy"
)

func main() {
	

	ldecorator :=  messagehandler.CreateMessageHandler(3)

	ldecorator.ProcessMessage(":) http://virus :)")

	fmt.Println(ldecorator.GetString())
	Mlogger:= Mlogger.NewMlogger()
	Mlogger.LevelDebug()

	data1 := database.NewUsers1(Mlogger)
	data1.Connect()
	data2 := database.NewUsers2(Mlogger)
	data2.Connect()

	

	users:=userproxy.NewUserProxy(Mlogger,data1)
	User:=users.GetData(1)
	fmt.Println(User)
	User = users.GetData(1)
	fmt.Println(User)

	
	users2:=userproxy.NewUserProxy(Mlogger,data2)
	User=users2.GetData(1)
	fmt.Println(User)
	User = users2.GetData(1)
	fmt.Println(User)

	sender := sender.NewSenderFacade(Mlogger,data1,3)
	sender.Send(1,3,"hello bratishka :)","text")

}