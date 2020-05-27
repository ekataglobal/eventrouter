/*
Copyright 2017 The Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
