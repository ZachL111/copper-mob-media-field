# copper-mob-media-field

`copper-mob-media-field` keeps a focused Go implementation around mobile workflows. The project goal is to create a Go reference implementation for media workflows, centered on policy evaluation, deny and allow fixtures, and explainable decision traces.

## Use Case

This is intentionally local and self-contained so it can be inspected without credentials, services, or seeded history.

## Copper Mob Media Field Review Notes

Start with `form pressure` and `local state`. Those cases create the widest score spread in this repo, so they are the best quick check when the model changes.

## Highlights

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/copper-mob-media-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `form pressure` and `local state`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Code Layout

The repository has two validation layers: the original compact policy fixture and the domain review fixture. They are separate so one can change without hiding failures in the other.

The Go implementation avoids hidden state so fixture changes are easy to reason about.

## Run The Check

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Regression Path

The check exercises the source code and the review fixture. `stale` is the high score at 227; `edge` is the low score at 145.

## Future Work

This remains a local project with deterministic fixtures. It does not depend on credentials, hosted services, or live data. Future work should add richer malformed inputs before widening the public API.
