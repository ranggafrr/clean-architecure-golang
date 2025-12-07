package main

import (
	"clean-architecture/config"
	httpHandler "clean-architecture/internal/delivery/http"
	"clean-architecture/internal/repository"
	"clean-architecture/internal/usecase"
	"clean-architecture/pkg/database"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	db, err := database.NewMySQL(cfg.Database.DSN())
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := httpHandler.NewUserHandler(userUsecase)

	router := httpHandler.NewRouter(userHandler)

	log.Printf("Server running on port %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}
