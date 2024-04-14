package main

import (
	"context"

	"github.com/watariRyo/baby-map-server/config"
	"github.com/watariRyo/baby-map-server/domain/repository"
	"github.com/watariRyo/baby-map-server/handler"
	"github.com/watariRyo/baby-map-server/infra/db"
	"github.com/watariRyo/baby-map-server/infra/firebase"
	"github.com/watariRyo/baby-map-server/infra/redis"
	"github.com/watariRyo/baby-map-server/server"
	"github.com/watariRyo/baby-map-server/usecase"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	println(cfg.Db.Host)

	// infra
	dbConnectionManager := db.NewConnectionManager(
		db.Username(cfg.Db.Username),
		db.Password(cfg.Db.Password),
		db.Host(cfg.Db.Host),
		db.Port(cfg.Db.Port),
		db.Schema(cfg.Db.Schema),
		db.DebugMode(cfg.Db.DebugMode),
	)
	conn, err := dbConnectionManager.Open()
	if err != nil {
		panic(err)
	}

	redisClient, err := redis.NewRedisClient(&cfg.Redis)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	firebaseApp, err := firebase.InitFirebaseApp(ctx, cfg.Server)

	allRepository := &repository.AllRepository{
		DBConnection:  conn,
		DBTransaction: db.Transaction,
		RedisClient:   redisClient,
		FirebaseApp:   firebaseApp,
		// UsersRepository: db.NewUsersRepository(),
	}

	usecase := usecase.NewUseCase(allRepository, cfg)
	handler := handler.NewHandler(usecase)

	server := server.NewServer(ctx, cfg, handler, allRepository)
	server.Run()
}
