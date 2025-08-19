package main

import (
	"fmt"
	"io"
	"os"

	"github.com/wkhere/indent"
)

type pargs struct {
	file string
	help func()
}

func parseArgs(args []string) (a pargs, _ error) {
	const usage = `usage: indent [FILE|-]  - reads FILE or stdin, prints to stdout`

	var rest []string
flags:
	for ; len(args) > 0; args = args[1:] {
		switch arg := args[0]; {

		case arg == "-h", arg == "--help":
			a.help = func() { fmt.Println(usage) }
			return a, nil

		case arg == "--":
			rest = append(rest, args[1:]...)
			break flags

		case len(arg) > 1 && arg[0] == '-':
			return a, fmt.Errorf("unknown flag %s\n%s", arg, usage)

		default:
			rest = append(rest, arg)
		}
	}

	switch len(rest) {
	case 0:
		a.file = "-"
	case 1:
		a.file = rest[0]
	case 2:
		return a, fmt.Errorf("too many file args\n%s", usage)
	}

	return a, nil
}

func open(file string) (*os.File, error) {
	if file == "-" {
		return os.Stdin, nil
	}
	return os.Open(file)
}

func run(a *pargs) error {
	f, err := open(a.file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := indent.NewReader(f, "\t")
	_, err = io.Copy(os.Stdout, r)
	return err
}

func main() {
	a, err := parseArgs(os.Args[1:])
	if err != nil {
		die(2, err)
	}
	if a.help != nil {
		a.help()
		os.Exit(0)
	}

	err = run(&a)
	if err != nil {
		die(1, err)
	}
}

func die(exitcode int, err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(exitcode)
}
