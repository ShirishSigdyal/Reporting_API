package main

import (
	"log"
	"net/http"
	"reporting-api/config"
	"reporting-api/src/controllers"
	"reporting-api/db"
	"reporting-api/src/repositories"
	"reporting-api/src/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	dbConn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize repository and services
	repo := repositories.NewReportRepository(dbConn)
	service := services.NewReportService(repo)
	controller := controllers.NewReportsController(service)

	// Set up routes and start server
	router := controller.SetupRoutes()

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
