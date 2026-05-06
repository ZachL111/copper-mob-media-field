# Copper Mob Media Field Walkthrough

I use this file as a small checklist before changing the Go implementation.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 190 | ship |
| stress | sync drift | 177 | ship |
| edge | local state | 145 | ship |
| recovery | conflict cost | 183 | ship |
| stale | form pressure | 227 | ship |

Start with `stale` and `edge`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The next useful expansion would be a malformed fixture around sync drift and conflict cost.
