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
	"testing"

	"github.com/Shopify/sarama"
)

func TestApplyUnknownSASLMechanism(t *testing.T) {
	conf := &sarama.Config{}

	saslCfg := SaslConfig{
		Enabled:   true,
		Mechanism: "foo",
		User:      "user",
		Password:  "password",
	}

	err := saslCfg.Apply(conf)
	if err != ErrUnsupportedSASLMechanism {
		t.Errorf("Expected '%v', got '%v'", ErrUnsupportedSASLMechanism, err)
	}
}

func TestApplyNoUsername(t *testing.T) {
	conf := &sarama.Config{}

	saslCfg := SaslConfig{
		Enabled: true,
	}

	err := saslCfg.Apply(conf)
	if err != ErrNoUsernameAndPassword {
		t.Errorf("Expected '%v', got '%v'", ErrNoUsernameAndPassword, err)
	}
}
