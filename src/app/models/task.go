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
	UpdatedAt time.Time
}

func (user *User) Tasks() ([]Task, error) {
	tasks := make([]Task, 0)
	rows, err := db.Query("SELECT id,uuid,body,user_id,created_at,updated_at FROM tasks WHERE user_id=$1 ORDER BY updated_at DESC ", user.Id)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		task := Task{}
		if err := rows.Scan(&task.Id, &task.Uuid, &task.Body, &task.UserId, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func FindTask(id int) (*Task, error) {
	task := &Task{UserId: -1}
	if err := db.QueryRow("SELECT id, uuid, body, user_id, created_at,updated_at FROM tasks WHERE id=$1", id).
		Scan(&task.Id, &task.Uuid, &task.Body, &task.UserId, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}
	return task, nil
}

func (user *User) CreateTask(body string) (*Task, error) {
	task := &Task{}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()
	if err := db.QueryRow("INSERT INTO tasks (uuid, body, user_id, created_at,updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id, uuid,body,user_id,created_at,updated_at", id, body, user.Id, timestamp, timestamp).
		Scan(&task.Id, &task.Uuid, &task.Body, &task.UserId, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}
	return task, nil
}

func (task *Task) UpdateTask(body string) error {
	if err := db.QueryRow("UPDATE tasks SET body=$1, updated_at=$2 WHERE id=$3 RETURNING body,updated_at", body, time.Now(), task.Id).
		Scan(&task.Body, &task.UpdatedAt); err != nil {
		return err
	}
	return nil
}
