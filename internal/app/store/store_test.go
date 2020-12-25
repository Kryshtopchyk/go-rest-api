package store_test

import (
	"os"
	"testing"
)

var (
	databeseURL string
)

func TestMain(m *testing.M) {
	databeseURL = os.Getenv("DATABASE_URL")
	if databeseURL == "" {
		databeseURL = "host=localhost password=1234 user=postgres dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
