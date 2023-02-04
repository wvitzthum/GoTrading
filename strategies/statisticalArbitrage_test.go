package strategies

import "testing"

func TestStatisticalArbitrage(t *testing.T) {
	pricesA := []float64{100, 105, 102, 110, 107, 109, 108}
	pricesB := []float64{90, 95, 97, 100, 95, 100, 105}
	window := 3

	StatisticalArbitrage(pricesA, pricesB, window)
}
