package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kirimi/rb2backend/libs/db"
	"github.com/kirimi/rb2backend/user_service/internal/config"
	"github.com/kirimi/rb2backend/user_service/internal/repository"
	"github.com/kirimi/rb2backend/user_service/internal/server"
	"github.com/kirimi/rb2backend/user_service/migrations"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

const exampleUID = "1X82du5Ou2UZt315HKR8i5o7vAY2"

func main() {
	cfg := config.GetConfig("config.yaml")

	ctx := context.Background()
	conn, err := db.ConnectDB(ctx, &cfg.DB)
	if err != nil {
		log.Panic().Err(err).Msg("Database connection error")
	}
	defer func(conn *sqlx.DB) {
		db.CloseDB(conn)
	}(conn)

	goose.SetBaseFS(migrations.EmbedFs)
	err = goose.SetDialect("postgres")
	if err != nil {
		log.Panic().Err(err).Msg("goose.SetDialect error")
	}
	err = goose.Up(conn.DB, ".")
	if err != nil {
		log.Panic().Err(err).Msg("goose migration error")
	}

	userRepo := repository.NewRepository(conn)

	u, err := userRepo.GetUserByUID(ctx, exampleUID)
	if err != nil {
		log.Warn().Err(err).Msg("Cannot fetch user")
	} else {
		log.Info().Msgf("User %+v", u)
	}

	if err := server.NewGrpcServer().Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")
		return
	}
}
