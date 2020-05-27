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
	"errors"

	"github.com/Shopify/sarama"
)

var (
	ErrUnsupportedSASLMechanism = errors.New("unsupported SASL mechanism")
	ErrNoUsernameAndPassword    = errors.New("did not supply a SASL user and password")
)

// SaslConfig holds the SASL configuration options
type SaslConfig struct {
	Enabled   bool   `yaml:"enabled"`
	Mechanism string `yaml:"mechanism,omitempty"`
	User      string `yaml:"user,omitempty"`
	Password  string `yaml:"password,omitempty"`
}

// Apply applies the SASL authentication to a Kafka producer config object
func (s SaslConfig) Apply(conf *sarama.Config) error {
	if s.Enabled == false {
		return nil
	}
	if s.User == "" || s.Password == "" {
		return ErrNoUsernameAndPassword
	}
	if s.Mechanism == "" {
		s.Mechanism = sarama.SASLTypePlaintext
	}

	switch s.Mechanism {
	case sarama.SASLTypeSCRAMSHA256:
		conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &XDGSCRAMClient{HashGeneratorFcn: SHA256}
		}
	case sarama.SASLTypeSCRAMSHA512:
		conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &XDGSCRAMClient{HashGeneratorFcn: SHA512}
		}
	default:
		return ErrUnsupportedSASLMechanism
	}

	conf.Net.SASL.Enable = true
	conf.Net.SASL.User = s.User
	conf.Net.SASL.Password = s.Password
	conf.Net.SASL.Mechanism = sarama.SASLMechanism(s.Mechanism)

	return nil
}
