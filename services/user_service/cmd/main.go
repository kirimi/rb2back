package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirimi/rb2backend/libs/db"
	"github.com/kirimi/rb2backend/user_service/internal/config"
	"github.com/kirimi/rb2backend/user_service/internal/repository"
	"github.com/kirimi/rb2backend/user_service/migrations"
	"github.com/pressly/goose/v3"
	"log"
)

const exampleUID = "1X82du5Ou2UZt315HKR8i5o7vAY2"

func main() {
	cfg := config.GetConfig("config.yaml")

	ctx := context.Background()
	conn, err := db.ConnectDB(ctx, &cfg.DB)
	if err != nil {
		log.Panic(err)
	}
	defer func(conn *sqlx.DB) {
		db.CloseDB(conn)
	}(conn)

	goose.SetBaseFS(migrations.EmbedFs)
	err = goose.SetDialect("postgres")
	if err != nil {
		log.Panic(err)
	}
	err = goose.Up(conn.DB, ".")
	if err != nil {
		log.Panic(err)
	}

	userRepo := repository.NewRepository(conn)

	u, err := userRepo.GetUserByUID(ctx, exampleUID)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("User %+v", u)
}
