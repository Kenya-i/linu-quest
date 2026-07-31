package main

import (
	"context"
	"log"
	"os"

	"github.com/Kenya-i/linu-quest/internal/handler"
	"github.com/Kenya-i/linu-quest/internal/repository"
	"github.com/Kenya-i/linu-quest/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = "postgres://linuquest_user:password@localhost:5432/linuquest"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-key"
	}

	db, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, []byte(jwtSecret))
	userHandler := handler.NewUserHandler(userUsecase)

	r := gin.Default()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	r.Run(":8080")

}
