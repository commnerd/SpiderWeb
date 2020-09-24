package config

import . "github.com/spf13/viper"

func GetConfig() *Viper {
	return GetViper()
}