package controllers

import (
	"log"
	"net/http"
	"todo_app_sample/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	// GETメソッドの場合
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
		// POSTメソッドの場合
	} else if r.Method == "POST" {
		// Fromの解析
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			// signup.htmlの属性のnameを渡す
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		// ユーザーを登録
		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
		}
		// トップページにリダイレクト
		http.Redirect(w, r, "/", 302)
	}
}
