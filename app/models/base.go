package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo_app_sample/config"

	_ "github.com/mattn/go-sqlite3"
)

// Usersテーブルの作成
var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		uuid string not null unique,
		email string,
		password string,
		created_at datetime)`, tableNameUser)

	Db.Exec(cmdU)
}
