package shortener

import "gorm.io/gorm"

type ShortURL struct {
	gorm.Model
	OriginalURL string `gorm:"column:original_url; not null"`
	Code        string `gorm:"column:code;uniqueIndex;not null"`
	Clicks      int    `gorm:"column:clicks;default:0"`
	ExpireAt    int64  `gorm:"column:expire_at;not null"`
}

func (ShortURL) TableName() string {
	return "short_urls"
}

type ShortenURLResponse struct {
	Code        string `json:"code"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	ExpireAt    int64  `json:"expire_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
