package models


type User struct{
	ID int `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Orders []Orders `gorm:foreignKey:UserID`
}

type Orders struct{
	ID int `gorm:"primaryKey"`
	ProductName string `gorm:"not null"`
	ProductCount int `gorm:"not null"`
	UserID int `gorm:"not null"`
}