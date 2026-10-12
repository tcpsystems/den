# Security, crypto

## Question

What security practices should den adopt for AppArmor, WireGuard, and cryptographic hygiene, as design rules a seat can follow? High level only: names, what they protect, what den should refuse. No procedures for attack. The path jail is `sections/security/match.py`. A secret never enters the vault or the packet.

## Sources

- AppArmor overview: https://apparmor.net/
- WireGuard overview: https://www.wireguard.com/
- Primitive names only: https://www.wireguard.com/protocol/
- Path jail: `sections/security/match.py` (`path_jail`, `Refused`)

## Finding

AppArmor is a Linux Security Module. A profile is mandatory access control for one program: which files it may touch and which capabilities it may use. Enforce mode makes that profile the limit. Complain mode only logs. The profile protects the host from a seat that can still run. It does not replace the path jail, and it does not see packet contents.

The path jail is the application rule. `path_jail` resolves the section root and the requested path, then refuses with `PathJail` unless the candidate stays inside that root. A seat follows that function. It does not invent a second jail.

WireGuard is a VPN that encapsulates IP packets over UDP. Each interface has a private key. Each peer has a public key. Cryptokey routing allows a source or destination only when it matches that peer's AllowedIPs. The project states that this provides privacy, authenticity, and perfect forward secrecy. The named suite is ChaCha20 with Poly1305 (RFC 7539 AEAD), Curve25519, BLAKE2s, SipHash24, and HKDF. Key distribution is out of scope for WireGuard. A seat may name those facts. A seat does not paste a private key, an AllowedIPs dump, or a handshake.

Cryptographic hygiene for a seat:

- A secret stays off the vault, the packet, the note, and the commit.
- A public key may be named as a peer identity. A private key may not be shown.
- Prefer the named suite above. Do not add a second cipher, a custom construction, or a "temporary" key in a prompt.
- Refuse a path that leaves the section root.
- Refuse a shell string. `run_argv` accepts an argument vector only.
- Rail changes only from the human side record.

## What den should not copy

Do not copy AppArmor profile syntax, complain-to-enforce transitions, or denial debugging. Do not copy WireGuard config examples, endpoint addresses, key material, handshake messages, or timer parameters from the protocol page. Do not copy attack steps, bypass procedures, or weaponized commands. A vendor comparison is not a den rule.

## One plan line

Confine the seat with an enforce-mode AppArmor profile, jail paths in `sections/security/match.py`, and treat WireGuard as the tunnel: public keys and AllowedIPs only, secrets never in the vault or the packet.

## Unproven

No den host was checked for an AppArmor profile, a WireGuard interface, or where keys live. Whether enforce mode is on, and whether any secret already sits in a vault, is unverified.
