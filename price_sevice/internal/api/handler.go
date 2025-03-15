package api

import (
	"encoding/json"
	"net/http"
	"price_service/internal/amm"
	"price_service/internal/models"
)

type PriceHandler struct {
	calculator *amm.Calculator
}

func NewPriceHandler(calculator *amm.Calculator) *PriceHandler {
	return &PriceHandler{calculator: calculator}
}

func (h *PriceHandler) HandlePriceCalculation(w http.ResponseWriter, r *http.Request) {
	var req models.PriceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch market state (in production, this would come from your database)
	marketState := models.MarketState{
		YesShares: 100, // Example values
		NoShares:  100,
	}

	currentPrice, impactedPrice, impact := h.calculator.CalculatePriceImpact(
		marketState.YesShares,
		marketState.NoShares,
		req.Outcome,
		req.Amount,
	)

	response := models.PriceResponse{
		CurrentPrice:   currentPrice,
		ImpactedPrice: impactedPrice,
		PriceImpact:   impact,
		TotalCost:     impactedPrice * req.Amount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}