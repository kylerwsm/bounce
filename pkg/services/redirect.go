package services

import (
	"errors"
	"fmt"
)

// RedirectFrom gets the redirect target of the provided short URL.
func RedirectFrom(shortURL string) (string, error) {
	urlMap := map[string]string{
		"pgadmin": "https://pgadmin.buirrito.com",
	}

	if val, ok := urlMap[shortURL]; ok {
		return val, nil
	}
	errMsg := fmt.Sprintf("The short url \"%s\" is invalid", shortURL)
	return "", errors.New(errMsg)
}
