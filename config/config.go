package config

import (
	"fmt"
	"time"

	validator "github.com/theflyingcodr/govalidator"
)

// Environment variable constants.
const (
	EnvServerPort     = "server.port"
	EnvServerHost     = "server.host"
	EnvEnvironment    = "env.environment"
	EnvMainNet        = "env.mainnet"
	EnvRegion         = "env.region"
	EnvVersion        = "env.version"
	EnvCommit         = "env.commit"
	EnvBuildDate      = "env.builddate"
	EnvLogLevel       = "log.level"
	EnvNodeHost       = "node.host"
	EnvNodePort       = "node.port"
	EnvNodeUser       = "node.username"
	EnvNodePassword   = "node.password"
	EnvNodeSSL        = "node.usessl"
	EnvAblyKey        = "ably.key"
	EnvAblyUsername   = "ably.username"
	EnvAblyMaxMessage = "ably.maxmessage"

	LogDebug = "debug"
	LogInfo  = "info"
	LogError = "error"
	LogWarn  = "warn"
)

// Config returns strongly typed config values.
type Config struct {
	Logging    *Logging
	Server     *Server
	Deployment *Deployment
	Node       *BitcoinNode
	Ably       *Ably
}

// Validate will check config values are valid and return a list of failures
// if any have been found.
func (c *Config) Validate() error {
	vl := validator.New()
	return vl.Err()
}

// Deployment contains information relating to the current
// deployed instance.
type Deployment struct {
	Environment string
	AppName     string
	Region      string
	Version     string
	Commit      string
	BuildDate   time.Time
	MainNet     bool
}

// IsDev determines if this app is running on a dev environment.
func (d *Deployment) IsDev() bool {
	return d.Environment == "dev"
}

func (d *Deployment) String() string {
	return fmt.Sprintf("Environment: %s \n AppName: %s\n Region: %s\n Version: %s\n Commit:%s\n BuildDate: %s\n",
		d.Environment, d.AppName, d.Region, d.Version, d.Commit, d.BuildDate)
}

// Logging contains log configuration.
type Logging struct {
	Level string
}

// Server contains all settings required to run a web server.
type Server struct {
	Port     string
	Hostname string
}

// BitcoinNode config params for connecting to a bitcoin node.
type BitcoinNode struct {
	Host     string
	Port     int
	Username string
	Password string
	UseSSL   bool
}

type Ably struct {
	Key        string
	Username   string
	MaxMessage int64
}
