package config

import "github.com/spf13/viper"

type Config struct {
	DBSource    string `mapstructure:"DB_SOURCE" env:"DB_SOURCE" envDefault:"postgresql://postgres@localhost:5432/postgres?sslmode=disable"`
	GRPCAddress string `mapstructure:"GRPC_ADDRESS" env:"GRPC_ADDRESS" envDefault:"localhost:8080"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
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
