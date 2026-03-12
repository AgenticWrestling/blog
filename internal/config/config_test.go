package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFromMissingFile(t *testing.T) {
	cfg, err := LoadFrom(filepath.Join(t.TempDir(), "nonexistent.conf"))
	if err != nil {
		t.Fatalf("expected no error for missing file, got: %v", err)
	}
	if cfg.Port != "22" {
		t.Errorf("default Port = %q, want 22", cfg.Port)
	}
	if cfg.Host != "" {
		t.Errorf("default Host = %q, want empty", cfg.Host)
	}
	if cfg.Path != "" {
		t.Errorf("default Path = %q, want empty", cfg.Path)
	}
}

func TestLoadFromAllFields(t *testing.T) {
	path := filepath.Join(t.TempDir(), "blog.conf")
	os.WriteFile(path, []byte("host=user@example.com\npath=/var/www/blog\nport=2222\n"), 0644)

	cfg, err := LoadFrom(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Host != "user@example.com" {
		t.Errorf("Host = %q, want user@example.com", cfg.Host)
	}
	if cfg.Path != "/var/www/blog" {
		t.Errorf("Path = %q, want /var/www/blog", cfg.Path)
	}
	if cfg.Port != "2222" {
		t.Errorf("Port = %q, want 2222", cfg.Port)
	}
}

func TestLoadFromIgnoresComments(t *testing.T) {
	path := filepath.Join(t.TempDir(), "blog.conf")
	os.WriteFile(path, []byte("# this is a comment\nhost=example.com\n# another comment\n"), 0644)

	cfg, err := LoadFrom(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Host != "example.com" {
		t.Errorf("Host = %q, want example.com", cfg.Host)
	}
}

func TestLoadFromIgnoresBlankLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "blog.conf")
	os.WriteFile(path, []byte("\n\nhost=example.com\n\n"), 0644)

	cfg, err := LoadFrom(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Host != "example.com" {
		t.Errorf("Host = %q, want example.com", cfg.Host)
	}
}

func TestLoadFromDefaultPort(t *testing.T) {
	path := filepath.Join(t.TempDir(), "blog.conf")
	os.WriteFile(path, []byte("host=example.com\n"), 0644)

	cfg, err := LoadFrom(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Port != "22" {
		t.Errorf("Port = %q, want default 22", cfg.Port)
	}
}

func TestLoadFromIgnoresUnknownKeys(t *testing.T) {
	path := filepath.Join(t.TempDir(), "blog.conf")
	os.WriteFile(path, []byte("host=example.com\nunknown_key=value\n"), 0644)

	cfg, err := LoadFrom(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Host != "example.com" {
		t.Errorf("Host = %q, want example.com", cfg.Host)
	}
}
