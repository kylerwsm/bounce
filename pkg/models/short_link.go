package models

import (
	"time"
)

// ShortLink represents columns that stores our short links.
type ShortLink struct {
	ShortLink    string `gorm:"primaryKey"`
	OriginalLink string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
