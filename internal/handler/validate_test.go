package handler_test

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestValidate(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}
	if os.Getenv("DB_HOST") == "localhost" {
		t.Skip("Tests will be skipped because DB_HOST is not set to docker postgres endpoint")
	}
	// ... rest of your test code ...
}
