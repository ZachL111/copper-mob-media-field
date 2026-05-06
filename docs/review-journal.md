# Review Journal

The review surface for `copper-mob-media-field` is deliberately narrow: one fixture, one scoring rule, and one local check.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 190, lane `ship`
- `stress`: `sync drift`, score 177, lane `ship`
- `edge`: `local state`, score 145, lane `ship`
- `recovery`: `conflict cost`, score 183, lane `ship`
- `stale`: `form pressure`, score 227, lane `ship`

## Note

This file is intentionally plain so the fixture remains the source of truth.
