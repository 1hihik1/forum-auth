package app

import (
	"github.com/1hihik1/forum-auth/internal/delivery/gin"
	"github.com/1hihik1/forum-auth/internal/repository"
	"github.com/1hihik1/forum-auth/internal/usecase"
	"github.com/1hihik1/forum-auth/pkg/database"
	"github.com/1hihik1/forum-auth/pkg/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
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
		logger.Logger.Fatal("Ошибка запуска сервера на порту :8080",
			zap.Error(err),
			zap.String("app", "database"))
	}
	logger.Logger.Info("Микросервис стартует на порту :8080")
}
