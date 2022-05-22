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

	// row = Db.QueryRow(cmd, id)
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

// todo取得(複数)

func GetTodos() (todos []Todo, err error) {
	cmd := `select * from todos`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}
