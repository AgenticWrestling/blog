package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"blog/internal/generator"

	"github.com/fsnotify/fsnotify"
)


func Serve(inputDir, outputDir string) error {
	// Initial build
	if err := generator.Build(inputDir, outputDir); err != nil {
		return fmt.Errorf("initial build: %w", err)
	}

	// Start file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("creating watcher: %w", err)
	}
	defer watcher.Close()

	// Watch inputDir recursively
	if err := watchRecursive(watcher, inputDir); err != nil {
		return fmt.Errorf("setting up watcher: %w", err)
	}

	// Debounced rebuild goroutine
	go func() {
		var timer *time.Timer
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Printf("change detected: %s\n", event.Name)
				if timer != nil {
					timer.Stop()
				}
				timer = time.AfterFunc(100*time.Millisecond, func() {
					fmt.Println("rebuilding...")
					if err := generator.Build(inputDir, outputDir); err != nil {
						fmt.Fprintf(os.Stderr, "rebuild error: %v\n", err)
					} else {
						fmt.Println("rebuild complete")
					}
					// Re-watch any new directories
					watchRecursive(watcher, inputDir)
				})
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Fprintf(os.Stderr, "watcher error: %v\n", err)
			}
		}
	}()

	// Serve static files
	fs := http.FileServer(http.Dir(outputDir))
	http.Handle("/", fs)

	fmt.Printf("Blog preview at http://localhost:9001 — watching ~/Documents/blog/\n")

	return http.ListenAndServe(":9001", nil)
}

func watchRecursive(watcher *fsnotify.Watcher, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
}
