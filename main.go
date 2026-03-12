package main

import (
	"fmt"
	"os"

	"blog/internal/config"
	"blog/internal/deploy"
	"blog/internal/generator"
	"blog/internal/server"
)

func main() {
	inputDir := config.InputDir()
	outputDir := config.OutputDir()

	// Ensure input dir exists
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(inputDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot create input dir %s: %v\n", inputDir, err)
			os.Exit(1)
		}
	}

	args := os.Args[1:]

	if len(args) == 0 {
		// Default: build + serve + watch
		if err := server.Serve(inputDir, outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	switch args[0] {
	case "build":
		if err := generator.Build(inputDir, outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "build error: %v\n", err)
			os.Exit(1)
		}
	case "deploy":
		if err := deploy.Deploy(inputDir, outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "deploy error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		fmt.Fprintln(os.Stderr, "usage: blog [build|deploy]")
		os.Exit(1)
	}
}
