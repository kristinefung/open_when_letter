package main

import (
	"log"

	"open_when_letter/config"
	"open_when_letter/internal/handler"
	"open_when_letter/internal/model"
	"open_when_letter/internal/repository"
	"open_when_letter/internal/service"
	"open_when_letter/pkg/database"
	"open_when_letter/route"
)

func main() {
	cfg := config.Load()

	db, err := database.Init(cfg)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	if err := db.AutoMigrate(&model.LetterBox{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// DI
	boxRepo := repository.NewLetterBoxRepository(db)
	boxService := service.NewLetterBoxService(boxRepo)
	boxHandler := handler.NewLetterBoxHandler(boxService)

	// Setup routes
	router := route.SetupRouter(boxHandler)

	log.Printf("Server starting on %s", cfg.ServerPort)
	router.Run(cfg.ServerPort)
}
