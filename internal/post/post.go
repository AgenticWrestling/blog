package post

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
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
	Summary   string
	CreatedAt time.Time
	Tags      []string
	HTMLBody  string
	Dir       string // absolute path to source directory
}

var slugRe = regexp.MustCompile(`^[a-z0-9]+([_-][a-z0-9]+)*$`)

func IsSnakeCase(s string) bool {
	return slugRe.MatchString(s)
}

// findMD returns the path to the first .md file in dir, or an error if none exists.
func findMD(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(strings.ToLower(e.Name()), ".md") {
			return filepath.Join(dir, e.Name()), nil
		}
	}
	return "", fmt.Errorf("no .md file found in %s", dir)
}

func Load(dir string) (*Post, error) {
	slug := filepath.Base(dir)
	if !IsSnakeCase(slug) {
		fmt.Fprintf(os.Stderr, "warning: post directory %q is not snake_case\n", slug)
	}

	mdPath, err := findMD(dir)
	if err != nil {
		return nil, err
	}
	raw, err := os.ReadFile(mdPath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", mdPath, err)
	}

	raw = processMermaid(dir, raw)

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

	if s, ok := metadata["summary"]; ok {
		p.Summary = fmt.Sprintf("%v", s)
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

	p.HTMLBody = buf.String()

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

var mermaidFenceRe = regexp.MustCompile("(?s)```mermaid[ \t]*\n(.*?)\n```")

// processMermaid finds fenced mermaid blocks in raw markdown, generates an SVG
// for each using mmdc, and replaces the block with the inlined SVG so that the
// page's font (Noto Sans) applies directly without needing an external request.
func processMermaid(dir string, raw []byte) []byte {
	n := 0
	return mermaidFenceRe.ReplaceAllFunc(raw, func(match []byte) []byte {
		sub := mermaidFenceRe.FindSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		n++
		svgPath := filepath.Join(dir, fmt.Sprintf("diagram-%d.svg", n))
		if err := renderMermaid(sub[1], svgPath); err != nil {
			fmt.Fprintf(os.Stderr, "warning: mermaid render failed for %s diagram %d: %v\n", filepath.Base(dir), n, err)
			return match // leave original fenced block on error
		}
		svgContent, err := os.ReadFile(svgPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: reading mermaid svg %s: %v\n", svgPath, err)
			return match
		}
		// Wrap in a div; two blank lines make goldmark treat it as an HTML block.
		return append([]byte("\n\n<div class=\"mermaid-diagram\">\n"), append(svgContent, []byte("\n</div>\n\n")...)...)
	})
}

const mermaidCSS = `text, tspan { font-family: 'Noto Sans', sans-serif; font-weight: 300; }`

func renderMermaid(diagram []byte, outPath string) error {
	tmp, err := os.CreateTemp("", "mermaid-*.mmd")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.Write(diagram); err != nil {
		tmp.Close()
		return err
	}
	tmp.Close()

	cssTmp, err := os.CreateTemp("", "mermaid-*.css")
	if err != nil {
		return err
	}
	defer os.Remove(cssTmp.Name())
	if _, err := cssTmp.WriteString(mermaidCSS); err != nil {
		cssTmp.Close()
		return err
	}
	cssTmp.Close()

	out, err := exec.Command("mmdc", "-i", tmp.Name(), "-o", outPath, "--cssFile", cssTmp.Name(), "--quiet").CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w\n%s", err, bytes.TrimSpace(out))
	}
	return nil
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
		if e.Name() == "drafts" {
			continue
		}
		postDir := filepath.Join(inputDir, e.Name())
		if _, err := findMD(postDir); err != nil {
			continue
		}
		p, err := Load(postDir)
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

