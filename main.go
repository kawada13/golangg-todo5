package main

import (
	"fmt"
	"golang_udemy/todo5/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	// ユーザー作成
	// u := models.User{}
	// u.Name = "test3"
	// u.Email = "test3@test.com"
	// u.PassWord = "testtest3"
	// fmt.Println(u)
	// u.CreateUser()

	u, _ := models.GetUser(1)
	fmt.Println(u)
}
