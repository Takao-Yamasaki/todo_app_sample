package controllers

import (
	"net/http"
)

// topハンドラ
func top(w http.ResponseWriter, r *http.Request) {

	generateHTML(w, "Hello", "layout", "top")

	/*
		// ファイルを解析
		t, err := template.ParseFiles("app/views/templates/top.html")
		if err != nil {
			log.Fatalln(err)
		}
		// 第二引数をHTMLファイルに渡すことが可能
		t.Execute(w, "Hello")
	*/
}
