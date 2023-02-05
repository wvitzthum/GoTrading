package app

import (
	"GoTrading/models"
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"golang.org/x/exp/slog"

	"GoTrading/pkg/alpaca"
	"GoTrading/pkg/badger"
)

type Application struct {
	Clients alpaca.AlpacaClient
	Badger *badger.BadgerClient
	logger *slog.Logger
}

func (a *Application) RetrieveAndStoreData() error {
	symbol := "AAPL"
	bars, err := a.Clients.MarketClient.GetTradeData(
		symbol,
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		return err
	}
	
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	for _, v := range bars {
		key := fmt.Sprintf("%s.%s", "AAPL", v.Timestamp.Format("2006-01-02|15:04"))
		err := encoder.Encode(v)
		if err != nil {
			a.logger.Error("Failed to encode bars", err)
			continue
		}
		a.Badger.InsertValue(key, buf.String())
	}
	return nil
}

func InitApp(config *models.Config, logger *slog.Logger) error {
	app := &Application{}
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
	app.Clients = alpaca.AlpacaClient{TradeClient: tradeCL, MarketClient: mclient}
	app.Badger, err = badger.InitBadger(logger, config.BadgerPath)
	if err != nil {
		return err
	}
	return err
}

func (a *Application) Start() error {
	a.RetrieveAndStoreData()
	a.Badger.RetrieveValuesForKey("AAPL")
	return nil
}