package internal

import (
	"github.com/c1r5/gormigrate"
	"github.com/c1r5/url-shortener/src/internal/shortener"
	"gorm.io/gorm"
)

func init() {
	gormigrate.Register(gormigrate.Migration{
		Version:     1,
		Description: "Create short_urls table",
		Up: func(db *gorm.DB) error {
			return db.AutoMigrate(&shortener.ShortURL{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&shortener.ShortURL{})
		},
	})
}
