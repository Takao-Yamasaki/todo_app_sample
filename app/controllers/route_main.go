package controllers

import (
	"log"
	"net/http"
	"text/template"
)

// topハンドラ
func top(w http.ResponseWriter, r *http.Request) {
	// ファイルを解析
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	// 第二引数をHTMLファイルに渡すことが可能
	t.Execute(w, "Hello")
}
