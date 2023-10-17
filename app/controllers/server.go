package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"todo_app_sample/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// NOTE: errの場合は、panic errorとなる
	templates := template.Must(template.ParseFiles(files...))
	// defineを使って宣言したtemplateを読み込む場合には、明示的に宣言する必要がある
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	// NOTE:/static/というディレクトリは無いので、/static/をパスから取り除く
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// topハンドラを呼び出す
	http.HandleFunc("/", top)

	// NOTE: 第二引数は、通常、デフォルトのマルチプレクサを使用するため、nil
	// デフォルトで登録されていないページにアクセスしたら、page not found を返してくれる
	log.Printf("server start at %v", config.Config.Port)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
