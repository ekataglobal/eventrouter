package util

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

// TLSConfig holds the TLS configuration for the Kafka sink producer
type TLSConfig struct {
	Enabled            bool     `yaml:"enabled"`
	RootCAFiles        []string `yaml:"rootCAFiles"`
	CertFile           string   `yaml:"certFile,omitempty"`
	KeyFile            string   `yaml:"keyFile,omitempty"`
	InsecureSkipVerify bool     `yaml:"insecureSkipVerify"`
}

// Get returns a valid *tls.Config based on the TLS Config settings
func (c *TLSConfig) Get() (*tls.Config, error) {
	var rootCAs *x509.CertPool
	var clientCerts []tls.Certificate

	rootCAs = x509.NewCertPool()

	for _, certFile := range c.RootCAFiles {
		caCert, err := ioutil.ReadFile(certFile)
		if err != nil {
			return nil, err
		}

		rootCAs.AppendCertsFromPEM(caCert)
	}

	if c.CertFile != "" && c.KeyFile != "" {
		cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
		if err != nil {
			return nil, err
		}

		clientCerts = []tls.Certificate{cert}
	}

	return &tls.Config{
		RootCAs:            rootCAs,
		Certificates:       clientCerts,
		InsecureSkipVerify: c.InsecureSkipVerify,
	}, nil
}
