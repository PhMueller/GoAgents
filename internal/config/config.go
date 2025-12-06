package config

import (
	"fmt"
	"log"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (Config, error) {
	var config Config

	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error unmarshalling config %v", err)
	}

	if err := v.Unmarshal(&config, func(decoderConfig *mapstructure.DecoderConfig) {
		// set ErrorUnset to make sure that each variable in the config has been retrieved from the env
		decoderConfig.ErrorUnset = true
	}); err != nil {
		log.Fatalf("Error unmarshalling config %v", err)
	}
	fmt.Println("Config:", config)
	return config, nil
}
