package services

import (
	"errors"
	"fmt"

	"github.com/kylerwsm/bounce/internal/database"
	"github.com/kylerwsm/bounce/pkg/models"
)

// CreateLink creates a new redirect link.
func CreateLink(shortURL, originalLink string) error {
	db := database.GetDatabase()
	shortLink := models.ShortLink{
		ShortLink:    shortURL,
		OriginalLink: originalLink,
	}

	result := db.Create(&shortLink)
	if result.Error != nil {
		errMsg := fmt.Sprintf("An error occurred: %s", result.Error.Error())
		return errors.New(errMsg)
	}
	return nil
}
