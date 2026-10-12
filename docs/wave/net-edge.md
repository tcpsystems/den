# Net edge: skill pack vs named human action

## Question

Which network knowledge belongs in a den skill pack (DNS, HTTP, SMTP, anycast, eBPF/XDP as defense names, OCI image layers) and what must stay a named human action? High level only.

## Sources

- RFC 1035 (STD 13), *Domain names - implementation and specification*, P. Mockapetris, November 1987. Status: the domain system is official protocol plus experimental types; official parts include standard queries, responses, and Internet-class resource-record formats. Updated by later RFCs, including DNSSEC (4033–4035) and dynamic update (2136). https://www.rfc-editor.org/info/rfc1035/
- eBPF.io, *What is eBPF?* eBPF runs sandboxed programs in a privileged context such as the kernel, without changing kernel source or loading a kernel module. Programs attach at hooks, including XDP; a verifier rejects unsafe programs before JIT. https://ebpf.io/what-is-ebpf/

Local map already splits the same line: niche `net` is design (BGP, routing policy, anycast, interface and DNS); niche `ebpf` is receive-path filters on hosts we operate, attached to `sections/security`.

## Finding

A skill pack may name the objects and the decision, not operate the edge.

- **DNS.** Names, the resolver-versus-authoritative split, record types as labels (A, AAAA, CNAME, MX, TXT, NS), and the fact that zones and delegations are data a human publishes. RFC 1035 is the vocabulary. Zone edits, registrar changes, and DNSSEC key custody stay human.
- **HTTP.** Method, status class, host, and path as request shape. Timeouts and size limits as policy names. No client that speaks to a host the packet did not name.
- **SMTP.** Envelope versus header, MX as a DNS fact, and that mail submission is a separate service. No sender, no open relay, no message body generation.
- **Anycast.** One address, many sites, routing picks the site. The skill may say "this name is anycast" and point at the niche. Announcing a prefix is a human routing action.
- **eBPF/XDP.** Defense names only: hook, verifier, map, drop or pass on a host we operate. The pack cites the name and the owned-host rule. It does not carry a program, a bytecode blob, or a load command.
- **OCI image layers.** A layer is content-addressed, ordered, and immutable; the config and the manifest name the stack. The pack may say which digest is already in the store. Pulling, pushing, or retagging a registry is a human or a named tool with a host we own.

The pack is a glossary plus a refuse list. Each row points at one niche or one section. It does not become a control plane.

## What den should not copy

- Protocol wire layouts, query construction, or update opcodes from RFC 1035 and its updates.
- Loader steps, helper calls, map layouts, or XDP return codes from ebpf.io.
- BGP announcements, firewall rule text, scrub recipes, or rate numbers.
- Registry credentials, layer extraction, or any command that changes DNS, mail, or routing.

Those stay a named human action: a person, on a host we already operate, recorded as the premise action. The seat may only name the control.

## One plan line

Add one skill row, `skills/net-edge`, that lists DNS, HTTP, SMTP, anycast, eBPF/XDP, and OCI layers as names and points each at an existing niche, with a single refuse line: no publish, no announce, no load, no send.

## Unproven

- Whether that row stays a glossary or later grows a check that the packet names only an owned host.
- How DNSSEC and dynamic update (RFC 4033, RFC 2136) should be labeled if the glossary mentions them at all.
- Whether OCI layer names belong next to `virt` instead of `net`.
