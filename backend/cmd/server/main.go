package main

import (
	"context"
	"log"
	"os"

	"github.com/Kenya-i/linu-quest/internal/handler"
	"github.com/Kenya-i/linu-quest/internal/middleware"
	"github.com/Kenya-i/linu-quest/internal/repository"
	"github.com/Kenya-i/linu-quest/internal/usecase"
	"github.com/gin-contrib/cors"
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
	stageRepo := repository.NewStageRepository(db)
	stageUsecase := usecase.NewStageUsecase(stageRepo)
	stageHandler := handler.NewStageHandler(stageUsecase)

	questionRepo := repository.NewQuestionRepository(db)
	questionUsecase := usecase.NewQuestionUsecase(questionRepo)
	questionHandler := handler.NewQuestionHandler(questionUsecase)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	r.GET("/me", middleware.AuthMiddleware([]byte(jwtSecret)), userHandler.Me)
	r.GET("/stages", stageHandler.List)
	r.GET("/stages/:id/questions", questionHandler.ListByStage)
	r.POST("/questions/:id/answer", questionHandler.CheckAnswer)

	r.Run(":8080")

}
