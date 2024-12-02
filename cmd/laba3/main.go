package main

import (
	"fmt"
	"misis/pkg/laba3/builder"
	. "misis/pkg/laba3/singleton"
)



func main(){

	builder := builderMessage.NewBuilder()

	message := builder.SetSender("Nikita").SetDestination("Artem").SetMessage("text","LABA3").Build()

	message.Message.Create()

	message = builder.SetSender("Avanesyan").SetDestination("Artem").SetMessage("photo","/papka/papka2/joto.png").Build()

	message.Message.Create()

	log1 := NewMlogger()
	log2 := NewMlogger()
	fmt.Println(log1 == log2)

	message1:= message.Message.Copy()

	fmt.Println(message.Message == message1)
}