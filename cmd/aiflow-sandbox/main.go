package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/umars28/aiflow-sandbox/internal/greet"
)

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("aiflow-sandbox", flag.ContinueOnError)
	fs.SetOutput(stderr)
	salutation := fs.String("salutation", "Hello", "salutation to address the name with")

	if err := fs.Parse(args); err != nil {
		return 2
	}

	name := strings.Join(fs.Args(), " ")
	if strings.TrimSpace(name) == "" {
		fmt.Fprintln(stderr, "aiflow-sandbox: a non-blank name is required")
		return 2
	}

	fmt.Fprintln(stdout, greet.Greet(name, *salutation))
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
