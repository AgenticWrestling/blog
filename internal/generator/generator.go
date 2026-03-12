package generator

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"blog/internal/post"
)

const indexTmplSrc = `<!DOCTYPE html>
<html lang="en" data-theme="light">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Blog</title>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Chelsea+Market&family=Noto+Sans:wght@300&display=swap" rel="stylesheet">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css">
  <style>
    h1, h2, h3, h4 { font-family: 'Chelsea Market', cursive; }
    body { font-family: 'Noto Sans', sans-serif; font-weight: 300; }
    .tags { display: flex; gap: 0.5rem; flex-wrap: wrap; margin-top: 0.5rem; }
    .tag { background: var(--pico-primary-background); color: var(--pico-primary); padding: 0.2rem 0.6rem; border-radius: 1rem; font-size: 0.8rem; }
    .post-card { margin-bottom: 1.5rem; padding: 1.5rem; border: 1px solid var(--pico-muted-border-color); border-radius: 8px; transition: box-shadow 0.2s; }
    .post-card:hover { box-shadow: 0 2px 12px rgba(0,0,0,0.08); }
    .post-meta { color: var(--pico-muted-color); font-size: 0.9rem; margin-bottom: 0.5rem; }
    .post-card h2 { margin-bottom: 0.25rem; }
    .post-card h2 a { text-decoration: none; }
    .site-header { padding: 2rem 0 1rem; border-bottom: 1px solid var(--pico-muted-border-color); margin-bottom: 2rem; }
  </style>
</head>
<body>
  <main class="container">
    <header class="site-header">
      <h1>Blog</h1>
    </header>
    {{range .Posts}}
    <article class="post-card">
      <p class="post-meta">{{.CreatedAt.Format "January 2, 2006"}}</p>
      <h2><a href="{{.Slug}}/index.html">{{.Title}}</a></h2>
      {{if .Tags}}
      <div class="tags">
        {{range .Tags}}<span class="tag">{{.}}</span>{{end}}
      </div>
      {{end}}
    </article>
    {{else}}
    <p>No posts yet.</p>
    {{end}}
  </main>
</body>
</html>`

const postTmplSrc = `<!DOCTYPE html>
<html lang="en" data-theme="light">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.Title}}</title>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Chelsea+Market&family=Noto+Sans:wght@300&display=swap" rel="stylesheet">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/github.min.css">
  <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"></script>
  <style>
    h1, h2, h3, h4 { font-family: 'Chelsea Market', cursive; }
    body { font-family: 'Noto Sans', sans-serif; font-weight: 300; }
    .tags { display: flex; gap: 0.5rem; flex-wrap: wrap; margin-top: 0.5rem; }
    .tag { background: var(--pico-primary-background); color: var(--pico-primary); padding: 0.2rem 0.6rem; border-radius: 1rem; font-size: 0.8rem; }
    .post-meta { color: var(--pico-muted-color); font-size: 0.9rem; margin-bottom: 1rem; }
    .back-link { display: inline-block; margin-bottom: 1.5rem; font-size: 0.9rem; }
    .post-header { margin-bottom: 2rem; padding-bottom: 1rem; border-bottom: 1px solid var(--pico-muted-border-color); }
    pre { background: var(--pico-code-background); border-radius: 6px; padding: 1rem; overflow-x: auto; }
    code { font-size: 0.9em; }
    .mermaid { text-align: center; margin: 1.5rem 0; }
  </style>
</head>
<body>
  <main class="container">
    <a class="back-link" href="../index.html">← Back to index</a>
    <article>
      <header class="post-header">
        <h1>{{.Title}}</h1>
        <p class="post-meta">{{.CreatedAt.Format "January 2, 2006"}}</p>
        {{if .Tags}}
        <div class="tags">
          {{range .Tags}}<span class="tag">{{.}}</span>{{end}}
        </div>
        {{end}}
      </header>
      <section>
        {{.HTMLBody}}
      </section>
    </article>
  </main>
  {{if .HasMermaid}}
  <script src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js"></script>
  <script>mermaid.initialize({startOnLoad:true});</script>
  {{end}}
  <script>hljs.highlightAll();</script>
</body>
</html>`

type indexData struct {
	Posts []*post.Post
}

type postData struct {
	*post.Post
	HTMLBody   template.HTML
	HasMermaid bool
}

func Build(inputDir, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	posts, err := post.LoadAll(inputDir)
	if err != nil {
		return fmt.Errorf("loading posts: %w", err)
	}

	indexTmpl, err := template.New("index").Parse(indexTmplSrc)
	if err != nil {
		return fmt.Errorf("parsing index template: %w", err)
	}

	postTmpl, err := template.New("post").Parse(postTmplSrc)
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

		hasMermaid := strings.Contains(p.HTMLBody, `class="mermaid"`)

		postFile, err := os.Create(filepath.Join(postOutDir, "index.html"))
		if err != nil {
			return fmt.Errorf("creating post html for %s: %w", p.Slug, err)
		}

		data := postData{
			Post:       p,
			HTMLBody:   template.HTML(p.HTMLBody),
			HasMermaid: hasMermaid,
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

