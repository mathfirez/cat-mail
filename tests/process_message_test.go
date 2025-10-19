package tests

import (
	"cat-mail/src/controllers"
	"testing"
)

func TestMessageTrim(t *testing.T) {
	message := "    MESSAGE WITHOUT TRIM    "
	expected := "MESSAGE WITHOUT TRIM"
	msg := controllers.ProcessMessage(message)
	if expected != msg {
		t.Errorf(`ProcessMessage = %q, expected: %#q, nil`, msg, expected)
	}
}
