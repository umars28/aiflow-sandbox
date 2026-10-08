# Slices

Source: `.aiflow/02-design.md`. Test command for every slice: `go test ./...` (from `.aiflow/project.yaml`).

| # | Slice | Primary files | Test that proves it |
|---|---|---|---|
| 1 | `Greet(name, salutation string) string` in a new `internal/greet` package: returns `"<salutation>, <name>"`, and the salutation alone when the name is blank | `internal/greet/greet.go` (new), `internal/greet/greet_test.go` (new) | `go test ./...` — table test over a populated name, an empty name, and a whitespace-only name, asserting the blank cases return the bare salutation with no trailing comma |
| 2 | CLI wiring: a `-salutation` flag feeding `Greet`, and rejection of an all-whitespace name with exit code 2 | `cmd/aiflow-sandbox/main.go`, `cmd/aiflow-sandbox/main_test.go` (new) | `go test ./...` — a test on the extracted `run(args, stdout, stderr) int` helper asserting exit code 2 plus a message on stderr for a whitespace name, exit code 0 and the formatted line on stdout otherwise, and that a custom `-salutation` reaches the output |

Slice 1 comes first because it is the only piece carrying the formatting rule and the blank-name decision; it is pure, has no flag parsing in the way, and is what slice 2 calls. Shipping it alone leaves the binary's current behaviour untouched and the suite green, so it is genuinely releasable on its own. Slice 2 then has nothing left to invent beyond argument handling. The risk sits in slice 2: exit codes and `flag` package behaviour are only testable if `main` stays a thin wrapper around a `run` function that returns an int rather than calling `os.Exit` directly — if that extraction is skipped, the exit-code 2 requirement from the design's failure-mode section ends up with no test behind it, and the slice stops satisfying the first rule. Note also that the two slices are ordered, not independent: slice 2 imports slice 1. Nothing depends backwards on a later slice.
