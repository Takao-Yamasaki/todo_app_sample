package main

import (
	"fmt"
	"log"
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
}
