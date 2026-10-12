#!/usr/bin/env python3
"""Parse a tree of markdown into one read-only library card.

Drops changelogs, license grants, and generated lists. Stores a heading
outline and a short claim per kept file. Does not copy the upstream file.
"""

import argparse
import hashlib
import json
import os
import re
from pathlib import Path

DROP_NAME = re.compile(
    r"(changelog|license|copying|code_of_conduct|contributing)",
    re.I,
)
DROP_DIR = (
    "node_modules",
    ".git",
    ".deer-flow",
    "dist",
    "coverage",
    "vendor",
)
HEADING = re.compile(r"^(#{1,3})\s+(.+?)\s*$")
DROP_SECTION = re.compile(
    r"^(changelog|table of contents|license|contributing|code of conduct)\b",
    re.I,
)


def keep_file(path: Path, root: Path) -> bool:
    rel = path.relative_to(root).as_posix()
    parts = rel.split("/")
    if any(part in DROP_DIR or part.startswith(".") for part in parts[:-1]):
        return False
    if DROP_NAME.search(path.name):
        return False
    return path.suffix.lower() == ".md"


def claim_of(text: str) -> tuple[list[str], str, str]:
    headings = []
    kept_lines = []
    skip = False
    skip_level = 0
    for line in text.splitlines():
        match = HEADING.match(line.strip())
        if match:
            level = len(match.group(1))
            title = match.group(2).strip()
            if DROP_SECTION.match(title):
                skip = True
                skip_level = level
                continue
            if skip and level <= skip_level:
                skip = False
            if skip:
                continue
            if len(headings) < 12:
                headings.append(title)
            continue
        if skip or not line.strip():
            continue
        if line.strip().startswith(("|", "```", "<!--")):
            continue
        kept_lines.append(line.strip())
    blob = " ".join(kept_lines)
    blob = re.sub(r"\s+", " ", blob).strip()
    if len(blob) > 180:
        blob = blob[:177] + "..."
    sha_heads = sum(1 for title in headings if title.lower().startswith("sha-256"))
    if sha_heads >= 3 or "deduplicated license" in " ".join(headings).lower():
        return headings, "", "dropped-notice"
    reason = "kept" if headings or blob else "empty-after-drop"
    return headings, blob, reason


def ingest(src: Path, dest: Path, subject: str, origin: str) -> None:
    dest.mkdir(parents=True, exist_ok=True)
    rows = []
    sections = []
    paths = []

    def _walk_error(_err: OSError) -> None:
        return None

    for dirpath, dirnames, filenames in os.walk(src, onerror=_walk_error):
        dirnames[:] = [
            name
            for name in dirnames
            if name not in DROP_DIR and not name.startswith(".")
        ]
        for name in filenames:
            if name.lower().endswith(".md"):
                paths.append(Path(dirpath) / name)
    paths.sort()
    for path in paths:
        try:
            if not path.is_file():
                continue
        except OSError:
            continue
        if not keep_file(path, src):
            rel = path.relative_to(src).as_posix()
            rows.append(
                {
                    "path": rel,
                    "sha256": "",
                    "bytes": path.stat().st_size if path.is_file() else 0,
                    "reason": "dropped-name",
                }
            )
            continue
        try:
            raw = path.read_bytes()
        except OSError:
            continue
        text = raw.decode("utf-8", "replace")
        digest = hashlib.sha256(raw).hexdigest()
        headings, claim, reason = claim_of(text)
        rel = path.relative_to(src).as_posix()
        rows.append(
            {
                "path": rel,
                "sha256": digest,
                "bytes": len(raw),
                "reason": reason,
            }
        )
        if reason != "kept":
            continue
        head = "; ".join(headings[:8]) if headings else "(no heading)"
        sections.append(f"### {rel}\n\n- Headings: {head}\n- Claim: {claim or '(structure only)'}\n")

    card = dest / "CARD.md"
    body = (
        f"# {subject}\n\n"
        "status: active\n\n"
        f"origin: {origin}\n\n"
        "memory: reconcile with `plan/` and with `research/library/cards/`. "
        "This card does not replace those files.\n\n"
        f"Kept files: {sum(1 for row in rows if row['reason'] == 'kept')}. "
        f"Dropped files: {sum(1 for row in rows if row['reason'] != 'kept')}.\n\n"
        "The lines below are a parse. The upstream tree remains the source.\n\n"
        + "\n".join(sections)
    )
    card.write_text(body)
    manifest = dest / "sources.jsonl"
    manifest.write_text("".join(json.dumps(row) + "\n" for row in rows))
    os.chmod(card, 0o444)
    os.chmod(manifest, 0o444)


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--src", type=Path, required=True)
    parser.add_argument("--dest", type=Path, required=True)
    parser.add_argument("--subject", required=True)
    parser.add_argument("--origin", required=True)
    args = parser.parse_args()
    ingest(args.src, args.dest, args.subject, args.origin)


if __name__ == "__main__":
    main()
