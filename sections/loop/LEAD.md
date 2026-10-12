# Turn loop and four exits

Draft detail under `plan/foundation.md`. The four exits are the only ways a turn ends.

## Module

One packet in. One result or one refusal out. The turn does not resume.

## Exits

1. Done. Write is filled and the check passed. The seat returns.
2. Refused. `path_jail`, `run_argv`, or `set_rail` refused. The row records the control name. The seat returns.
3. No progress. The same read set and argv already ran, two notes added nothing, or the seat left the Function. The seat returns the packet to the operator.
4. Cut. A lead cuts Next. The mesh keeps the current section and a refusal row. The cut is not a new section.

## Invariants

A thought that does not change the packet is not an exit and not a step. The seat does not start a second call to escape exit 3. Rail is not an exit the seat can take.

## Files

`plan/foundation.md` holds the runtime and the no-progress table. `sections/security/match.py` holds the three controls.

## Adopt

Record the exit name on the row. Exit 3 names the field that did not change.

## Refuse

A repair of cut-off arguments. A retry after a refusal. A loop whose only output is another thought. A child started so the parent can avoid an exit.
