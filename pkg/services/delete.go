package services

import (
	"errors"
	"fmt"

	"github.com/kylerwsm/bounce/internal/database"
	"github.com/kylerwsm/bounce/pkg/models"
)

// DeleteLink deletes an existing redirect link.
func DeleteLink(shortURL string) error {
	db := database.GetDatabase()
	shortLink := models.ShortLink{
		ShortLink: shortURL,
	}

	result := db.Delete(&shortLink)
	if result.Error != nil {
		errMsg := fmt.Sprintf("An error occurred: %s", result.Error.Error())
		return errors.New(errMsg)
	}
	return nil
}
