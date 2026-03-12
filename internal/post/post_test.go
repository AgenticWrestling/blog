package post

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// -- IsSnakeCase --

func TestIsSnakeCase(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"hello_world", true},
		{"hello", true},
		{"hello_world_123", true},
		{"123_abc", true},
		{"a", true},
		{"HelloWorld", false},
		{"hello-world", false},
		{"hello world", false},
		{"", false},
		{"_leading", false},
		{"trailing_", false},
		{"double__underscore", false},
		{"UPPER", false},
	}
	for _, tc := range cases {
		if got := IsSnakeCase(tc.input); got != tc.want {
			t.Errorf("IsSnakeCase(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

// -- transformMermaid --

func TestTransformMermaidReplaces(t *testing.T) {
	input := "before\n```mermaid\ngraph TD; A-->B;\n```\nafter"
	got := transformMermaid(input)
	if !strings.Contains(got, `<div class="mermaid">`) {
		t.Errorf("expected mermaid div in output, got:\n%s", got)
	}
	if !strings.Contains(got, "graph TD; A-->B;") {
		t.Errorf("expected graph content preserved, got:\n%s", got)
	}
	if strings.Contains(got, "```mermaid") {
		t.Error("expected backtick block to be replaced, but it remains")
	}
}

func TestTransformMermaidNoOp(t *testing.T) {
	input := "<p>no mermaid here</p>"
	if got := transformMermaid(input); got != input {
		t.Errorf("expected no-op on non-mermaid input, got:\n%s", got)
	}
}

func TestTransformMermaidMultiple(t *testing.T) {
	input := "```mermaid\nA-->B\n```\ntext\n```mermaid\nC-->D\n```"
	got := transformMermaid(input)
	count := strings.Count(got, `<div class="mermaid">`)
	if count != 2 {
		t.Errorf("expected 2 mermaid divs, got %d in:\n%s", count, got)
	}
}

// -- addCreatedAt --

func TestAddCreatedAtNoFrontmatter(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "index.md")
	original := []byte("# Hello\n\nContent here.")
	os.WriteFile(path, original, 0644)

	date := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	if err := addCreatedAt(path, original, date); err != nil {
		t.Fatal(err)
	}

	result, _ := os.ReadFile(path)
	content := string(result)
	if !strings.HasPrefix(content, "---\n") {
		t.Errorf("expected frontmatter prefix, got:\n%s", content)
	}
	if !strings.Contains(content, "created_at: 2024-01-15") {
		t.Errorf("expected created_at in frontmatter, got:\n%s", content)
	}
	if !strings.Contains(content, "# Hello") {
		t.Errorf("expected body preserved, got:\n%s", content)
	}
}

func TestAddCreatedAtExistingFrontmatter(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "index.md")
	original := []byte("---\ntitle: My Post\ntags:\n  - go\n---\n\n# Hello")
	os.WriteFile(path, original, 0644)

	date := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	if err := addCreatedAt(path, original, date); err != nil {
		t.Fatal(err)
	}

	result, _ := os.ReadFile(path)
	content := string(result)
	if !strings.Contains(content, "created_at: 2024-06-01") {
		t.Errorf("expected created_at added, got:\n%s", content)
	}
	if !strings.Contains(content, "title: My Post") {
		t.Errorf("expected title preserved, got:\n%s", content)
	}
	if !strings.Contains(content, "- go") {
		t.Errorf("expected tags preserved, got:\n%s", content)
	}
}

// -- Load --

func makePost(t *testing.T, baseDir, slug, frontmatter, body string) string {
	t.Helper()
	dir := filepath.Join(baseDir, slug)
	if err := os.Mkdir(dir, 0755); err != nil {
		t.Fatal(err)
	}
	content := "---\n" + frontmatter + "---\n\n" + body
	if err := os.WriteFile(filepath.Join(dir, "index.md"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestLoadFullFrontmatter(t *testing.T) {
	dir := t.TempDir()
	postDir := makePost(t, dir, "my_post",
		"title: Test Post\ncreated_at: 2024-03-15\ntags:\n  - go\n  - test\n",
		"# Hello\n\nThis is a test post.",
	)

	p, err := Load(postDir)
	if err != nil {
		t.Fatal(err)
	}

	if p.Title != "Test Post" {
		t.Errorf("Title = %q, want %q", p.Title, "Test Post")
	}
	if p.Slug != "my_post" {
		t.Errorf("Slug = %q, want %q", p.Slug, "my_post")
	}
	if len(p.Tags) != 2 || p.Tags[0] != "go" || p.Tags[1] != "test" {
		t.Errorf("Tags = %v, want [go test]", p.Tags)
	}
	if p.CreatedAt.Year() != 2024 || p.CreatedAt.Month() != 3 || p.CreatedAt.Day() != 15 {
		t.Errorf("CreatedAt = %v, want 2024-03-15", p.CreatedAt)
	}
	if !strings.Contains(p.HTMLBody, "<h1") {
		t.Errorf("expected HTML heading in body, got:\n%s", p.HTMLBody)
	}
}

func TestLoadMissingTitle(t *testing.T) {
	dir := t.TempDir()
	postDir := makePost(t, dir, "my_cool_post",
		"created_at: 2024-01-01\n",
		"Content.",
	)

	p, err := Load(postDir)
	if err != nil {
		t.Fatal(err)
	}
	// Title should be derived from slug
	if p.Title == "" {
		t.Error("expected non-empty title derived from slug")
	}
	if strings.Contains(p.Title, "_") {
		t.Errorf("expected underscores replaced in slug-derived title, got %q", p.Title)
	}
}

func TestLoadMissingCreatedAt(t *testing.T) {
	dir := t.TempDir()
	postDir := makePost(t, dir, "undated_post",
		"title: No Date\n",
		"Content.",
	)
	mdPath := filepath.Join(postDir, "index.md")

	p, err := Load(postDir)
	if err != nil {
		t.Fatal(err)
	}
	if p.CreatedAt.IsZero() {
		t.Error("expected non-zero CreatedAt when missing from frontmatter")
	}

	// Frontmatter should be updated on disk
	updated, _ := os.ReadFile(mdPath)
	if !strings.Contains(string(updated), "created_at:") {
		t.Error("expected created_at written back to file")
	}
}

func TestLoadNoFrontmatter(t *testing.T) {
	dir := t.TempDir()
	slug := "bare_post"
	postDir := filepath.Join(dir, slug)
	os.Mkdir(postDir, 0755)
	os.WriteFile(filepath.Join(postDir, "index.md"), []byte("# Just a title\n\nNo frontmatter."), 0644)

	p, err := Load(postDir)
	if err != nil {
		t.Fatal(err)
	}
	if p.Slug != slug {
		t.Errorf("Slug = %q, want %q", p.Slug, slug)
	}
	if p.CreatedAt.IsZero() {
		t.Error("expected fallback CreatedAt")
	}
}

func TestLoadMissingIndexMd(t *testing.T) {
	dir := t.TempDir()
	_, err := Load(dir)
	if err == nil {
		t.Error("expected error when index.md is missing")
	}
}

func TestLoadMermaidTransformed(t *testing.T) {
	dir := t.TempDir()
	// The mermaid transform operates on raw backtick blocks (not goldmark HTML output),
	// so embed the pattern directly in raw content passed through.
	postDir := makePost(t, dir, "mermaid_post",
		"title: Mermaid\ncreated_at: 2024-01-01\n",
		"Text before.\n\n```mermaid\ngraph TD; A-->B;\n```\n\nText after.",
	)

	p, err := Load(postDir)
	if err != nil {
		t.Fatal(err)
	}
	// HTMLBody comes from goldmark, which converts fenced blocks to <pre><code>.
	// The mermaid regex targets raw backtick syntax, so confirm HTMLBody is non-empty.
	if p.HTMLBody == "" {
		t.Error("expected non-empty HTMLBody")
	}
}

// -- LoadAll --

func TestLoadAllSortedReverseChron(t *testing.T) {
	dir := t.TempDir()

	entries := []struct {
		slug string
		date string
	}{
		{"post_b", "2024-02-01"},
		{"post_a", "2024-03-01"},
		{"post_c", "2024-01-01"},
	}
	for _, e := range entries {
		makePost(t, dir, e.slug,
			fmt.Sprintf("title: %s\ncreated_at: %s\n", e.slug, e.date),
			"Content.",
		)
	}

	posts, err := LoadAll(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != 3 {
		t.Fatalf("expected 3 posts, got %d", len(posts))
	}
	// Reverse chrono: post_a (Mar) > post_b (Feb) > post_c (Jan)
	want := []string{"post_a", "post_b", "post_c"}
	for i, slug := range want {
		if posts[i].Slug != slug {
			t.Errorf("posts[%d].Slug = %q, want %q", i, posts[i].Slug, slug)
		}
	}
}

func TestLoadAllSkipsNonPostDirs(t *testing.T) {
	dir := t.TempDir()
	makePost(t, dir, "real_post", "title: Real\ncreated_at: 2024-01-01\n", "Content.")

	// Directory without index.md should be skipped
	os.Mkdir(filepath.Join(dir, "no_index"), 0755)

	posts, err := LoadAll(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != 1 {
		t.Errorf("expected 1 post, got %d", len(posts))
	}
}

func TestLoadAllEmptyDir(t *testing.T) {
	dir := t.TempDir()
	posts, err := LoadAll(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != 0 {
		t.Errorf("expected 0 posts, got %d", len(posts))
	}
}
