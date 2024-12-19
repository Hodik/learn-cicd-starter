package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 1234567890")
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal(err)
	}
	expected := "1234567890"
	if apiKey != expected {
		t.Fatalf("expected %s, got %s", expected, apiKey)
	}
}

func TestGetAPIKey_NoHeader(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected %s, got %s", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "1234567890")
	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err.Error() != "malformed authorization header" {
		t.Fatalf("expected %s, got %s", "malformed authorization header", err)
	}
}
