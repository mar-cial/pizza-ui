package adapters

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockHandler struct{}

func (m *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// TestNewHTTPServer verifies that NewHTTPServer correctly sets up the server with the provided port and handler.
func TestNewHTTPServer(t *testing.T) {
	handler := &mockHandler{}
	port := ":8080"

	server := NewHTTPServer(port, handler)

	// Use assertions to check if the server's properties are set correctly
	assert.Equal(t, port, server.(*httpServer).server.Addr)
	assert.Equal(t, handler, server.(*httpServer).server.Handler)
}

// TestHTTPServerStart verifies that the Start method correctly starts the server and handles shutdowns.
func TestHTTPServerStart(t *testing.T) {
	handler := &mockHandler{}
	port := ":8081" // Use a non-standard port for testing

	server := NewHTTPServer(port, handler)

	// Start the server in a goroutine to simulate it running in the background
	go func() {
		// We expect an error because no actual client will connect to it in this test
		err := server.Start()
		if err != nil && err != http.ErrServerClosed {
			t.Errorf("Expected server to close with error 'http.ErrServerClosed', got: %v", err)
		}
	}()

	// Wait for a short period to simulate server running
	time.Sleep(100 * time.Millisecond)

	t.Run("confirm server started", func(t *testing.T) {
		// Try to connect to the server to confirm it started (optional)
		resp, err := http.Get("http://localhost" + port)
		assert.Nil(t, err, "Expected no error when connecting to test server")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status OK from mock handler")
		resp.Body.Close()
	})

	t.Run("shutdown server", func(t *testing.T) {
		// Shutdown the server after the test to avoid lingering resources
		err := server.(*httpServer).server.Close()
		assert.Nil(t, err, "Expected server to close without error")
	})
}
