package strategies
/*
This example implements a Dual Moving Average Crossover strategy,
which uses two moving averages to determine whether to buy, sell,
or hold a stock. The prices array represents a historical list of
stock prices, the shortWindow variable determines the size of the
short moving average window, and the longWindow variable determines
the size of the long moving average window.
For each iteration, the code calculates both the short and long
moving averages of the prices in their respective windows.

If the short moving average is greater than the long moving
average and the current price is above the short moving average,
the strategy signals to sell. If the short moving average is less
than the long moving average and the current price is below the
short moving average, the strategy signals to buy. If neither
condition is met, the strategy signals to hold.

Note that this is just one example of a Dual Moving Average Crossover
strategy and there are many variations and modifications that can be
made to fit different trading objectives and market conditions.
*/
import (
	"fmt"
)

func DualMovingAverage(prices []float64, shortWindow, longWindow int) {
	for i := longWindow - 1; i < len(prices); i++ {
		var shortSum float64
		for j := i; j > i-shortWindow; j-- {
			shortSum += prices[j]
		}
		shortMA := shortSum / float64(shortWindow)

		var longSum float64
		for j := i; j > i-longWindow; j-- {
			longSum += prices[j]
		}
		longMA := longSum / float64(longWindow)

		price := prices[i]

		if shortMA > longMA && price > shortMA {
			fmt.Println("Sell at price:", price)
		} else if shortMA < longMA && price < shortMA {
			fmt.Println("Buy at price:", price)
		} else {
			fmt.Println("Hold at price:", price)
		}
	}
}