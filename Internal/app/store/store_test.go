package store_test

import (
	"os"
	"testing"
)

var (
	DatabaseURL string
)

func TestMain(m *testing.M) {
	DatabaseURL := os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		DatabaseURL = "host=localhost db_name=stocks_test sslmode=disable"
	}
	os.Exit(m.Run())
}
