package observer

import "fmt"

type EmailListener struct{
	User User
}

func (e EmailListener)Update(news string){
	fmt.Printf("Отправка новости '%s' на почту: %s пользователю: %s",news,e.User.email,e.User.name)
}

