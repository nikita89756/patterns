package observer

type User struct{
	name string
	email string
	phoneNumber string
}

func NewUser(name,email,phoneNumber string)User{
	return User{name,email,phoneNumber}
}