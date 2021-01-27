package services

import (
	"errors"
	"fmt"
)

var urlMap = map[string]string{
	"pgadmin": "https://pgadmin.buirrito.com",
}

// RedirectFrom gets the redirect target of the provided short URL.
func RedirectFrom(shortURL string) (string, error) {
	if val, ok := urlMap[shortURL]; ok {
		return val, nil
	}
	errMsg := fmt.Sprintf("The short url \"%s\" is invalid", shortURL)
	return "", errors.New(errMsg)
}
