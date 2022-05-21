package models

import (
	"database/sql"
	"fmt"
	"golang_udemy/todo5/config"
	"log"

	// sqllite3のドライバーを手動でインストール
	_ "github.com/mattn/go-sqlite3"
)

//テーブルを作成するコードを書く

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

// main.goのmain関数が呼び出される前に呼び出されるメソッド
func init() {
	// ドライバーと接続
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	if err != nil {
		log.Fatalln(err)
	}

	// テーブル作成コマンドを作成
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	// テーブル作成を実行
	_, err := Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}
}
