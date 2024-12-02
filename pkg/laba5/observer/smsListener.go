package observer

import "fmt"

type SmsListener struct{
	User User
}

func (e SmsListener)Update(news string){
	fmt.Printf("Отправка новости '%s' по SMS на номер телефона: %s пользователю: %s\n",news,e.User.phoneNumber,e.User.name)
}