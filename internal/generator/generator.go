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
  <title>AgenticWrestl.ing</title>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Chelsea+Market&family=Noto+Sans:wght@300&display=swap" rel="stylesheet">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css">
  <script async src="https://www.googletagmanager.com/gtag/js?id=YOUR_GA4_MEASUREMENT_ID"></script>
  <script>window.dataLayer=window.dataLayer||[];function gtag(){dataLayer.push(arguments);}gtag('js',new Date());gtag('config','YOUR_GA4_MEASUREMENT_ID');</script>
  <style>
    h1, h2, h3, h4 { font-family: 'Chelsea Market', cursive; }
    body { font-family: 'Noto Sans', sans-serif; font-weight: 300; font-size: 0.85rem; line-height: 1.8; }
    main { max-width: 860px; margin-inline: auto; }
    .tags { display: flex; gap: 0.5rem; flex-wrap: wrap; margin-top: 0.5rem; }
    .tag { background: var(--pico-primary-background); color: white; padding: 0.2rem 0.6rem; border-radius: 1rem; font-size: 0.8rem; cursor: pointer; user-select: none; transition: opacity 0.15s; }
    .tag:hover { opacity: 0.8; }
    .tag.active { outline: 2px solid white; outline-offset: -2px; }
    .post-card { margin-bottom: 1.5rem; padding: 1.5rem; border: 1px solid var(--pico-muted-border-color); border-radius: 8px; transition: box-shadow 0.2s; }
    .post-card:hover { box-shadow: 0 2px 12px rgba(0,0,0,0.08); }
    .post-meta { color: var(--pico-muted-color); font-size: 0.75rem; margin-top: 0.75rem; margin-bottom: 0; }
    .post-card h2 { margin-bottom: 0.5rem; }
    .post-card h2 a { text-decoration: none; }
    .post-summary { margin-top: 0.5rem; }
    .site-header { padding: 2rem 0 1rem; border-bottom: 1px solid var(--pico-muted-border-color); margin-bottom: 2rem; }
    .byline { color: var(--pico-muted-color); font-size: 0.85rem; margin-top: 0.4rem; }
    #filter-bar { display: none; align-items: center; gap: 0.5rem; flex-wrap: wrap; padding: 0.75rem 1rem; margin-bottom: 1.5rem; background: var(--pico-card-background-color); border: 1px solid var(--pico-muted-border-color); border-radius: 8px; }
    #filter-bar.visible { display: flex; }
    #filter-bar .filter-label { color: var(--pico-muted-color); font-size: 0.8rem; margin-right: 0.25rem; }
    .filter-tag { display: inline-flex; align-items: center; gap: 0.3rem; background: var(--pico-primary-background); color: white; padding: 0.2rem 0.4rem 0.2rem 0.6rem; border-radius: 1rem; font-size: 0.8rem; }
    .filter-tag button { all: unset; cursor: pointer; font-size: 1rem; line-height: 1; opacity: 0.8; }
    .filter-tag button:hover { opacity: 1; }
    #no-results { display: none; color: var(--pico-muted-color); padding: 2rem 0; }
  </style>
