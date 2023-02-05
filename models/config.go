package models

type Config struct {
	Alpaca AlpacaConfigurations `yaml:"alpaca"`
}

type AlpacaConfigurations struct {
	TradeClient      AlpacaConfig `yaml:"tradeClient"`
	MarketDataClient AlpacaConfig `yaml:"marketDataClient"`
}

type AlpacaConfig struct {
	BaseUrl   string `yaml:"baseUrl"`
	ApiKey    string `yaml:"apiKey"`
	ApiSecret string `yaml:"apiSecret"`
}
