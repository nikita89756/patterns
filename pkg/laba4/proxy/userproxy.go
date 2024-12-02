package userproxy

import (
	"fmt"
	Mlogger "misis/pkg/laba3/singleton"
	database "misis/pkg/laba4/adapter"
	"misis/pkg/laba4/user"
)

type UserProxy struct {
	logger   Mlogger.Logger
	database database.Database
	cahed    map[int]user.User
}

func (u *UserProxy) GetData(id int) *user.User {
	if user, ok := u.cahed[id]; ok {
		fmt.Println(u.logger.LogInfo("Данные взяты из кэша"))
		return &user
	}
	fmt.Println(u.logger.LogInfo("Данные взяты из базы"))
	user := u.database.GetUser(id)
	if user != nil {
		u.cahed[id] = *user
		return user
	}
	return nil
}

func NewUserProxy(logger Mlogger.Logger, database database.Database) *UserProxy {
	return &UserProxy{logger, database, make(map[int]user.User)}
}
