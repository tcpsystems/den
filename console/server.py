#!/usr/bin/env python3
"""Local controller for den. Read-only. Binds to loopback."""

from __future__ import annotations

import json
import subprocess
import urllib.request
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
HOST = "127.0.0.1"
PORT = 8792
PROBES = {
    "mesh": "http://127.0.0.1:8787/api/health",
    "embed": "http://127.0.0.1:8091/v1/models",
    "memory": "http://127.0.0.1:1933/health",
    "omniroute": "https://ai.tcp.systems/api/health",
}


def git(*args: str) -> str:
    result = subprocess.run(
        ["git", "-C", str(ROOT), *args],
        check=False,
        capture_output=True,
        text=True,
        timeout=8,
    )
    return result.stdout.strip()


def branches() -> list[dict]:
    raw = git(
        "for-each-ref",
        "--format=%(refname:short)|%(objectname:short)|%(committerdate:relative)|%(subject)",
        "refs/heads",
    )
    rows = []
    for line in raw.splitlines():
        name, sha, when, subject = (line.split("|", 3) + ["", "", ""])[:4]
        ahead = git("rev-list", "--count", f"develop..{name}") or "0"
        kind = "archive" if name == "main" else "section" if name.startswith("section/") else "draft" if name.startswith("r") else "line"
        rows.append(
            {
                "name": name,
                "sha": sha,
                "when": when,
                "subject": subject,
                "ahead": int(ahead or 0),
                "kind": kind,
            }
        )
    rows.sort(key=lambda row: (row["kind"], row["name"]))
    return rows


def probe(url: str) -> str:
    try:
        with urllib.request.urlopen(url, timeout=5) as res:
            return "up" if res.status == 200 else f"http {res.status}"
    except Exception:
        return "down"


def guests() -> list[str]:
    try:
        result = subprocess.run(
            ["smolvm", "machine", "ls"],
            check=False,
            capture_output=True,
            text=True,
            timeout=5,
        )
    except FileNotFoundError:
        return []
    return [line for line in result.stdout.splitlines() if line.strip()]


def state() -> dict:
    rows = branches()
    return {
        "repo": "tcpsystems/den",
        "services": {name: probe(url) for name, url in PROBES.items()},
        "counts": {
            "branches": len(rows),
            "sections": sum(1 for row in rows if row["kind"] == "section"),
            "drafts": sum(1 for row in rows if row["kind"] == "draft"),
            "ahead": sum(1 for row in rows if row["ahead"] > 0),
        },
        "branches": rows,
        "guests": guests(),
        "schedules": [],
    }


class Handler(BaseHTTPRequestHandler):
    def log_message(self, fmt: str, *args) -> None:
        return

    def send_json(self, payload: dict) -> None:
        body = json.dumps(payload).encode()
        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(body)))
        self.end_headers()
        self.wfile.write(body)

    def do_GET(self) -> None:
        if self.path == "/api/state":
            self.send_json(state())
            return
        if self.path in ("/", "/index.html"):
            page = Path(__file__).with_name("index.html").read_bytes()
            self.send_response(200)
            self.send_header("Content-Type", "text/html; charset=utf-8")
            self.send_header("Content-Length", str(len(page)))
            self.end_headers()
            self.wfile.write(page)
            return
        self.send_response(404)
        self.end_headers()


def main() -> None:
    ThreadingHTTPServer((HOST, PORT), Handler).serve_forever()


if __name__ == "__main__":
    main()
