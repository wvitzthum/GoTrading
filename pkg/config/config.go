package config

import (
	"GoTrading/models"

	"github.com/spf13/viper"
)

// GetConfig returns read config or error form passed path
func GetConfig(path, filename string) (*models.Config, error) {
	viper.SetConfigType("yaml")
    viper.SetConfigName(filename)
    viper.AddConfigPath(path)
    if err := viper.ReadInConfig();  err != nil {
        return nil, err
    }
    var config models.Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }
    return &config, nil
}