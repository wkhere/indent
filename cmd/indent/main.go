package main

import (
	"fmt"
	"io"
	"os"

	"github.com/wkhere/indent"
)

type pargs struct {
	help func()
}

func parseArgs(args []string) (p pargs, _ error) {
	const usage = `usage: indent  (reads stdin, outputs to stdout)`
	if len(args) > 0 {
		if len(args) == 1 {
			switch arg := args[0]; {
			case arg == "-h", arg == "--help":
				p.help = func() { fmt.Println(usage) }
				return p, nil
			}
		}
		return p, fmt.Errorf(usage)
	}
	return p, nil
}

func run() error {
	r := indent.NewReader(os.Stdin, "\t")
	_, err := io.Copy(os.Stdout, r)
	return err
}

func main() {
	p, err := parseArgs(os.Args[1:])
	if err != nil {
		die(2, err)
	}
	if p.help != nil {
		p.help()
		os.Exit(0)
	}

	err = run()
	if err != nil {
		die(1, err)
	}
}

func die(exitcode int, err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(exitcode)
}
