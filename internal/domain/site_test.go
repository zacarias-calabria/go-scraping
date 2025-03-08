package domain

import (
	"testing"
)

func TestSite_NewSite(t *testing.T) {
	t.Run("Should return error when url is empty", func(t *testing.T) {
		site, err := NewSite("")
		if err == nil {
			t.Error("Expected error when url is empty")
		}
		if err.Error() != "error.site.url_is_empty" {
			t.Errorf("Expected error message 'error.site.url_is_empty', got %v", err.Error())
		}
		if site != nil {
			t.Errorf("Expected nil site when url is empty, got %v", site)
		}
	})

	t.Run("Should create a site successfully", func(t *testing.T) {
		url := "http://domain.com"
		site, err := NewSite(url)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if site == nil {
			t.Errorf("Expected site, got nil")
		}
		if site.URL != url {
			t.Errorf("Expected URL to be %s, got %s", url, site.URL)
		}
	})
}
