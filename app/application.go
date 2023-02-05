package app

import (
	"GoTrading/models"

	"golang.org/x/exp/slog"

	"GoTrading/pkg/alpaca"
)

func StartApp(config *models.Config, logger *slog.Logger) error {
	tradeCL, err := alpaca.InitAlpacaClient(
		config.Alpaca.TradeClient.ApiKey, config.Alpaca.TradeClient.ApiSecret, *logger)
	if err != nil {
		return err
	}
	mclient := alpaca.MarketDataClient{}
	err = mclient.Init(config.Alpaca.MarketDataClient.ApiKey, config.Alpaca.MarketDataClient.ApiSecret, *logger)
	if err != nil {
		return err
	}
	client := alpaca.AlpacaClient{TradeClient: tradeCL, MarketClient: mclient}

	client.MarketClient.GetTradeData()

	//alpaca.GetTradeData()
	return err
}