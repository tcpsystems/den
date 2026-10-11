# Exercises

A match is a test of a control on a host we own. Seats document the limit. They do not drop it.

Two seats play. The hill is one control. The score is whether that control held. It is not whether an attack landed.

## Roles

| Seat | Writes | Does not write |
| --- | --- | --- |
| Premise | One sentence of what an attacker would want, and which owned host | Steps, commands, payloads, packet captures, or a bypass |
| Defense | The control, the test that shows it holds, and the result | A weaker rule, a deleted limit, or a way around the control |

The premise is the attacker thought, kept to intent and impact. The defense is the function. The match runs: the premise seat submits the disallowed action to that function, and the defense seat returns the refusal. The spectator writeup is those two actions. It is produced by `sections/security/match.py`.

## Owned range

- This machine, including smolvm guests on the `harness` network
- Services we already run on `tcp.systems`
- A Hetzner lab or staging host in the foxcore project
- A host named in a written permission, for the hours named there

A competitor, a customer, and a public game panel are not in range.

## One match

1. Hill. Name one control. Path jail, argv, Rail, fail closed, or a firewall rule we already run.
2. Premise. One sentence. What we are afraid that control might miss.
3. Host. The owned name.
4. Test. A unit test or a config read. It fails if the control is absent. It does not perform the attack.
5. Stop. One run, or the time on the written permission.
6. Result. Pass or fail, and the path of the test.
7. Safety. Copy the limits that apply. Do not edit them.

## King of the hill

Three matches, one hill. Each premise seat writes a different fear of the same control. The defense seat runs the three tests. The hill holds if all three tests pass. A failed test is a bug in the control, fixed by a smaller function and a reeval. It is not a prize for the premise seat.

## What gets expanded

A new fear becomes a new six-line match and a new test. The catalog of fears grows. The catalog of steps does not. Platform security work is more controls with tests, not more ways to break them.

## Safety

These limits are copied, not revised:

- One direct model id. OmniRoute is a named tool.
- Argv only. No shell.
- Paths stay inside the section root.
- A prompt cannot flip Rail.
- A child cannot write the parent's Write, Next, or Rail.
- No secrets in the packet or the repo.
- No exploit steps, payloads, or flood generation.
- A test runs only on a host we own, or a range named in written permission.
