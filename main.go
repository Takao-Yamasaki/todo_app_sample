package main

import (
	"fmt"
	"log"
	"todo_app_sample/app/models"
	"todo_app_sample/config"
)

func main() {
	// 環境変数読み込みのテスト
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	// ログ出力のテスト
	log.Println("test")

	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(1)

	fmt.Println(u)
}
