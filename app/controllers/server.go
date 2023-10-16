package controllers

import (
	"log"
	"net/http"
	"todo_app_sample/config"
)

func StartMainServer() error {
	// topハンドラを呼び出す
	http.HandleFunc("/", top)

	// NOTE: 第二引数は、通常、デフォルトのマルチプレクサを使用するため、nil
	// デフォルトで登録されていないページにアクセスしたら、page not found を返してくれる
	log.Printf("server start at %v", config.Config.Port)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
