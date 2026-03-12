package deploy

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"blog/internal/config"
	"blog/internal/generator"
)

func Deploy(inputDir, outputDir string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	if cfg.Host == "" {
		return fmt.Errorf("deploy config missing: set host in ~/.config/blog.conf")
	}
	if cfg.Path == "" {
		return fmt.Errorf("deploy config missing: set path in ~/.config/blog.conf")
	}

	fmt.Println("Building site...")
	if err := generator.Build(inputDir, outputDir); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	// scp -P {port} -r {output_dir}/* {host}:{path}/
	src := filepath.Join(outputDir, "*")
	dst := fmt.Sprintf("%s:%s/", cfg.Host, cfg.Path)

	fmt.Printf("Deploying to %s:%s...\n", cfg.Host, cfg.Path)

	// Use shell expansion for the glob
	args := []string{"-P", cfg.Port, "-r", src, dst}
	fmt.Printf("Running: scp %v\n", args)

	cmd := exec.Command("sh", "-c", fmt.Sprintf("scp -P %s -r %s/* %s:%s/",
		cfg.Port, outputDir, cfg.Host, cfg.Path))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("scp failed: %w", err)
	}

	fmt.Println("Deploy complete.")
	return nil
}
