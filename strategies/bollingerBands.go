package strategies

import (
	"math"
)

// BollingerBands returns the Bollinger Bands for a given set of prices, using a
// window size and number of standard deviations for the upper and lower bands.
func BollingerBands(prices []float64, window int, numStdDev float64) ([]float64, []float64) {
	// Calculate the simple moving average (SMA) for the prices
	sma := SMA(prices, window)

	// Calculate the standard deviation for the prices
	stddev := make([]float64, len(prices)-window+1)
	for i := range stddev {
		var sum float64
		for j := i; j < i+window; j++ {
			sum += math.Pow(prices[j]-sma[i], 2)
		}
		stddev[i] = math.Sqrt(sum / float64(window))
	}

	// Calculate the upper and lower Bollinger Bands
	upper := make([]float64, len(sma))
	lower := make([]float64, len(sma))
	for i := range sma {
		upper[i] = sma[i] + numStdDev*stddev[i]
		lower[i] = sma[i] - numStdDev*stddev[i]
	}

	return upper, lower
}