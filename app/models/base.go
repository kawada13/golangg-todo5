package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"golang_udemy/todo5/config"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

//テーブルを作成するコードを書く

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

// main.goのmain関数が呼び出される前に呼び出されるメソッド
func init() {
	// ドライバーと接続
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	if err != nil {
		log.Fatalln(err)
	}

	// テーブル作成コマンドを作成(ユーザー)
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	// ユーザーテーブル作成を実行
	_, errU := Db.Exec(cmdU)
	if errU != nil {
		log.Fatalln(errU)
	}

	// テーブル作成コマンドを作成(todo)
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	// ユーザーテーブル作成を実行
	_, errT := Db.Exec(cmdT)
	if errT != nil {
		log.Fatalln(errT)
	}
}

// uuidパッケージを使って、uuidを作成
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// ハッシュ値を作成(パスワード用)
func Encypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))

	return cryptext
}