</head>
<body>
  <main class="container">
    <header class="site-header">
      <h1>AgenticWrestl.ing</h1>
      <p class="byline">Running into agentic walls and clambering over them.<br>
      <a href="mailto:nicholas@reid.contact">Nicholas Reid</a></p>
    </header>
    <div id="filter-bar">
      <span class="filter-label">Filtered by:</span>
    </div>
    {{range .Posts}}
    <article class="post-card" data-tags="{{range .Tags}}{{.}} {{end}}">
      <h2><a href="{{.Slug}}/index.html">{{.Title}}</a></h2>
      {{if .Summary}}<p class="post-summary">{{.Summary}}</p>{{end}}
      {{if .Tags}}
      <div class="tags">
        {{range .Tags}}<span class="tag" data-tag="{{.}}">{{.}}</span>{{end}}
      </div>
      {{end}}
      <p class="post-meta">Last updated: {{.CreatedAt.Format "January 2, 2006"}}</p>
    </article>
    {{else}}
    <p>No posts yet.</p>
    {{end}}
    <p id="no-results">No posts match the selected filters.</p>
  </main>
  <script>
    const active = new Set();

    function render() {
      const bar = document.getElementById('filter-bar');
      // Rebuild filter chips (keep the label, replace the rest)
      const label = bar.querySelector('.filter-label');
      bar.innerHTML = '';
      bar.appendChild(label);
      active.forEach(tag => {
        const chip = document.createElement('span');
        chip.className = 'filter-tag';
        chip.innerHTML = tag + ' <button aria-label="Remove filter">\u00d7</button>';
        chip.querySelector('button').addEventListener('click', () => { active.delete(tag); render(); });
        bar.appendChild(chip);
      });
      bar.classList.toggle('visible', active.size > 0);

      // Show/hide posts
      let visible = 0;
      document.querySelectorAll('.post-card').forEach(card => {
        const cardTags = card.dataset.tags.trim().split(/\s+/);
        const match = active.size === 0 || [...active].every(t => cardTags.includes(t));
        card.style.display = match ? '' : 'none';
        if (match) visible++;
      });
      document.getElementById('no-results').style.display = (active.size > 0 && visible === 0) ? '' : 'none';

      // Sync active class on all tag pills
      document.querySelectorAll('.tag').forEach(el => {
        el.classList.toggle('active', active.has(el.dataset.tag));
      });
    }

    document.querySelectorAll('.tag').forEach(el => {
      el.addEventListener('click', () => {
        active.has(el.dataset.tag) ? active.delete(el.dataset.tag) : active.add(el.dataset.tag);
        render();
      });
    });
  </script>
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
  <script async src="https://www.googletagmanager.com/gtag/js?id=YOUR_GA4_MEASUREMENT_ID"></script>
  <script>window.dataLayer=window.dataLayer||[];function gtag(){dataLayer.push(arguments);}gtag('js',new Date());gtag('config','YOUR_GA4_MEASUREMENT_ID');</script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"></script>
  <style>
    h1, h2, h3, h4 { font-family: 'Chelsea Market', cursive; }
    body { font-family: 'Noto Sans', sans-serif; font-weight: 300; font-size: 0.85rem; line-height: 1.8; }
    main { max-width: 1100px; margin-inline: auto; }
    .page-layout { display: flex; gap: 2.5rem; align-items: flex-start; }
    .post-content { flex: 1; min-width: 0; }
    .tags { display: flex; gap: 0.5rem; flex-wrap: wrap; margin-top: 0.5rem; }
    .tag { background: var(--pico-primary-background); color: white; padding: 0.2rem 0.6rem; border-radius: 1rem; font-size: 0.8rem; }
    .post-meta { color: var(--pico-muted-color); font-size: 0.9rem; margin-bottom: 1rem; }
    .back-link { display: inline-block; margin-bottom: 1.5rem; font-size: 0.9rem; }
    .post-header { margin-bottom: 2rem; padding-bottom: 1rem; border-bottom: 1px solid var(--pico-muted-border-color); }
    pre { background: var(--pico-code-background); border-radius: 6px; padding: 1rem; overflow-x: auto; }
    code { font-size: 0.9em; }
    .mermaid-diagram { text-align: center; margin: 1.5rem 0; }
    .mermaid-diagram svg { max-width: 100%; height: auto; }
    table img { min-width: 3rem; min-height: 3rem; }
    .toc-sidebar { width: 20%; flex-shrink: 0; position: sticky; top: 1.5rem; max-height: calc(100vh - 3rem); overflow-y: auto; overflow-x: hidden; }
    .toc-sidebar nav { border-left: 2px solid var(--pico-muted-border-color); padding-left: 1rem; }
    .toc-title { font-size: 0.75rem; font-weight: 600; text-transform: uppercase; letter-spacing: 0.05em; color: var(--pico-muted-color); margin-bottom: 0.5rem; }
    .toc-list { list-style: none; padding: 0; margin: 0; }
    .toc-list li { margin: 0; }
    .toc-list a { display: block; padding: 0.2rem 0; font-size: 0.8rem; color: var(--pico-muted-color); text-decoration: none; line-height: 1.4; transition: color 0.15s; }
    .toc-list a:hover, .toc-list a.active { color: var(--pico-primary); }
    .toc-list .toc-h3 a { padding-left: 0.75rem; font-size: 0.75rem; }
    @media (max-width: 768px) { .toc-sidebar { display: none; } main { max-width: 860px; } }
  </style>
</head>
<body>
  <main class="container">
    <a class="back-link" href="../index.html">← Back to index</a>
    <div class="page-layout">
      <article class="post-content">
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
      <aside class="toc-sidebar">
        <nav id="toc">
          <p class="toc-title">Contents</p>
          <ul class="toc-list" id="toc-list"></ul>
        </nav>
      </aside>
    </div>
  </main>
  <script>hljs.highlightAll();</script>
  <script>
    var __semio__params = {
      graphcommentId: "YOUR_GRAPHCOMMENT_ID",
      behaviour: { uid: "{{.Slug}}" }
    };
    function __semio__onload() { __semio__init_graphcomment(); }
    (function() {
      var gc = document.createElement('script'); gc.type = 'text/javascript'; gc.async = true;
      gc.onload = __semio__onload; gc.defer = true;
      gc.src = 'https://integration.graphcomment.com/gc_graphlogin.js?' + Date.now();
      (document.head || document.documentElement).appendChild(gc);
    })();
  </script>
  <script>
    (function() {
      const headings = document.querySelector('.post-content section').querySelectorAll('h2, h3');
      const list = document.getElementById('toc-list');
      if (headings.length < 2) { document.querySelector('.toc-sidebar').style.display = 'none'; return; }
      headings.forEach(h => {
        const li = document.createElement('li');
        li.className = h.tagName === 'H3' ? 'toc-h3' : 'toc-h2';
        const a = document.createElement('a');
        a.href = '#' + h.id;
        a.textContent = h.textContent;
        li.appendChild(a);
        list.appendChild(li);
      });

      // Scroll-spy
      const links = list.querySelectorAll('a');
      const observer = new IntersectionObserver(entries => {
        entries.forEach(e => {
          if (e.isIntersecting) {
            links.forEach(a => a.classList.remove('active'));
            const active = list.querySelector('a[href="#' + e.target.id + '"]');
            if (active) active.classList.add('active');
          }
        });
      }, { rootMargin: '0px 0px -70% 0px' });
      headings.forEach(h => observer.observe(h));
    })();
  </script>
</body>
</html>`

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

