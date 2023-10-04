package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	MySQL  map[string]MySQLConfig
}

type Server struct {
	Port string `mapstructure:"port"`
}

type MySQLConfig struct {
	Master string `mapstructure:"master"`
	Slave  string `mapstructure:"slave"`
}

func LoadConfiguation() (*Config, error) {

	config := Config{}

	viper.AddConfigPath("config") //配置文件路径
	viper.SetConfigName("config") // 配置文件名称
	viper.SetConfigType("yaml")   // 配置文件类型
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil

}
