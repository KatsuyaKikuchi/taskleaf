package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Id        int
	Uuid      string
	UserId    int
	CreatedAt time.Time
}

func (user *User) CreateSession() (*Session, error) {
	session := &Session{}
	if err := db.QueryRow("SELECT id,uuid,user_id,created_at FROM sessions WHERE user_id=$1", user.Id).
		Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt); err != nil {
		// 存在しないならば作る
		stmt, err := db.Prepare("INSERT INTO sessions (uuid, user_id, created_at) VALUES ($1,$2,$3) RETURNING id,uuid,user_id,created_at")
		if err != nil {
			return nil, err
		}
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		err = stmt.QueryRow(id.String(), user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return session, nil
}

func FindSession(sessionUuid string) (*Session, error) {
	session := &Session{}
	if err := db.QueryRow("SELECT id,uuid,user_id,created_at FROM sessions WHERE uuid=$1", sessionUuid).
		Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt); err != nil {
		return nil, err
	}
	return session, nil
}

func (session *Session) DeleteSession() error {
	stmt, err := db.Prepare("DELETE FROM sessions WHERE uuid=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.Uuid)
	return err
}
