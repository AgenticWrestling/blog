# Deployment Checklist

Tasks required to get agenticwrestl.ing live.

## 1. Third-party accounts

- [x] **Google Analytics 4** — Configured with measurement ID `G-NWZQREB58C`.
- [x] **Giscus** — Configured. Discussions enabled on AgenticWrestling/blog.

## 2. Update placeholders in source

Both templates in `internal/generator/generator.go` contain placeholder strings that must be replaced:

- [x] Replace both occurrences of `YOUR_GA4_MEASUREMENT_ID` with your GA4 Measurement ID
- [x] Enable GitHub Discussions on the repo: **Settings → Features → Discussions**

## 3. Set up GitHub repository

- [ ] Create a new **private** GitHub repository (e.g. `agenticwrestling-blog`)
- [ ] Push the code: `git remote add origin <repo-url> && git push -u origin main`
- [ ] Add the following repository secrets under **Settings → Secrets and variables → Actions**:
  - `PORKBUN_FTP_HOST` — FTP hostname from PorkBun hosting dashboard
  - `PORKBUN_FTP_USER` — FTP username
  - `PORKBUN_FTP_PASS` — FTP password

## 4. Add posts to the repository

The GitHub Actions workflow expects posts at `./posts/` in the repo root.

- [ ] Copy (or symlink) your post directories from `~/Documents/blog/` into `posts/` in the repo
- [ ] Verify the directory structure: each post should be `posts/<slug>/index.md` with any assets alongside it
- [ ] Exclude the `drafts/` directory (it is already ignored by the generator, but don't commit it)

## 5. Configure DNS on PorkBun

- [ ] Log in to PorkBun and navigate to the domain `agenticwrestl.ing`
- [ ] Point DNS to the PorkBun static hosting nameservers / A records (follow PorkBun's static hosting setup guide)
- [ ] Enable HTTPS / SSL in the PorkBun hosting dashboard

## 6. Test deployment

- [ ] Push a commit to `main` and confirm the GitHub Actions workflow completes successfully
- [ ] Visit [https://agenticwrestl.ing](https://agenticwrestl.ing) and verify the site loads
- [ ] Confirm GA4 is receiving hits (Realtime report in GA4 dashboard)
- [ ] Confirm GraphComment widget appears on a post page
