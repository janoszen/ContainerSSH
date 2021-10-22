//go:build windows
// +build windows

package config

// DockerConnectionConfig configures how to connect to dockerd.
type DockerConnectionConfig struct {
	// Host is the docker connect URL
	Host string `json:"host" yaml:"host" default:"npipe:////./pipe/docker_engine"`
	// CaCert is the CA certificate for Docker connection embedded in the configuration in PEM format.
	CaCert string `json:"cacert" yaml:"cacert"`
	// Cert is the client certificate in PEM format embedded in the configuration.
	Cert string `json:"cert" yaml:"cert"`
	// Key is the client key in PEM format embedded in the configuration.
	Key string `json:"key" yaml:"key"`
}
