#!/usr/bin/env python3
"""Score harness-winner models locally and drop only the ones that never answer.

A model stays when one reply is complete. A failed, empty, or invalid reply
counts. After three of those, it gets two more tries. Five bad replies remove
it from the combo. Survivors are ordered by a local score.
"""

from __future__ import annotations

import json
import os
import time
import urllib.error
import urllib.request

BASE = "https://ai.tcp.systems"
COMBO_ID = "a6fd3248-799f-49ff-84b5-21f824fabb10"
KEY_PATH = os.path.expanduser("~/.config/omniroute/api.key")
OUT = os.path.expanduser("~/Projects/harness-plan/sections/consensus/winner-scores.json")
PROMPT = (
    "Reply with exactly two lines and nothing else.\n"
    "TYPE SectionPacket\n"
    "FUNC Validate"
)

WEIGHTS = {
    "quality": 0.18,
    "efficacy": 0.16,
    "efficiency": 0.12,
    "performance": 0.12,
    "cost": 0.10,
    "reasoning": 0.10,
    "thinking": 0.08,
    "maintainability": 0.14,
}


def key() -> str:
    return open(KEY_PATH, encoding="utf-8").read().strip()


def request(method: str, path: str, body: dict | None = None, timeout: int = 50) -> tuple[int, dict, dict]:
    data = None if body is None else json.dumps(body).encode()
    req = urllib.request.Request(
        BASE + path,
        data=data,
        method=method,
        headers={
            "Authorization": "Bearer " + key(),
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
    )
    try:
        with urllib.request.urlopen(req, timeout=timeout) as res:
            raw = res.read().decode()
            payload = json.loads(raw) if raw else {}
            return res.status, payload, dict(res.headers)
    except urllib.error.HTTPError as err:
        raw = err.read().decode(errors="replace")
        try:
            payload = json.loads(raw) if raw else {}
        except json.JSONDecodeError:
            payload = {"error": raw[:300]}
        return err.code, payload, dict(err.headers)


def text_of(payload: dict) -> str:
    choices = payload.get("choices") or []
    if not choices:
        return ""
    msg = choices[0].get("message") or {}
    content = msg.get("content")
    if isinstance(content, str):
        return content.strip()
    if isinstance(content, list):
        parts = []
        for block in content:
            if isinstance(block, dict) and isinstance(block.get("text"), str):
                parts.append(block["text"])
        return "\n".join(parts).strip()
    return ""


def valid(text: str) -> bool:
    lines = [line.strip() for line in text.splitlines() if line.strip()]
    if len(lines) != 2:
        return False
    return lines[0] == "TYPE SectionPacket" and lines[1] == "FUNC Validate"


def clamp(n: float) -> float:
    return max(0.0, min(100.0, n))


def score(text: str, elapsed_ms: int, usage: dict) -> dict:
    lines = [line.strip() for line in text.splitlines() if line.strip()]
    exact = valid(text)
    has_type = any(line == "TYPE SectionPacket" for line in lines)
    has_func = any(line == "FUNC Validate" for line in lines)
    quality = 100.0 if exact else (40.0 if has_type and has_func else 0.0)
    efficacy = 100.0 if exact else 0.0
    out_tokens = int((usage or {}).get("completion_tokens") or (usage or {}).get("output_tokens") or 0)
    in_tokens = int((usage or {}).get("prompt_tokens") or (usage or {}).get("input_tokens") or 0)
    reasoning = int(((usage or {}).get("completion_tokens_details") or {}).get("reasoning_tokens") or (usage or {}).get("reasoning_tokens") or 0)
    # A two-line answer should be small. Longer replies cost efficiency.
    efficiency = clamp(100 - max(0, out_tokens - 16) * 2)
    performance = clamp(100 - (elapsed_ms / 300))
    # No priced field is required. Fewer total tokens is the local cost proxy.
    cost = clamp(100 - (in_tokens + out_tokens) / 20)
    thinking = 70.0 if 0 < reasoning <= 200 else (40.0 if reasoning == 0 else clamp(100 - reasoning / 20))
    reasoning_score = 80.0 if exact and reasoning <= 200 else (50.0 if exact else 10.0)
    # The names match the harness packet. Exact shape is the maintainable reply.
    maintainability = 100.0 if exact else 20.0
    parts = {
        "quality": quality,
        "efficacy": efficacy,
        "efficiency": efficiency,
        "performance": performance,
        "cost": cost,
        "reasoning": reasoning_score,
        "thinking": thinking,
        "maintainability": maintainability,
    }
    total = round(sum(parts[k] * WEIGHTS[k] for k in WEIGHTS), 2)
    return {"total": total, "parts": parts, "elapsedMs": elapsed_ms, "outputTokens": out_tokens}


def probe(model_id: str) -> dict:
    bad = 0
    attempts = []
    kept = None
    while bad < 5:
        started = time.monotonic()
        status, payload, _headers = request(
            "POST",
            "/v1/chat/completions",
            {
                "model": model_id,
                "messages": [{"role": "user", "content": PROMPT}],
                "max_tokens": 80,
                "temperature": 0,
            },
        )
        elapsed = int((time.monotonic() - started) * 1000)
        text = text_of(payload)
        served = payload.get("model") or ""
        ok = status == 200 and valid(text)
        attempts.append(
            {
                "status": status,
                "ok": ok,
                "elapsedMs": elapsed,
                "servedModel": served,
                "text": text[:240],
            }
        )
        if ok:
            kept = score(text, elapsed, payload.get("usage") or {})
            break
        bad += 1
    return {"model": model_id, "bad": bad, "kept": kept is not None, "score": kept, "attempts": attempts}


def main() -> None:
    status, combo, _ = request("GET", f"/api/combos/{COMBO_ID}")
    if status != 200:
        raise SystemExit(f"combo read failed: {status}")
    models = []
    for item in combo.get("models") or []:
        model_id = item.get("model")
        if model_id:
            models.append(item)
    results = [probe(item["model"]) for item in models]
    survivors = []
    removed = []
    for item, result in zip(models, results):
        if result["kept"] and result["score"]:
            item = dict(item)
            item["weight"] = result["score"]["total"]
            survivors.append((item, result))
        else:
            removed.append(result["model"])
    survivors.sort(key=lambda pair: pair[1]["score"]["total"], reverse=True)
    ranked = []
    for index, (item, _result) in enumerate(survivors, start=1):
        item["priority"] = index
        ranked.append(item)
    config = {
        "explorationRate": 0.12,
        "routerStrategy": "lkgp",
        "auto": {
            "explorationRate": 0.12,
            "modePack": "quality-first",
            "routerStrategy": "lkgp",
            "weights": {
                "quota": 0.08,
                "health": 0.12,
                "costInv": 0.14,
                "latencyInv": 0.14,
                "taskFit": 0.28,
                "stability": 0.14,
                "tierPriority": 0.10,
            },
        },
    }
    put_status, put_body, _ = request(
        "PUT",
        f"/api/combos/{COMBO_ID}",
        {
            "strategy": "auto",
            "models": ranked,
            "config": config,
            "description": (
                "Winner stays, losers are retested. A model leaves only after "
                "three bad replies plus two more failed tries."
            ),
        },
    )
    report = {
        "comboId": COMBO_ID,
        "putStatus": put_status,
        "removed": removed,
        "ranking": [
            {
                "model": item["model"],
                "priority": item.get("priority"),
                "total": result["score"]["total"] if result["score"] else None,
                "parts": result["score"]["parts"] if result["score"] else None,
                "badBeforeKeep": result["bad"],
            }
            for item, result in survivors
        ],
        "attempts": results,
        "putError": None if put_status == 200 else put_body,
    }
    os.makedirs(os.path.dirname(OUT), exist_ok=True)
    with open(OUT, "w", encoding="utf-8") as fh:
        json.dump(report, fh, indent=2)
    print(f"put {put_status} kept {len(ranked)} removed {len(removed)}")
    for row in report["ranking"]:
        print(f"{row['priority']:2} {row['total']:6} {row['model']}")
    for model_id in removed:
        print(f"removed {model_id}")


if __name__ == "__main__":
    main()
