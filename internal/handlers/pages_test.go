package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	ErrDisplayingPage = errors.New("error displaying page")
)

func TestHomepage(t *testing.T) {
	handler := NewPagesHandler()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler.Homepage(w, req)

	t.Run("should return 200 OK status", func(t *testing.T) {
		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200 OK but got %d", w.Code)
		}
	})

	t.Run("should display title correctly", func(t *testing.T) {
		expected := "Home"

		if !contains(w.Body.String(), expected) {
			t.Error(ErrDisplayingPage)
		}
	})
}

// Helper function to check if the response contains expected content
func contains(response, expected string) bool {
	return strings.Contains(response, expected)
}
