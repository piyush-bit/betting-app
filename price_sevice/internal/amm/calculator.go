package amm

import (
	"math"
	"sync"
)

type Calculator struct {
	liquidityParameter float64
	mu                sync.RWMutex
}

func NewCalculator(liquidityParameter float64) *Calculator {
	return &Calculator{
		liquidityParameter: liquidityParameter,
	}
}

func (c *Calculator) CalculatePrice(yesShares, noShares float64, outcome string, amount float64) float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if outcome == "yes" {
		return c.calculateYesPrice(yesShares, noShares, amount)
	}
	return c.calculateNoPrice(yesShares, noShares, amount)
}

func (c *Calculator) calculateYesPrice(yesShares, noShares, amount float64) float64 {
	b := c.liquidityParameter
	return math.Exp((yesShares+amount)/b) / (math.Exp((yesShares+amount)/b) + math.Exp(noShares/b))
}

func (c *Calculator) calculateNoPrice(yesShares, noShares, amount float64) float64 {
	b := c.liquidityParameter
	return math.Exp((noShares+amount)/b) / (math.Exp(yesShares/b) + math.Exp((noShares+amount)/b))
}

func (c *Calculator) CalculatePriceImpact(yesShares, noShares float64, outcome string, amount float64) (currentPrice, impactedPrice, impact float64) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	currentPrice = c.CalculatePrice(yesShares, noShares, outcome, 0)
	impactedPrice = c.CalculatePrice(yesShares, noShares, outcome, amount)
	impact = math.Abs((impactedPrice - currentPrice) / currentPrice * 100)

	return currentPrice, impactedPrice, impact
}