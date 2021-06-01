package client

import (
	"testing"
)

func TestNew(t *testing.T) {
	client := New("asdf", "qwer")

	if client.Token != "asdf" {
		t.Error("Expected token to be 'asdf'")
	}

	if client.Secret != "qwer" {
		t.Error("Expected secret to be 'qwer'")
	}
}
