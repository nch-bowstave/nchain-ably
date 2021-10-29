package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

// NewViperConfig will setup and return a new viper based configuration handler.
func NewViperConfig(appname string) *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return &Config{}
}

// WithServer will setup the web server configuration if required.
func (c *Config) WithServer() *Config {
	viper.SetDefault(EnvServerPort, ":8443")
	viper.SetDefault(EnvServerHost, "localhost:8443")
	c.Server = &Server{
		Port:     viper.GetString(EnvServerPort),
		Hostname: viper.GetString(EnvServerHost),
	}
	return c
}

// WithDeployment sets up the deployment configuration if required.
func (c *Config) WithDeployment(appName string) *Config {
	viper.SetDefault(EnvEnvironment, "dev")
	viper.SetDefault(EnvRegion, "test")
	viper.SetDefault(EnvCommit, "test")
	viper.SetDefault(EnvVersion, "test")
	viper.SetDefault(EnvBuildDate, time.Now().UTC())
	viper.SetDefault(EnvMainNet, false)

	c.Deployment = &Deployment{
		Environment: viper.GetString(EnvEnvironment),
		Region:      viper.GetString(EnvRegion),
		Version:     viper.GetString(EnvVersion),
		Commit:      viper.GetString(EnvCommit),
		BuildDate:   viper.GetTime(EnvBuildDate),
		AppName:     appName,
		MainNet:     viper.GetBool(EnvMainNet),
	}
	return c
}

// WithLog sets up and returns log config.
func (c *Config) WithLog() *Config {
	c.Logging = &Logging{Level: viper.GetString(EnvLogLevel)}
	return c
}

// WithBitcoinNode sets up and returns bitcoin node configuration.
func (c *Config) WithBitcoinNode() *Config {
	c.Node = &BitcoinNode{
		Host:     viper.GetString(EnvNodeHost),
		Port:     viper.GetInt(EnvNodePort),
		Username: viper.GetString(EnvNodeUser),
		Password: viper.GetString(EnvNodePassword),
		UseSSL:   viper.GetBool(EnvNodeSSL),
	}
	return c
}

// WithAbly sets up the header client with the type of
// syncing we wish to do.
func (c *Config) WithAbly() *Config {
	c.Ably = &Ably{
		Key:        viper.GetString(EnvAblyKey),
		Username:   viper.GetString(EnvAblyUsername),
		MaxMessage: viper.GetInt64(EnvAblyMaxMessage),
	}
	return c
}
