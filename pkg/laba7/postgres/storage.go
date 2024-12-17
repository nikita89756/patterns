package storage

import (
	"fmt"
  "misis/pkg/laba7/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct{
  db *gorm.DB
}

func GetDB(host,port,user,password,dbname string)*Storage{
  connstring := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",host,user,password,dbname,port)
  db,err:=gorm.Open(postgres.Open(connstring),&gorm.Config{Logger: logger.Default.LogMode(logger.Info),
  })
  if err!=nil{
    panic("db connection error")
  }
  db.AutoMigrate(&models.User{},&models.Orders{})
  return &Storage{db:db}

}

func (s *Storage)GetAllUsersWithOrders() []*models.User{
  var users []*models.User
  s.db.Preload("Orders").Find(&users)
  return users
}

func (s *Storage)CreateUser(username string) *models.User{
  user := &models.User{Username: username}
  s.db.Create(user)
  return user
}

func (s *Storage)CreateOrder(productName string , count int , userID int) *models.Orders{
  order := &models.Orders{ProductName: productName,ProductCount: count,UserID: userID}
  s.db.Create(order)
  return order
}

func (s *Storage)DeleteUser(userid int){
  s.db.Where("user_id = ?",userid).Delete(&models.Orders{})
  s.db.Delete(&models.User{},userid)
}