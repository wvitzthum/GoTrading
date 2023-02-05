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

func (m *MarketDataClient)  GetTradeData(symbol string, from, to time.Time) ([]marketdata.Bar, error) {
	return m.client.GetBars(symbol, marketdata.GetBarsRequest{
		TimeFrame: marketdata.NewTimeFrame(1, marketdata.Min),
		Start:      from,
		End:        to,
		TotalLimit: 100,
		PageLimit:  200,
	})
}