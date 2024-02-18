package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type UserConfig struct {
	MaxGwei        int     `yaml:"max_gwei"`
	UseProxy       bool    `yaml:"use_proxy"`
	NeedDelayAcc   bool    `yaml:"need_delay_acc"`
	DelayAccMin    int     `yaml:"delay_acc_min"`
	DelayAccMax    int     `yaml:"delay_acc_max"`
	TelegramAlerts bool    `yaml:"telegram_alerts"`
	BotToken       string  `yaml:"bot_token"`
	ChatID         int     `yaml:"chat_id"`
	NeedOkx        bool    `yaml:"need_okx"`
	OkxValueMin    float64 `yaml:"okx_value_min"`
	OkxValueMax    float64 `yaml:"okx_value_max"`
	OxkAPIKey      string  `yaml:"oxk_apiKey"`
	OxkSecret      string  `yaml:"oxk_secret"`
	OxkPassword    string  `yaml:"oxk_password"`
}

func ReadSettings(filepath string) UserConfig {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error reading config file: ", err)
		return UserConfig{}
	}

	var config UserConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal("Error decoding YAML: ", err)
		return UserConfig{}
	}

	return config
}
