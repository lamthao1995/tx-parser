package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"tx-parser/api"
	"tx-parser/config"
	"tx-parser/middleware"
	"tx-parser/repository"
	"tx-parser/scripts"
	"tx-parser/service"
	"tx-parser/utils"
)

func main() {
	// Ensure the log file is closed on exit
	defer utils.CloseLogFile()

	utils.InfoLogger.Println("Starting the server...")

	// Set up application dependencies
	repo := repository.NewMemoryRepository()
	parserService := service.NewParserService(repo)
	handler := api.NewHandler(parserService)

	// Use the logging and recovery middleware
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)
	http.Handle("/", middleware.RecoveryMiddleware(middleware.LoggingMiddleware(mux)))

	// Start the HTTP server
	port := config.AppConfig.ServerPort
	fmt.Println("Server running on http://localhost:" + port)

	// for testing only
	if strings.EqualFold(config.AppConfig.IsNeedTestingData, "true") {
		go scripts.FillData()
	}

	utils.InfoLogger.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
