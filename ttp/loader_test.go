package ttp

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPorblem(t *testing.T) {
	rootDir, _ := os.Getwd()
	path := filepath.Join(filepath.Dir(rootDir), "data", "trivial_0.ttp")
	cities, err := LoadPorblem(path)

	if err != nil {
		t.Error(err)
	}

	if len(cities) != 10 {
		t.Errorf("Failed to laod cities. Total cities loaded %v, expected to load %v", len(cities), 10)
	}
}
