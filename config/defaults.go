package config

import (
	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault(EnvHeaderType, "node")

	// Node defaults
	viper.SetDefault(EnvNodeHost, "node1")
	viper.SetDefault(EnvNodePort, 28332)
	viper.SetDefault(EnvNodeUser, "bitcoin")
	viper.SetDefault(EnvNodePassword, "bitcoin")
}
