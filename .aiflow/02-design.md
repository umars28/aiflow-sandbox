# Greet command with a configurable salutation

## Public surface

`Greet(name, salutation string) string` in `internal/greet`, and a `-salutation` flag on the CLI.

## Modules touched

- `internal/greet/greet.go` new, holds the formatting rule
- `cmd/aiflow-sandbox/main.go` changed, parses the flag and prints the result

## Failure mode

An empty name produces a greeting addressed to nobody, which reads as a bug to whoever sees it in a log. `Greet` returns the salutation alone when the name is blank, and the CLI refuses an all-whitespace name with exit code 2 rather than printing something meaningless.

## Options rejected

A template string flag such as `-format "Hello, {name}"` was rejected: it moves validation to runtime and invites format-string mistakes for a feature with exactly one variable.
