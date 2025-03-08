package domain

import (
	"testing"
)

func TestLink_NewLink(t *testing.T) {
	t.Run("Should return error when URL is empty", func(t *testing.T) {
		URL := ""
		content := "test content"
		attributes := map[string]string{"class": "test-class"}
		link, err := NewLink(URL, content, attributes)
		if err == nil {
			t.Error("Expected error when URL is empty")
		}
		if err.Error() != "error.link.url_is_empty" {
			t.Errorf("Expected 'error.link.url_is_empty' error, got %v", err)
		}
		if link != nil {
			t.Errorf("Expected no link, got %v", link)
		}
	})

	t.Run("Should return error when content is empty", func(t *testing.T) {
		URL := "http://domain.com"
		content := ""
		attributes := map[string]string{"class": "test-class"}
		link, err := NewLink(URL, content, attributes)
		if err == nil {
			t.Error("Expected error when content is empty")
		}
		if err.Error() != "error.link.content_is_empty" {
			t.Errorf("Expected 'error.link.content_is_empty' error, got %v", err)
		}
		if link != nil {
			t.Errorf("Expected no link, got %v", link)
		}
	})

	t.Run("Should create a link successfully", func(t *testing.T) {
		URL := "http://domain.com"
		content := "test content"
		attributes := map[string]string{"class": "test-class"}
		link, err := NewLink(URL, content, attributes)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if link == nil {
			t.Errorf("Expected link, got nil")
		}
		if link.URL != URL {
			t.Errorf("Expected URL to be %s, got %s", URL, link.URL)
		}
		if link.Content != content {
			t.Errorf("Expected content to be %s, got %s", content, link.Content)
		}
		if len(link.Attributes) != len(attributes) {
			t.Errorf("Expected attributes length to be %d, got %d", len(attributes), len(link.Attributes))
		}
		for key, value := range attributes {
			if link.Attributes[key] != value {
				t.Errorf("Expected attribute %s to be %s, got %s", key, value, link.Attributes[key])
			}
		}
	})
}
