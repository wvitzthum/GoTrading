package alpaca

import (
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"golang.org/x/exp/slog"
)
type MarketDataClient struct {
	client *marketdata.Client
	logger *slog.Logger
}

func (m *MarketDataClient) Init(apiKeyEnv, apiSecretEnv string, logger slog.Logger) (error) {
	var err error
	m.client, err = InitMarketClient(apiKeyEnv, apiSecretEnv, logger)
	return err
}

func (m *MarketDataClient)  GetTradeData() ([]marketdata.Bar, error) {
	return m.client.GetBars("AAPL", marketdata.GetBarsRequest{
		TimeFrame: marketdata.NewTimeFrame(1, marketdata.Day),
		Start:      time.Date(2021, 4, 3, 0, 0, 0, 0, time.UTC),
		End:        time.Date(2021, 4, 4, 5, 0, 0, 0, time.UTC),
		TotalLimit: 5,
		PageLimit:  2,
	})
}