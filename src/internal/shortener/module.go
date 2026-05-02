package shortener

import (
	"encoding/json"
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

func writeJSON(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func errorResponse(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, ErrorResponse{Error: message})
}

func (m *ShortenerModule) ShortenURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		originalURL := r.FormValue("url")
		if originalURL == "" {
			errorResponse(w, http.StatusBadRequest, "URL is required")
			return
		}

		code := internal.GenerateCode(originalURL)
		shortURL := ShortURL{
			Code:        code,
			OriginalURL: originalURL,
			ExpireAt:    time.Now().Add(24 * time.Hour).Unix(),
		}

		if err := m.db.Create(&shortURL).Error; err != nil {
			errorResponse(w, http.StatusInternalServerError, "Failed to create short URL")
			return
		}

		writeJSON(w, http.StatusCreated, ShortenURLResponse{
			Code:        code,
			OriginalURL: originalURL,
			ShortURL:    "http://" + r.Host + "/" + code,
			ExpireAt:    shortURL.ExpireAt,
		})
	default:
		w.Header().Set("Allow", http.MethodPost)
		errorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (m *ShortenerModule) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		errorResponse(w, http.StatusNotFound, "Short URL not found")
		return
	}

	switch r.Method {
	case http.MethodGet:
		code := r.URL.Path[1:]
		shortURL := ShortURL{}
		result := m.db.Where("code = ?", code).First(&shortURL, "expire_at > ?", time.Now().Unix())
		if result.Error != nil {
			errorResponse(w, http.StatusNotFound, "Short URL not found")
			return
		}

		http.Redirect(w, r, shortURL.OriginalURL, http.StatusFound)
		m.db.Model(&shortURL).UpdateColumn("clicks", gorm.Expr("clicks + ?", 1))

	default:
		w.Header().Set("Allow", http.MethodGet)
		errorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
