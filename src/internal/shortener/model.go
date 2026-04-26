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
