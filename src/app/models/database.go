package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var db *sql.DB

func init() {
	var err error
	if db, err = sql.Open("postgres", "dbname=taskleaf sslmode=disable"); err != nil {
		log.Fatal().Msg(err.Error())
	}
	return
}
