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

// Todoの全件データを取得するメソッド
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
