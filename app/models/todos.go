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

// todo作成

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(
		cmd,
		content,
		u.ID,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// todo取得(一つ)

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `select * from todos where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)

	if err != nil {
		log.Fatalln(err)
	}
	return todo, err
}
