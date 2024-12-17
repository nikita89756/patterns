package main

import (
	"fmt"
	storage "misis/pkg/laba7/postgres"
)

func main(){
	db := storage.GetDB("localhost","5432","postgres","postgres","lgdb")
	user1:=db.CreateUser("Nikita")
	user2:=db.CreateUser("Anton")
	db.CreateOrder("Варежки",2,user1.ID)
	db.CreateOrder("Нож",3,user2.ID)
	users:=db.GetAllUsersWithOrders()
	for _, user := range users{
		fmt.Printf("User: %s\n",user.Username)
		fmt.Println("Orders:")
		for _,order := range user.Orders{
			fmt.Printf("Product: %s, Count: %d\n",order.ProductName,order.ProductCount)
		}
	}
	db.DeleteUser(user1.ID)
	users = db.GetAllUsersWithOrders()
	for _, user := range users{
		fmt.Printf("User: %s\n",user.Username)
		fmt.Println("Orders:")
		for _,order := range user.Orders{
			fmt.Printf("Product: %s, Count: %d\n",order.ProductName,order.ProductCount)
		}
	}
}