import os
import tempfile
import unittest

from match import path_jail, play, run_argv, set_rail


class Attempt:
    def __init__(self, premise: str, run) -> None:
        self.premise = premise
        self.run = run


class MatchTests(unittest.TestCase):
    def test_path_outside_root_is_refused_and_written_up(self) -> None:
        with tempfile.TemporaryDirectory() as root:
            match = play(
                "path-jail",
                Attempt(
                    "Submitted a path outside the section root on this machine.",
                    lambda: path_jail(root, os.path.join(os.pardir, "outside")),
                ),
            )
        self.assertTrue(match.held)
        text = match.writeup()
        self.assertIn("Premise: Submitted a path outside the section root", text)
        self.assertIn("Defense: PathJail refused", text)
        self.assertIn("Result: pass", text)

    def test_shell_string_is_refused(self) -> None:
        match = play(
            "argv",
            Attempt(
                "Submitted a shell string on a smolvm guest we created.",
                lambda: run_argv(["tool", "a|b"]),
            ),
        )
        self.assertTrue(match.held)
        self.assertIn("RunArgv refused", match.writeup())

    def test_packet_cannot_flip_rail(self) -> None:
        match = play(
            "rail",
            Attempt(
                "A packet tried to flip Rail on this machine.",
                lambda: set_rail(False),
            ),
        )
        self.assertTrue(match.held)
        self.assertIn("Rail refused", match.writeup())

    def test_plain_argv_is_allowed(self) -> None:
        self.assertEqual(run_argv(["tool", "read"]), ["tool", "read"])


if __name__ == "__main__":
    unittest.main()
