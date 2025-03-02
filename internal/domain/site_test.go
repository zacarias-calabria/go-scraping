package domain

import (
	"testing"
)

func TestNewSite(t *testing.T) {
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
}
