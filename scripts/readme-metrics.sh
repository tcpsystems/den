#!/bin/sh
# Regenerate the measured block in README.md from files and live checks.
set -eu
cd "$(dirname "$0")/.."
commit=$(git rev-parse --short HEAD)
skills=$(find skills -mindepth 2 -maxdepth 2 -name SKILL.md | wc -l)
go_files=$(find . -name '*.go' -not -path './.git/*' | wc -l)
plans=$(find plan -name '*.md' | wc -l)
niches=$(grep -c '^| ' plan/niches/INDEX.md || true)
commits=$(git rev-list --count HEAD)
memory_files=$(find "${HOME}/.local/share/harness/memory" -type f 2>/dev/null | wc -l)
embed=$(curl -fsS --max-time 3 http://127.0.0.1:8091/health >/dev/null && echo healthy || echo down)
memory=$(curl -fsS --max-time 3 http://127.0.0.1:1933/health >/dev/null && echo healthy || echo down)
block=$(cat <<EOF
<!-- metrics:start -->
| Metric | Value |
| --- | ---: |
| Commit | \`$commit\` |
| Skills | $skills |
| Go files | $go_files |
| Plan documents | $plans |
| Niche rows | $niches |
| Commits | $commits |
| Embedder | $embed |
| Long memory | $memory |
| Local memory files | $memory_files |
<!-- metrics:end -->
EOF
)
python3 - "$block" <<'PY'
import pathlib, sys
block = sys.argv[1]
path = pathlib.Path("README.md")
text = path.read_text()
start = text.index("<!-- metrics:start -->")
end = text.index("<!-- metrics:end -->") + len("<!-- metrics:end -->")
path.write_text(text[:start] + block.strip() + text[end:])
PY
