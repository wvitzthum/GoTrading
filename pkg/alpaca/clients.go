package alpaca

import (
	"fmt"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"golang.org/x/exp/slog"
)

type AlpacaClient struct {
	TradeClient *alpaca.Client
	MarketClient MarketDataClient
}

func InitAlpacaClient(apiKeyEnv, apiSecretEnv string, logger slog.Logger) (*alpaca.Client, error) {
	key, secret := os.Getenv(apiKeyEnv), os.Getenv(apiSecretEnv)
	if key == "" || secret == "" {
		return nil, fmt.Errorf("Failed to retrieve env vars")
	}
	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:   key,
		APISecret: secret,
		BaseURL:   "https://paper-api.alpaca.markets",
	})
	// Check for client to be working properly
	_, err := client.GetAccount()
	if err != nil {
		logger.Info("Successfully connected client")
	}
	return client, err
}

func InitMarketClient(apiKeyEnv, apiSecretEnv string, logger slog.Logger) (*marketdata.Client, error) {
	key, secret := os.Getenv(apiKeyEnv), os.Getenv(apiSecretEnv)
	if key == "" || secret == "" {
		return nil, fmt.Errorf("Failed to retrieve env vars")
	}
	client := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:   key,
		APISecret: secret,
	})
	// Check for client to be working properly
	client.GetNews(marketdata.GetNewsRequest{
		Symbols:    []string{"AAPL", "TSLA"},
		Start:      time.Date(2021, 4, 3, 0, 0, 0, 0, time.UTC),
		End:        time.Date(2021, 4, 4, 5, 0, 0, 0, time.UTC),
		TotalLimit: 1,
		PageLimit:  1,
	})
	return client, nil
}
