# Public proxy list sources

Probed 2026-10-12, live curl GETs, no account/auth used. Raw dumps, parsed
TSVs, and the sample-check live under `/tmp/harness-wave/proxies/` (not
committed — scratch only).

| source | URL | format | coverage | ToS | refresh | account-needed | verified-sample-ok/dead |
|---|---|---|---|---|---|---|---|
| proxyscrape.com | `https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=5000&country=all&ssl=all&anonymity=all` | plain text, `ip:port`/line, CRLF | 500 rows; no country field | free tier, "use responsibly", no published SLA | continuous, pull-on-demand | no | 0 ok / 2 dead |
| TheSpeedX/PROXY-List | `https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/http.txt` | plain text, `ip:port`/line | 2794 lines, capped 2000; no country/type metadata (http per filename) | open GitHub repo, MIT-style, no license file confirmed | cron commits, roughly every few hours (repo activity, not live-confirmed) | no | 0 ok / 2 dead |
| geonode.com (proxylist API) | `https://proxylist.geonode.com/api/proxy-list?limit=500&page=1&sort_by=lastChecked&sort_type=desc` | JSON: ip, port, protocols[], country (ISO2), anonymityLevel, lastChecked | total 3781 proxies; page 1 = 500 parsed, country populated | personal/eval use, no heavy-use guarantee | rolling `lastChecked`, continuous re-verification | no | 0 ok / 2 dead |
| free-proxy-list.net | `https://free-proxy-list.net/` | HTML table (IP, Port, Code, Country, Anonymity, Google, Https, Last Checked) | 298 rows via `<td>` regex; IP/port/country code captured | free to scrape for personal use, no formal API | page claims ~10 min (not independently confirmed) | no | 1 ok / 1 dead — ok = 217.69.121.5:5670, HTTP 407 (proxy auth required, TCP+handshake reached) |
| openproxylist.xyz | `https://openproxylist.xyz/http.txt` | plain text, `ip:port`/line | 6050 lines (1 junk `0.0.0.0:80` dropped), capped 2000; no metadata | none found, use at own risk | unknown/unstated | no | 0 ok / 2 dead |
| proxy-list.download (v2 API) | `https://www.proxy-list.download/api/v2/get?l=en&t=http` | JSON array of `{IP, PORT}` per docs | n/a — endpoint down | n/a | n/a | no | **unreachable: HTTP 502 on two tries (with/without `-L`)** |
| proxy-daily.com | `https://www.proxy-daily.com/` | HTML, but list is client-side JS/paid-API rendered | n/a — zero ip:port pairs in raw GET | n/a | n/a | no | **JS-gated: 301→200 landing page, no scrapable list via plain GET** |
| spys.one | `https://spys.one/en/http-proxy-list/` | HTML, ports obfuscated via inline JS (`spy14`-style functions) | 50 IP octet-groups found, ports not resolvable without JS | n/a | n/a | no | **JS-obfuscated: not parsed, needs headless-JS fetch** |

## Parsing strategy

- Plain-text sources (proxyscrape, TheSpeedX, openproxylist): regex-split each
  line on `ip:port`, type defaults to `http` (source has no type field).
- geonode: parse JSON `data[]`, join `protocols[]` with `+`, keep `country`.
- free-proxy-list.net: regex over `<td>...</td>` table cells (ip, port,
  2-letter country code).
- All five: dedupe on `(ip, port)`, cap 2000 rows/source, write to
  `/tmp/harness-wave/proxies/<source>.tsv` with header `ip, port, type, country`
  (country blank where the source has none).
- proxy-list.download, proxy-daily.com, spys.one excluded from parsing — down
  (502), JS-rendered, and JS-obfuscated respectively; none expose a plain-GET
  list at probe time.

## Validation

10 proxies sampled (2 per parsed source), `curl -x http://ip:port -m 5
--connect-timeout 5 http://example.com`. 1/10 reachable (free-proxy-list.net,
HTTP 407 counted ok as the TCP+proxy handshake succeeded); 9/10 dead
(refused/timeout). Consistent with the known high churn of free,
unauthenticated public proxy lists — treat any entry as ephemeral and
re-validate before use. Results: `/tmp/harness-wave/proxies/sample-check.tsv`.
