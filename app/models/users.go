package models

import (
	"log"
	"time"
)

type User struct {
	ID       int
	UUID     string
	Name     string
	Email    string
	PassWord string
	CreateAt time.Time
}

// Userの登録を行うメソッド
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?,?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Userの取得を行う関数
// NOTE: ポインタ型にすると、nilポインタになる
func GetUser(id int) (user User, err error) {
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreateAt,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Userを更新するメソッド
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// Userを削除するメソッド
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
