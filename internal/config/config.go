package config

import "github.com/spf13/viper"

type Config struct {
	AccessToken        string
	ConsumerKey        string
	Host               string
	Path               string
	PrometheusEndpoint string
	PrometheusUsername string
	PrometheusPassword string
}

func GetConfig() Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	conf := Config{
		AccessToken:        viper.GetString("ACCESS_TOKEN"),
		ConsumerKey:        viper.GetString("CONSUMER_KEY"),
		Host:               viper.GetString("HOST"),
		Path:               viper.GetString("PATH"),
		PrometheusEndpoint: viper.GetString("PROMETHEUS_ENDPOINT"),
		PrometheusUsername: viper.GetString("PROMETHEUS_USERNAME"),
		PrometheusPassword: viper.GetString("PROMETHEUS_PASSWORD"),
	}
	return conf
}
