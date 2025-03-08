package domain

import "testing"

func TestResource_NewResource(t *testing.T) {
	t.Run("Should return a error when URL is empty", func(t *testing.T) {
		URL := ""
		MIMEType := "text/html"
		referer := "http://dominio.com/referer"
		content := "test"
		resource, err := NewResource(URL, MIMEType, referer, content)
		if err == nil {
			t.Errorf("Expected error when url is empty")
		}
		if err.Error() != "error.resource.url_is_empty" {
			t.Errorf("Expected 'error.resource.url_is_empty' error, got %v", err)
		}
		if resource != nil {
			t.Errorf("Expected no resource, got %v", resource)
		}
	})

	t.Run("Should return a error when MIME type is empty", func(t *testing.T) {
		URL := "http://dominio.com"
		MIMEType := ""
		referer := "http://dominio.com/referer"
		content := "test"
		resource, err := NewResource(URL, MIMEType, referer, content)
		if err == nil {
			t.Error("Expected error when type is empty")
		}
		if err.Error() != "error.resource.mime_type_is_empty" {
			t.Errorf("Expected 'error.resource.mime_type_is_empty' error, got %v", err)
		}
		if resource != nil {
			t.Errorf("Expected no resource, %v got", resource)
		}
	})

	t.Run("Should return a error when referer is empty", func(t *testing.T) {
		URL := "http://dominio.com"
		MIMEType := "text/html"
		referer := ""
		content := "test"
		resource, err := NewResource(URL, MIMEType, referer, content)
		if err == nil {
			t.Error("Expected error when referer is empty")
		}
		if err.Error() != "error.resource.referer_is_empty" {
			t.Errorf("Expected 'error.resource.referer_is_empty' error, got %v", err)
		}
		if resource != nil {
			t.Errorf("Expected no resource, %v got", resource)
		}
	})

	t.Run("Should return a error when content is empty", func(t *testing.T) {
		URL := "http://dominio.com"
		MIMEType := "text/html"
		referer := "http://dominio.com/referer"
		content := ""
		resource, err := NewResource(URL, MIMEType, referer, content)
		if err == nil {
			t.Error("Expected error when content is empty")
		}
		if err.Error() != "error.resource.content_is_empty" {
			t.Errorf("Expected 'error.resource.content_is_empty' error, got %v", err)
		}
		if resource != nil {
			t.Errorf("Expected no resource, %v got", resource)
		}
	})

	t.Run("Should create a resource successfully", func(t *testing.T) {
		URL := "http://dominio.com"
		MIMEType := "text/html"
		referer := "http://dominio.com/referer"
		content := "test"
		resource, err := NewResource(URL, MIMEType, referer, content)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if resource == nil {
			t.Errorf("Expected resource, got nil")
		}
		if resource.URL != URL {
			t.Errorf("Expected URL to be %s, got %s", URL, resource.URL)
		}
		if resource.MIMEType != MIMEType {
			t.Errorf("Expected MIMEType to be %s, got %s", MIMEType, resource.MIMEType)
		}
		if resource.Referer != referer {
			t.Errorf("Expected Referer to be %s, got %s", referer, resource.Referer)
		}
		if resource.Content != content {
			t.Errorf("Expected Content to be %s, got %s", content, resource.Content)
		}
	})
}
