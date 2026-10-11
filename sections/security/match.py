"""Run a match against a control we own.

The premise seat submits one disallowed action. The defense seat answers.
The spectator writeup is those two actions and whether the control held.
"""

from __future__ import annotations

import os
from dataclasses import dataclass


class Refused(Exception):
    def __init__(self, control: str) -> None:
        self.control = control
        super().__init__(control)


def path_jail(root: str, requested: str) -> str:
    """Open a path only when it stays inside the section root."""
    base = os.path.realpath(root)
    candidate = os.path.realpath(os.path.join(base, requested))
    if candidate != base and not candidate.startswith(base + os.sep):
        raise Refused("PathJail")
    return candidate


def run_argv(argv: list[str]) -> list[str]:
    """Accept an argument vector. Refuse a shell string."""
    if len(argv) == 0:
        raise Refused("RunArgv")
    for arg in argv:
        if any(ch in arg for ch in "|&;<>\n"):
            raise Refused("RunArgv")
    return list(argv)


def set_rail(from_human: bool) -> str:
    """Rail changes only when the human side record says so."""
    if not from_human:
        raise Refused("Rail")
    return "on"


@dataclass(frozen=True)
class Action:
    seat: str
    did: str


@dataclass(frozen=True)
class Match:
    hill: str
    premise: Action
    defense: Action
    held: bool

    def writeup(self) -> str:
        result = "pass" if self.held else "fail"
        return (
            f"# {self.hill}\n"
            f"Premise: {self.premise.did}\n"
            f"Defense: {self.defense.did}\n"
            f"Result: {result}\n"
        )


def play(hill: str, attempt) -> Match:
    premise = Action("premise", attempt.premise)
    try:
        attempt.run()
    except Refused as refused:
        return Match(
            hill,
            premise,
            Action("defense", f"{refused.control} refused. The action did not proceed."),
            True,
        )
    return Match(
        hill,
        premise,
        Action("defense", "The control let the action proceed."),
        False,
    )
