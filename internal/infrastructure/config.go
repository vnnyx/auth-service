package infrastructure

import "github.com/spf13/viper"

type Config struct {
	RedisHost       string `mapstructure:"REDIS_HOST"`
	RedisPassword   string `mapstructure:"REDIS_PASSWORD"`
	AppPort         string `mapstructure:"APP_PORT"`
	PrivateKey      string `mapstructure:"JWT_PRIVATE_KEY"`
	PublicKey       string `mapstructure:"JWT_PUBLIC_KEY"`
	ServiceAuthPort string `mapstructure:"SERVICE_AUTH_PORT"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
