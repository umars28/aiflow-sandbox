package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr bool
	}{
		{"default salutation", []string{"Ada"}, 0, "Hello, Ada\n", false},
		{"custom salutation", []string{"-salutation", "Howdy", "Ada"}, 0, "Howdy, Ada\n", false},
		{"whitespace name", []string{"   "}, 2, "", true},
		{"missing name", nil, 2, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)

			if code != tt.wantCode {
				t.Errorf("run(%q) = %d, want %d", tt.args, code, tt.wantCode)
			}
			if stdout.String() != tt.wantStdout {
				t.Errorf("stdout = %q, want %q", stdout.String(), tt.wantStdout)
			}
			if gotStderr := strings.TrimSpace(stderr.String()) != ""; gotStderr != tt.wantStderr {
				t.Errorf("stderr non-empty = %v, want %v (stderr = %q)", gotStderr, tt.wantStderr, stderr.String())
			}
		})
	}
}
