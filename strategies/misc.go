package strategies

import ()

// The SMA function takes as input an array of prices and a window size,
// and returns an array of the simple moving averages of the price data.
// The function uses a loop to calculate the sum of prices within the
// window for each point in time, and then divides the sum by the window
// size to calculate the moving average. The moving average is then
// appended to the output array.
func SMA(prices []float64, window int) []float64 {
	var sma []float64
	for i := window - 1; i < len(prices); i++ {
		var sum float64
		for j := i; j > i-window; j-- {
			sum += prices[j]
		}
		sma = append(sma, sum/float64(window))
	}
	return sma
}

// The EMA function takes as input an array of prices,
// a window size, and a smoothing factor alpha. The function
// uses the formula for exponential moving average to calculate
// the EMA of the price data. The EMA is initialized with the
// first price value and the subsequent values are calculated using
// the formula (price[i]-ema[i-1]) * alpha + ema[i-1]. Finally,
// the function returns the EMA values for the window size specified.
func EMA(prices []float64, window int) []float64 {
	var ema []float64
	k := 2.0 / float64(window+1)
	ema = append(ema, prices[0])
	for i := 1; i < len(prices); i++ {
		ema = append(ema, ema[i-1]+k*(prices[i]-ema[i-1]))
	}
	return ema
}

// The RSI function takes as input an array of prices and
// a window size, and returns an array of the relative
// strength index (RSI) values. The function first calculates
// the arrays of gains and losses based on the difference
// between consecutive prices. It then uses the SMA function
// to calculate the average gain and average loss over the
// specified window. The RSI values are then calculated using
// the formula 100 - 100 / (1 + (avgGain / avgLoss)) and returned.
func RSI(prices []float64, window int) []float64 {
	var rsi []float64
	var gains, losses []float64
	for i := 1; i < len(prices); i++ {
		change := prices[i] - prices[i-1]
		if change > 0 {
			gains = append(gains, change)
			losses = append(losses, 0.0)
		} else {
			gains = append(gains, 0.0)
			losses = append(losses, -change)
		}
	}
	avgGain := SMA(gains, window)[window-1]
	avgLoss := SMA(losses, window)[window-1]
	for i := window - 1; i < len(prices)-1; i++ {
		if avgLoss == 0 {
			rsi = append(rsi, 100.0)
		} else {
			rs := avgGain / avgLoss
			rsi = append(rsi, 100.0-(100.0/(1.0+rs)))
		}
		avgGain = (avgGain*float64(window-1) + gains[i]) / float64(window)
		avgLoss = (avgLoss*float64(window-1) + losses[i]) / float64(window)
	}
	return rsi
}

