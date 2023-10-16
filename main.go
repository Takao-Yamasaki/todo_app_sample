package main

import (
	"fmt"
	"todo_app_sample/app/controllers"
	"todo_app_sample/app/models"
)

func main() {
	fmt.Println(models.Db)

	// サーバーの起動
	controllers.StartMainServer()
}
