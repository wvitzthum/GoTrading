package strategies

import (
	"reflect"
	"testing"
)

func TestBollingerBands(t *testing.T) {
	// Test case 1
	prices := []float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	window := 5
	numStdDev := 2.0
	expectedUpper := []float64{42, 52, 62, 72, 82, 92, 102, 112, 122}
	expectedLower := []float64{18, 8, -2, -12, -22, -32, -42, -52, -62}
	upper, lower := BollingerBands(prices, window, numStdDev)
	if !reflect.DeepEqual(upper, expectedUpper) || !reflect.DeepEqual(lower, expectedLower) {
		t.Errorf("BollingerBands(%v, %d, %f) = (%v, %v), expected (%v, %v)", prices, window, numStdDev, upper, lower, expectedUpper, expectedLower)
	}

	// Test case 2
	prices = []float64{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}
	window = 3
	numStdDev = 1.0
	expectedUpper = []float64{93, 83, 73, 63, 53, 43, 33}
	expectedLower = []float64{77, 67, 57, 47, 37, 27, 17}
	upper, lower = BollingerBands(prices, window, numStdDev)
	if !reflect.DeepEqual(upper, expectedUpper) || !reflect.DeepEqual(lower, expectedLower) {
		t.Errorf("BollingerBands(%v, %d, %f) = (%v, %v), expected (%v, %v)", prices, window, numStdDev, upper, lower, expectedUpper, expectedLower)
	}

	// Test case 3
	prices = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	window = 2
	numStdDev = 0.5
	expectedUpper = []float64{2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}
	expectedLower = []float64{1.5, 0.5, -0.5, -1.5, -2.5, -3.5, -4.5, -5.5}
	upper, lower = BollingerBands(prices, window, numStdDev)
	if !reflect.DeepEqual(upper, expectedUpper) || !reflect.DeepEqual(lower, expectedLower) {
		t.Errorf("BollingerBands(%v, %d, %f) = (%v, %v), expected (%v, %v)", prices, window, numStdDev, upper, lower, expectedUpper, expectedLower)
	}
}