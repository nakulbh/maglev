package main

import (
	"testing"
)

func testBlankKeyIsInvalid(t *testing.T) {
	app := &application{
		config: config{
			apiKeys: []string{"key"},
		},
	}
	result := app.isInvalidAPIKey("")

	if result != false {
		t.Error("isInvalidAPIKey('') = true")
	}
}
