package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Host string
	Path string
	Port string
}

func Load() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot determine home dir: %w", err)
	}
	return LoadFrom(filepath.Join(home, ".config", "blog.conf"))
}

func LoadFrom(confPath string) (*Config, error) {
	cfg := &Config{Port: "22"}

	f, err := os.Open(confPath)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("opening config: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		switch key {
		case "host":
			cfg.Host = val
		case "path":
			cfg.Path = val
		case "port":
			cfg.Port = val
		}
	}
	return cfg, scanner.Err()
}

func InputDir() string {
	if v := os.Getenv("BLOG_INPUT_DIR"); v != "" {
		return v
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Documents", "blog")
}

func OutputDir() string {
	if v := os.Getenv("BLOG_OUTPUT_DIR"); v != "" {
		return v
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Documents", "blog-output")
}
