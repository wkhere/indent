## indent

[![Build Status](https://github.com/wkhere/indent/workflows/Go/badge.svg?branch=master)](https://github.com/wkhere/indent/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/wkhere/indent/badge.svg?branch=master&kill_cache=1)](https://coveralls.io/github/wkhere/indent?branch=master)


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
	b := bytes.NewBufferString(example)
	io.Copy(os.Stdout, bytes.NewBufferString("new_root:\n"))
	io.Copy(os.Stdout, indent.NewReader(b, "    "))
}

const example = `
hey: I am yaml
key: val
	nested: something
`
```

produces the output:
```
new_root:
    
    hey: I am yaml
    key: val
        nested: something
```

