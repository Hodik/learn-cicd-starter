package auth

import (
	"errors"
	"net/http"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {

	return "", ErrNoAuthHeaderIncluded
}
