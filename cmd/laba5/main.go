package main

import (
	"fmt"
	Mlogger "misis/pkg/laba3/singleton"
	handler "misis/pkg/laba5/chainOfResponsibility"
	news "misis/pkg/laba5/observer"
	payment "misis/pkg/laba5/strategy"
)

func main(){
	logger:= Mlogger.NewMlogger()
	logger.LevelDebug()
	paymentCard:=payment.NewPaymentCard(123,23,124,"sber",*logger)

	paymentCard.Pay(100)

	paymentCard = payment.NewPaymentCard(777,24,333,"tbank",*logger)

	paymentCard.Pay(666)

	user:=news.NewUser("ALBLAK52","email@mail.ru","+7(952)812")
	newsPublisher:= news.NewNewsPublisher()

	smsListener:=news.SmsListener{User:user}
	emailListener:=news.EmailListener{User:user}

	newsPublisher.Subscribe(smsListener)
	newsPublisher.Subscribe(emailListener)

	newsPublisher.NotifyAll("Это второй")

	text := `
	{
		"name": "John Doe",
		"age": 30,
		"email": "john.doe@example.com"
	}
	`
	text2 :=`
	<person>
		<name>John Doe</name>
		<age>30</age>
		<email>john.doe@example.com</email>
	</person>`
	text3:=`
name = "John Doe"
age = 30
[address]
street = "123 Main St"
city = "Anytown"
state = "CA"
`

	base:=&handler.BaseParseHandler{}
	json:=&handler.JSONHandler{}
	json.SetNext(base)
	toml:=&handler.TomlHandler{}
	toml.SetNext(json)
	xml:=&handler.XmlHandler{}
	xml.SetNext(toml)
	
	t:=xml.Handle(text2)
	fmt.Println(t)
	t=xml.Handle(text)
	fmt.Println(t)
	t=xml.Handle(text3)
	fmt.Println(t)
}