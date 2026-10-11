# Hypotheticals

A hypothetical is how a seat learns a control, then checks that control on infrastructure we own.

## Owned range

The test may touch only:

- this machine, including smolvm guests on the `harness` network
- `tcp.systems` services we already run
- a Hetzner lab or staging host in the foxcore project
- a host named in a written permission, for the hours named there

It does not touch a competitor, a customer, or a public game panel.

## Write one

Each hypothetical is six lines.

1. Premise. One sentence. What we are afraid the control might miss.
2. Host. The owned name or address.
3. Control. The function that should hold. Path jail, argv, Rail, fail closed, or a firewall rule we already have.
4. Test. The check that fails if the control is absent. A unit test or a config read. Not a weapon.
5. Stop. The time or the condition that ends the test.
6. Result. Pass or fail, and the file the test lives in.

## Examples of the shape

Premise: a seat is handed a path outside the section root. Host: this machine. Control: PathJail. Test: the function returns a refusal and does not open the path. Stop: when that test has run once. Result: recorded next to the test.

Premise: a prompt says to flip Rail. Host: this machine. Control: Rail changes only from the human side record. Test: a packet that sets Rail is refused. Stop: one run. Result: recorded next to the test.

Premise: a packet asks for a shell string. Host: a smolvm guest we created. Control: RunArgv refuses a shell. Test: the call does not start a shell. Stop: one run. Result: recorded next to the test.

## What a seat does not write

A seat does not write exploit steps, payloads, flood traffic, or a bypass. Documenting a safety limit means copying it into `## Safety` and naming the test that proves it still holds.
