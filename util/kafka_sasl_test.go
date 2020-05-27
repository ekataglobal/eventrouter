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
