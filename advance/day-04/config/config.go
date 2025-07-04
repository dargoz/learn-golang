package config

import "github.com/spf13/viper"

type Config struct {
	DBSource    string `mapstructure:"DB_SOURCE"`
	GRPCAddress string `mapstructure:"GRPC_ADDRESS" env:"GRPC_ADDRESS" envDefault:"localhost:8080"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
