package config

import (
	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault(EnvHeaderType, "node")

	// Node defaults
	viper.SetDefault(EnvNodeHost, "localhost")
	viper.SetDefault(EnvNodePort, 28332)
	viper.SetDefault(EnvNodeUser, "bitcoin")
	viper.SetDefault(EnvNodePassword, "bitcoin")

	viper.SetDefault(EnvLogLevel, "info")

	viper.SetDefault(EnvAblyMaxMessage, 16384)
}
