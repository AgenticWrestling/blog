package post

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type Post struct {
	Slug      string
	Title     string
	CreatedAt time.Time
	Tags      []string
	HTMLBody  string
	Dir       string // absolute path to source directory
}

var snakeCaseRe = regexp.MustCompile(`^[a-z0-9]+(_[a-z0-9]+)*$`)

func IsSnakeCase(s string) bool {
	return snakeCaseRe.MatchString(s)
}

func Load(dir string) (*Post, error) {
	slug := filepath.Base(dir)
	if !IsSnakeCase(slug) {
		fmt.Fprintf(os.Stderr, "warning: post directory %q is not snake_case\n", slug)
	}

	mdPath := filepath.Join(dir, "index.md")
	raw, err := os.ReadFile(mdPath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", mdPath, err)
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			meta.Meta,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	ctx := parser.NewContext()
	if err := md.Convert(raw, &buf, parser.WithContext(ctx)); err != nil {
		return nil, fmt.Errorf("rendering markdown: %w", err)
	}

	metadata := meta.Get(ctx)

	p := &Post{
		Slug: slug,
		Dir:  dir,
	}

	if t, ok := metadata["title"]; ok {
		p.Title = fmt.Sprintf("%v", t)
	} else {
		// Use slug as title if not set
		p.Title = strings.ReplaceAll(strings.ReplaceAll(slug, "_", " "), "-", " ")
		p.Title = toTitleCase(p.Title)
	}

	if tags, ok := metadata["tags"]; ok {
		switch v := tags.(type) {
		case []interface{}:
			for _, tag := range v {
				p.Tags = append(p.Tags, fmt.Sprintf("%v", tag))
			}
		}
	}

	if ca, ok := metadata["created_at"]; ok {
		switch v := ca.(type) {
		case time.Time:
			p.CreatedAt = v
		case string:
			t, err := time.Parse("2006-01-02", v)
			if err == nil {
				p.CreatedAt = t
			}
		}
	}

	if p.CreatedAt.IsZero() {
		fi, err := os.Stat(mdPath)
		if err == nil {
			p.CreatedAt = fi.ModTime()
		} else {
			p.CreatedAt = time.Now()
		}
		if err := addCreatedAt(mdPath, raw, p.CreatedAt); err != nil {
			fmt.Fprintf(os.Stderr, "warning: could not update frontmatter in %s: %v\n", mdPath, err)
		}
	}

	// Validate file references
	validateRefs(dir)

	// Transform mermaid code blocks
	p.HTMLBody = transformMermaid(buf.String())

	return p, nil
}

func addCreatedAt(path string, raw []byte, t time.Time) error {
	content := string(raw)
	dateStr := t.Format("2006-01-02")

	if strings.HasPrefix(content, "---") {
		// Find the closing ---
		rest := content[3:]
		idx := strings.Index(rest, "\n---")
		if idx == -1 {
			return fmt.Errorf("malformed frontmatter")
		}
		frontmatter := rest[:idx]
		body := rest[idx+4:]

		newFrontmatter := strings.TrimRight(frontmatter, "\n") + "\ncreated_at: " + dateStr + "\n"
		newContent := "---\n" + newFrontmatter + "---" + body
		return os.WriteFile(path, []byte(newContent), 0644)
	}

	// No frontmatter — prepend it
	newContent := fmt.Sprintf("---\ncreated_at: %s\n---\n\n%s", dateStr, content)
	return os.WriteFile(path, []byte(newContent), 0644)
}

var refRe = regexp.MustCompile(`(?:!\[[^\]]*\]|]\s*)\(([^)#\s][^)]*)\)`)

func validateRefs(dir string) {
	// Read original markdown for link checking
	mdPath := filepath.Join(dir, "index.md")
	raw, err := os.ReadFile(mdPath)
	if err != nil {
		return
	}
	matches := refRe.FindAllSubmatch(raw, -1)
	for _, m := range matches {
		ref := string(m[1])
		if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") || strings.HasPrefix(ref, "//") {
			continue
		}
		refPath := filepath.Join(dir, ref)
		if _, err := os.Stat(refPath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "warning: missing referenced file %q in post %s\n", ref, filepath.Base(dir))
		}
	}
}

var mermaidRe = regexp.MustCompile("(?s)```mermaid\\s*\\n(.*?)\\n```")

func transformMermaid(html string) string {
	return mermaidRe.ReplaceAllStringFunc(html, func(match string) string {
		sub := mermaidRe.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		return `<div class="mermaid">` + sub[1] + `</div>`
	})
}

func toTitleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			runes := []rune(w)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

func LoadAll(inputDir string) ([]*Post, error) {
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return nil, fmt.Errorf("reading input dir %s: %w", inputDir, err)
	}

	var posts []*Post
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		mdPath := filepath.Join(inputDir, e.Name(), "index.md")
		if _, err := os.Stat(mdPath); os.IsNotExist(err) {
			continue
		}
		p, err := Load(filepath.Join(inputDir, e.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping post %s: %v\n", e.Name(), err)
			continue
		}
		posts = append(posts, p)
	}

	// Sort reverse chronological
	for i := 0; i < len(posts); i++ {
		for j := i + 1; j < len(posts); j++ {
			if posts[i].CreatedAt.Before(posts[j].CreatedAt) {
				posts[i], posts[j] = posts[j], posts[i]
			}
		}
	}

	return posts, nil
}

