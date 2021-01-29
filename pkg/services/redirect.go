package services

import (
	"errors"
	"fmt"

	"github.com/kylerwsm/bounce/internal/database"
	"github.com/kylerwsm/bounce/pkg/models"
	"gorm.io/gorm"
)

// RedirectFrom gets the redirect target of the provided short URL.
func RedirectFrom(shortURL string) (string, error) {
	db := database.GetDatabase()
	var shortLink models.ShortLink
	result := db.Take(&shortLink, "short_link = ?", shortURL)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errMsg := fmt.Sprintf("The short url \"%s\" is invalid", shortURL)
		return "", errors.New(errMsg)
	}
	return shortLink.OriginalLink, nil
}
