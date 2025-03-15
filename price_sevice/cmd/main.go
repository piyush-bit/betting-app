package main

import (
	"log"
	"net/http"
	"price_service/internal/amm"
	"price_service/internal/api"
)

func main() {
	// Initialize AMM calculator
	calculator := amm.NewCalculator(100) // Liquidity parameter of 100

	// Initialize handler
	handler := api.NewPriceHandler(calculator)

	// Setup routes
	http.HandleFunc("/api/prices/calculate", handler.HandlePriceCalculation)

	// Start server
	log.Printf("Starting price service on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}