package handler_test

import (
	"github.com/jexroid/gopi/internal/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogout(t *testing.T) {
	// Create a request recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a request with the identity cookie set
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "identity",
		Value: "some_value",
		Path:  "/",
	})

	// Call the Logout function
	handler.Logout(rr, req)

	// Check if the response status code is 200 (OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}

	// Check if the "cookie deleted" message is in the response body
	expectedBody := "cookie deleted"
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected response body to contain '%s', got '%s'", expectedBody, rr.Body.String())
	}

	// Check if the identity cookie  is deleted
	cookies := rr.Header()["Set-Cookie"]
	t.Log(cookies)
	if len(cookies) != 1 {
		t.Error("Expected 1 Set-Cookie header, got", len(cookies))
	}

	assert.Equal(t, rr.Code, 200, "response was not 200")
}
