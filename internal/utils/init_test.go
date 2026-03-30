package utils

import (
	"os"
	"testing"
	"github.com/bbc18/decypharr/internal/config"
)

func TestMain(m *testing.M) {
	// Setup config path for tests
	tmpDir, err := os.MkdirTemp("", "decypharr_test_*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	config.SetConfigPath(tmpDir)

	// Run tests
	os.Exit(m.Run())
}
