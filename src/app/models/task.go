package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	CreatedAt time.Time
}

func (user *User) Tasks() ([]Task, error) {
	tasks := make([]Task, 0)
	rows, err := db.Query("SELECT id,uuid,body,user_id,created_at FROM tasks WHERE user_id=$1", user.Id)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		task := Task{}
		if err := rows.Scan(&task.Id, &task.Uuid, &task.Body, &task.UserId, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (user *User) CreateTask(body string) (*Task, error) {
	task := &Task{}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	if err := db.QueryRow("INSERT INTO tasks (uuid, body, user_id, created_at) VALUES ($1,$2,$3,$4) RETURNING id, uuid,body,user_id,created_at", id, body, user.Id, time.Now()).
		Scan(&task.Id, &task.Uuid, &task.Body, &task.UserId, &task.CreatedAt); err != nil {
		return nil, err
	}
	return task, nil
}
