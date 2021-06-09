package models

import (
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (session *Session) FindUser() (*User, error) {
	if session == nil {
		return nil, nil
	}
	user := &User{}
	if err := db.QueryRow("SELECT id, uuid, name, email FROM users WHERE id=$1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return user, nil
}

func FindUser(email string, password string) (*User, error) {
	user := &User{}
	if err := db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email=$1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(name string, email string, password string) (*User, error) {
	var (
		stmt *sql.Stmt
		user *User
		err  error
	)
	if stmt, err = db.Prepare("INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1,$2,$3,$4,$5) RETURNING id, uuid,name,email,created_at"); err != nil {
		return nil, err
	}
	defer stmt.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}

	user = &User{}
	err = stmt.QueryRow(id.String(), name, email, string(hash), time.Now()).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
