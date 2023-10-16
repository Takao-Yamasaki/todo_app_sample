package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// Todoの作成をするメソッド
func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
			content,
			user_id,
			created_at) values (?, ?, ?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Todoのデータを取得する関数
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

// Todoのリストを取得するメソッド
func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, nil
}

// 特定のユーザーのTodoのリストを取得するメソッド
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, nil
}

// Todoのデータを更新するメソッド
func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// TODO: 仮実装
// Todoのデータを削除するメソッド
func (u *User) DeleteTodo(id int) (err error) {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// TODO: 次はサーバーから
// 構成の変更
// MySQLに変更
// Gormに変更
// Sessionが出てこなかったら、Session管理入れてみてもいいかも
// TODO: Dockerfileの作成
// TODO: WebコンテナとDBコンテナを作成(docker-comppose.yml)
// TODO: EC2でもいいかもしれない
// TODO: ECSにプッシュ
