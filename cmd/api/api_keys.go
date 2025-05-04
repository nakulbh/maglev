package main

import "net/http"

func (app *application) requestHasInvalidAPIKey(r *http.Request) bool {
	key := r.URL.Query().Get("key")
	return app.isInvalidAPIKey(key)
}

func (app *application) isInvalidAPIKey(key string) bool {
	// This is a placeholder. In a real application, you would:
	// - Check against keys stored in a database
	// - Check against keys in your configuration
	// - Potentially validate expiration, permissions, etc.

	if key == "" {
		return true
	}

	// For example, checking against keys stored in your app config:
	validKeys := app.config.apiKeys
	for _, validKey := range validKeys {
		if key == validKey {
			return false
		}
	}

	return true
}
