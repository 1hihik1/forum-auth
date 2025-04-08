package app

import (
	"github.com/DrusGalkin/forum-auth-grpc/internal/delivery/gin"
	"github.com/DrusGalkin/forum-auth-grpc/internal/repository"
	"github.com/DrusGalkin/forum-auth-grpc/internal/usecase"
	"github.com/DrusGalkin/forum-auth-grpc/pkg/database"
	_ "github.com/lib/pq"
	"log"
)

func Run() {
	db, err := database.NewSQLiteConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	router := gin.SetupRouter(userUseCase)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
