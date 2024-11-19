package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirimi/rb2backend/libs/db"
	"github.com/kirimi/rb2backend/user_service/internal/config"
	"log"
)

func main() {
	cfg := config.GetConfig("config.yaml")

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		log.Panic(err)
	}
	defer func(conn *sqlx.DB) {
		db.CloseDB(conn)
	}(conn)

	err = conn.Ping()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Connection to db established")
}
