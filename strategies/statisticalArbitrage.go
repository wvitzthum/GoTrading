package strategies

import (
	"fmt"
	"math"
)

/*
In this example, the SMA function is used to calculate the simple
moving average of the price data for two currency pairs (A and B).
The expected ratio of the two currencies is calculated as the ratio
of their moving averages, and the current ratio is calculated as the
ratio of their current prices. If the difference between the expected
and current ratios is greater than 0.05, the trader would execute a
trade to profit from the price differential, either buying currency A
and selling currency B or vice versa. If the difference is not greater
than 0.05, the trader would do nothing.
*/

func StatisticalArbitrage(pricesA, pricesB []float64, window int) {
	smaA := SMA(pricesA, window)
	smaB := SMA(pricesB, window)
	for i := window; i < len(pricesA); i++ {
		expectedRatio := smaA[i-window] / smaB[i-window]
		currentRatio := pricesA[i] / pricesB[i]
		if math.Abs(currentRatio-expectedRatio) > 0.05 {
			if currentRatio > expectedRatio {
				fmt.Println("Buy A, Sell B at price", pricesA[i], pricesB[i])
			} else {
				fmt.Println("Sell A, Buy B at price", pricesA[i], pricesB[i])
			}
		} else {
			fmt.Println("Do nothing at price", pricesA[i], pricesB[i])
		}
	}
}
