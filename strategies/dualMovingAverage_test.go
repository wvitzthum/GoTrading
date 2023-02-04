package strategies

import "testing"

func TestDualMovingAverage(t *testing.T) {
	prices := []float64{100, 105, 102, 110, 107, 109, 108}
	shortWindow := 3
	longWindow := 5

	DualMovingAverage(prices, shortWindow, longWindow)
}
