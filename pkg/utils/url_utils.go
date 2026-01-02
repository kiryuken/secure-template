package utils

// URL utilities
// Helper functions for URL parsing and building

import (
	"net/url"
	"strings"
)

// BuildURL builds URL from base and path
func BuildURL(base, path string) (string, error) {
	// TODO: Implement URL building
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	baseURL.Path = strings.TrimRight(baseURL.Path, "/") + "/" + strings.TrimLeft(path, "/")
	return baseURL.String(), nil
}

// AddQueryParams adds query parameters to URL
func AddQueryParams(urlStr string, params map[string]string) (string, error) {
	// TODO: Implement query parameter addition
	// - Parse URL
	// - Add parameters
	// - Return encoded URL
	return "", nil
}

// ExtractDomain extracts domain from URL
func ExtractDomain(urlStr string) (string, error) {
	// TODO: Implement domain extraction
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

// IsValidURL validates URL format
func IsValidURL(urlStr string) bool {
	// TODO: Implement URL validation
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}

// SanitizeURL sanitizes URL for safety
func SanitizeURL(urlStr string) string {
	// TODO: Implement URL sanitization
	// - Remove dangerous protocols (javascript:, data:)
	// - Validate scheme (http, https only)
	return urlStr
}
