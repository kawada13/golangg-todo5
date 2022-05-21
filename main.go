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

	//ユーザー取得
	// u, _ := models.GetUser(2)
	// fmt.Println(u)

	// ユーザー更新
	// u.Name = "test1change"
	// u.Email = "test1@example.com"

	// u.UpdateUser()

	// u1, _ := models.GetUser(2)
	// fmt.Println(u1)

	// ユーザー削除
	// まずは対象ユーザーを取得

	u, _ := models.GetUser(2)
	fmt.Println(u)

	// ユーザーの削除
	u.DeleteUser()

	// 実際に削除されているのかを確認
	fmt.Println(u)
}
