# Stack

```
dan.tcp.systems (later, Cloudflare Workers, design language)
        |
        v
den API binary (cmd/denapi)  OpenAPI, loopback now, auth when it is public
        |
        v
den daemon socket (cmd/denctl)  checks, queues, tooling
        ^
        |
den TUI (cmd/dentui) dials the socket when it is local
        |
        v
harness  skills, rules, impl flow, eval flow
        |
        +-- tcp.systems   expand the host and the price desk
        +-- den           this executor
        +-- fox, foxd, floki   company products, separate repos
```

The harness holds the rules and the skills. den consumes them. den executes, deploys, injects, validates, optimises, expands, and analyses. A harness change is a branch on `tcpsystems/den`. A project change is a branch in that project's repo. They are not the same commit.

The daemon is the only process that reads git and probes services. The TUI dials `~/.local/share/harness/den.sock` directly. The web API is a second binary on the same machine. It dials that socket and publishes HTTP. Browsers, admin dashboards, and later Zabbix, Grafana, or OpenTelemetry collectors use the published API. They do not open the socket.

`dan.tcp.systems` is a later Workers frontend. It calls the published API. It does not call the socket. Auth sits on that API before it leaves the machine. It is not on the public internet in this starting point.

Multi-agent work starts from `plan/BASE.md`. One orchestrator, one skill per seat, evidence then a labeled user-context block. Innovation is a section branch with a check, not a new framework.
