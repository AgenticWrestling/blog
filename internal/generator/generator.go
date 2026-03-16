package generator

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"blog/internal/post"
)

//go:embed templates/index.go.tmpl templates/post.go.tmpl
var tmplFS embed.FS

type indexData struct {
	Posts []*post.Post
}

type postData struct {
	*post.Post
	HTMLBody template.HTML
}

func Build(inputDir, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	posts, err := post.LoadAll(inputDir)
	if err != nil {
		return fmt.Errorf("loading posts: %w", err)
	}

	indexSrc, err := tmplFS.ReadFile("templates/index.go.tmpl")
	if err != nil {
		return fmt.Errorf("reading index template: %w", err)
	}
	indexTmpl, err := template.New("index").Parse(string(indexSrc))
	if err != nil {
		return fmt.Errorf("parsing index template: %w", err)
	}

	postSrc, err := tmplFS.ReadFile("templates/post.go.tmpl")
	if err != nil {
		return fmt.Errorf("reading post template: %w", err)
	}
	postTmpl, err := template.New("post").Parse(string(postSrc))
	if err != nil {
		return fmt.Errorf("parsing post template: %w", err)
	}

	// Generate index
	idxFile, err := os.Create(filepath.Join(outputDir, "index.html"))
	if err != nil {
		return fmt.Errorf("creating index.html: %w", err)
	}
	if err := indexTmpl.Execute(idxFile, indexData{Posts: posts}); err != nil {
		idxFile.Close()
		return fmt.Errorf("rendering index: %w", err)
	}
	idxFile.Close()

	// Generate each post
	for _, p := range posts {
		postOutDir := filepath.Join(outputDir, p.Slug)
		if err := os.MkdirAll(postOutDir, 0755); err != nil {
			return fmt.Errorf("creating post dir %s: %w", p.Slug, err)
		}

		// Copy non-md files
		if err := copyAssets(p.Dir, postOutDir); err != nil {
			fmt.Fprintf(os.Stderr, "warning: copying assets for %s: %v\n", p.Slug, err)
		}

		postFile, err := os.Create(filepath.Join(postOutDir, "index.html"))
		if err != nil {
			return fmt.Errorf("creating post html for %s: %w", p.Slug, err)
		}

		data := postData{
			Post:     p,
			HTMLBody: template.HTML(p.HTMLBody),
		}
		if err := postTmpl.Execute(postFile, data); err != nil {
			postFile.Close()
			return fmt.Errorf("rendering post %s: %w", p.Slug, err)
		}
		postFile.Close()
	}

	fmt.Printf("Built %d posts -> %s\n", len(posts), outputDir)
	return nil
}

func copyAssets(srcDir, dstDir string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToLower(e.Name()), ".md") {
			continue
		}
		src := filepath.Join(srcDir, e.Name())
		dst := filepath.Join(dstDir, e.Name())
		if err := copyFile(src, dst); err != nil {
			fmt.Fprintf(os.Stderr, "warning: copying %s: %v\n", e.Name(), err)
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
