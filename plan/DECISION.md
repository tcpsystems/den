# Decision

A broad plan does not pick a direction by volume. Three clusters write three proposals. One judge keeps one.

## What a proposal must contain

- The choice in one sentence.
- The file or command that supports each claim. A claim with no path is dropped.
- What the current code does, and what a document claims that the code does not do.
- One operator scenario: the command they type and the line they read back.
- What this choice refuses to add.

## How the judge scores

| Check | Pass |
| --- | --- |
| Evidence | Every kept sentence cites a path. |
| Fit | The sentence matches a function, a flag, or a test in that path. |
| Split | Two commands cannot be confused from the help line alone. |
| Size | Help for one command is two lines or fewer. |
| Safety | No push, no secret value, no model call unless that call is in the file. |

The highest score is implemented. A tie goes to the smaller change. The other proposals stay on their branches as the record.

## What is not a win

A longer document. A new name with no migration. A feature the proposal wishes the binary had.
