package internal

import "strings"

func NormalizeURL(url string) string {
	// lowercase the URL and trim whitespace
	url = strings.TrimSpace(strings.ToLower(url))

	// add http:// if missing
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	return url
}

func GenerateCode(url string) string {
	normalizedURL := NormalizeURL(url)
	hash := Hash(normalizedURL)
	code := EncodeBase62(hash)
	return code[:6] // return first 6 characters for the short code
}
