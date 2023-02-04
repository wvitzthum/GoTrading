package strategies

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrategies() {
	prices := []float64{100, 105, 102, 110, 107, 109, 108}
	shortWindow := 3
	longWindow := 5

	sma := SMA(prices, shortWindow)
	ema := EMA(prices, longWindow)
	rsi := RSI(prices, longWindow)

	for i := longWindow - 1; i < len(prices); i++ {
		if ema[i] > sma[i-shortWindow+1] && rsi[i-longWindow+1] < 50 {
			fmt.Println("Buy at price", prices[i])
		} else if ema[i] < sma[i-shortWindow+1] && rsi[i-longWindow+1] > 50 {
			fmt.Println("Sell at price", prices[i])
		} else {
			fmt.Println("Do nothing at price", prices[i])
		}
	}
}

func TestRSI(t *testing.T) {
	prices := []float64{100, 105, 103, 108, 110, 115}
	window := 3
	expectedRSI := []float64{50, 56.25, 58.33, 61.54, 66.67}

	rsi := RSI(prices, window)
	if !reflect.DeepEqual(rsi, expectedRSI) {
		t.Errorf("Expected RSI to be %v, but got %v", expectedRSI, rsi)
	}
}

func TestRSIWithEmptyPrices(t *testing.T) {
	prices := []float64{}
	window := 3
	expectedRSI := []float64{}

	rsi := RSI(prices, window)
	if !reflect.DeepEqual(rsi, expectedRSI) {
		t.Errorf("Expected RSI to be %v, but got %v", expectedRSI, rsi)
	}
}

func TestRSIWithWindowGreaterThanPrices(t *testing.T) {
	prices := []float64{100, 105}
	window := 3
	expectedRSI := []float64{}

	rsi := RSI(prices, window)
	if !reflect.DeepEqual(rsi, expectedRSI) {
		t.Errorf("Expected RSI to be %v, but got %v", expectedRSI, rsi)
	}
}


func TestEMA(t *testing.T) {
	prices := []float64{100, 105, 103, 108, 110, 115}
	window := 3
	expectedEMA := []float64{103, 104.67, 106.33, 108.67, 111.67}

	ema := EMA(prices, window)
	if !reflect.DeepEqual(ema, expectedEMA) {
		t.Errorf("Expected EMA to be %v, but got %v", expectedEMA, ema)
	}
}

func TestEMAWithEmptyPrices(t *testing.T) {
	prices := []float64{}
	window := 3
	expectedEMA := []float64{}

	ema := EMA(prices, window)
	if !reflect.DeepEqual(ema, expectedEMA) {
		t.Errorf("Expected EMA to be %v, but got %v", expectedEMA, ema)
	}
}

func TestEMAWithWindowGreaterThanPrices(t *testing.T) {
	prices := []float64{100, 105}
	window := 3
	expectedEMA := []float64{}

	ema := EMA(prices, window)
	if !reflect.DeepEqual(ema, expectedEMA) {
		t.Errorf("Expected EMA to be %v, but got %v", expectedEMA, ema)
	}
}

func TestSMA(t *testing.T) {
	prices := []float64{100, 105, 103, 108, 110, 115}
	window := 3
	expectedSMA := []float64{103, 104, 106, 107, 111}

	sma := SMA(prices, window)
	if !reflect.DeepEqual(sma, expectedSMA) {
		t.Errorf("Expected SMA to be %v, but got %v", expectedSMA, sma)
	}
}

func TestSMAWithEmptyPrices(t *testing.T) {
	prices := []float64{}
	window := 3
	expectedSMA := []float64{}

	sma := SMA(prices, window)
	if !reflect.DeepEqual(sma, expectedSMA) {
		t.Errorf("Expected SMA to be %v, but got %v", expectedSMA, sma)
	}
}

func TestSMAWithWindowGreaterThanPrices(t *testing.T) {
	prices := []float64{100, 105}
	window := 3
	expectedSMA := []float64{}

	sma := SMA(prices, window)
	if !reflect.DeepEqual(sma, expectedSMA) {
		t.Errorf("Expected SMA to be %v, but got %v", expectedSMA, sma)
	}
}