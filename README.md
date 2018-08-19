## indent

[![Build Status](https://travis-ci.org/wkhere/indent.svg?branch=master)](https://travis-ci.org/wkhere/indent)
[![Coverage Status](https://coveralls.io/repos/github/wkhere/indent/badge.svg?branch=master)](https://coveralls.io/github/wkhere/indent?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/wkhere/indent)](https://goreportcard.com/report/github.com/wkhere/indent)


This Go library gives you an `indent.Reader` which will prepend each line
of underlying input with configured indentation.

### API

You create the reader as the doc says:
```
% go doc indent.NewReader
package indent // import "github.com/wkhere/indent"

func NewReader(r io.Reader, indent string) *Reader
    NewReader returns a Reader which will prepend each line read from underlying
    reader with given indentation.
```

then you use it as any other `io.Reader`.

### example

This shows an example of `indent` usefulness.

The program:
```
package main

import (
	"bytes"
	"io"
	"os"

	"github.com/wkhere/indent"
)

func main() {
	b := bytes.NewBufferString(`
hey: I am yaml
key: val
    nested: something`)

	io.Copy(os.Stdout, bytes.NewBufferString("new_root:\n"))
	io.Copy(os.Stdout, indent.NewReader(b, "    "))
}
```

produces the output:
```
new_root:
    
    hey: I am yaml
    key: val
        nested: something
```

