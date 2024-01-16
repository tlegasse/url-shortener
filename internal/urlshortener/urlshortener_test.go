package urlshortener

import (
	"testing"
)

func TestSetup(t *testing.T) {
	const (
		hostname = "http://localhost"
		port = "8080"
	)

	Setup(hostname, port)

	// Test if the Shortener contains the hostname and port as specified
	if Shortener.Url != hostname {
		t.Errorf("Shortener.Hostname is not initialized")
	}
	if Shortener.Port != port {
		t.Errorf("Shortener.Port is not initialized")
	}
}

func TestSetupRoutes(*testing.T) {

}

func TestRandStringRunes(*testing.T) {

}

func TestCleanUrl(*testing.T) {

}

func TestShorten(*testing.T) {

}

func TestRedirect(*testing.T) {
}

