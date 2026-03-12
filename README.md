# blog

A static site generator and local dev server for a personal blog. Converts Markdown posts (with YAML front matter) to HTML, with live-reload serving and S3 deploy.

## Usage

```
blog              # build, serve, and watch for changes
blog build        # build static site to output dir
blog deploy       # build then deploy to S3
```

## Configuration

Directories are configured via environment variables (see `internal/config`):

| Variable | Default |
|---|---|
| `BLOG_INPUT_DIR` | `~/posts` |
| `BLOG_OUTPUT_DIR` | `~/public` |

## Development

```sh
task build        # compile binary to ./blog
task install      # install binary to ~/bin
task test         # run tests
task check        # lint + test
task clean        # remove build artefacts
```

## Dependencies

- [goldmark](https://github.com/yuin/goldmark) — Markdown parser
- [goldmark-meta](https://github.com/yuin/goldmark-meta) — YAML front matter
- [fsnotify](https://github.com/fsnotify/fsnotify) — file watching for live reload
