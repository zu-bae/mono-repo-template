package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadUsesFrontendDirFromEnvFile(t *testing.T) {
	tempDir := t.TempDir()
	prevWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer func() {
		if err := os.Chdir(prevWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	if err := os.WriteFile(filepath.Join(tempDir, ".env"), []byte("FRONTEND_DIR=./custom-build\n"), 0o600); err != nil {
		t.Fatalf("write .env: %v", err)
	}

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.ClientDir != "./custom-build" {
		t.Fatalf("expected ClientDir to be ./custom-build, got %q", cfg.ClientDir)
	}
}
