package shortener

import (
	"net/http"
	"time"

	"github.com/c1r5/url-shortener/src/internal"
	"gorm.io/gorm"
)

type ShortenerModule struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ShortenerModule {
	return &ShortenerModule{db: db}
}

func (m *ShortenerModule) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/shorten", m.ShortenURL)
	mux.HandleFunc("/", m.Redirect)
}

func (m *ShortenerModule) ShortenURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		originalURL := r.FormValue("url")
		if originalURL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		code := internal.GenerateCode(originalURL)
		shortURL := ShortURL{
			Code:        code,
			OriginalURL: originalURL,
			ExpireAt:    time.Now().Add(24 * time.Hour).Unix(),
		}

		if err := m.db.Create(&shortURL).Error; err != nil {
			http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Short URL created: " + code))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (m *ShortenerModule) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		code := r.URL.Path[1:]
		shortURL := ShortURL{}
		result := m.db.Where("code = ?", code).First(&shortURL, "expire_at > ?", time.Now().Unix())
		if result.Error != nil {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, shortURL.OriginalURL, http.StatusFound)
		m.db.Model(&shortURL).UpdateColumn("clicks", gorm.Expr("clicks + ?", 1))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
