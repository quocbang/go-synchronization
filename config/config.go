package config

import "time"

// FlagOptions definition.
type FlagOptions struct {
	Host    string    `long:"host" description:"" default:"localhost" env:"SYNC_HOST"`
	Port    string    `long:"port" description:"" default:"9091" env:"SYNC_PORT"`
	DevMode bool      `long:"dev-mode" description:"" env:"SYNC_DEV_MODE"`
	AuthKey string    `long:"auth-key" description:"" env:"SYNC_AUTH_KEY"`
	Timeout time.Time `long:"timeout" description:"request timeout, need to fetch unit time" default:"1m" env:"SYNC_TIMEOUT"`
	Config  string    `long:"config" env:"SYNC_CONFIG"`
}

// TLSOptionsType definition.
type TLSOptionsType struct {
	Cert   string `long:"tls-cert" description:"path to TLS certificate (PUBLIC). To enable TLS handshake, you must set this value" env:"MESYNC_TLS_CERT"`
	Key    string `long:"tls-key" description:"path to TLS certificate key (PRIVATE), To enable TLS handshake, you must set this value" env:"MESYNC_TLS_KEY"`
	RootCA string `long:"root-ca" description:"path to the root certificate"`
}

// UseTLS check whether use tls mode or not.
func (t TLSOptionsType) UseTLS() bool {
	return t.Cert != "" && t.Key != ""
}

// Config difinition.
type Config struct {
	PostGres *PostGesConfig `yaml:"postgres"`
	Cloud    *CloudConfig   `yaml:"cloud"`
}

// PostGesConfig connection info.
type PostGesConfig struct {
	Schema   string `yaml:"schema"`
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}

// Cloud connection info.
// TODO: needs cloud account.
type CloudConfig struct {
	Schema   string `yaml:"schema"`
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}
