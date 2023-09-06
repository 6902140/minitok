package initialize

import (
	"github.com/spf13/viper"
	"minitok/cmd/publish/global"
)

func Viper(path string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&global.Configs); err != nil {
		return nil, err
	}
	if err := parseDuration(); err != nil {
		return nil, err
	}
	return v, nil
}