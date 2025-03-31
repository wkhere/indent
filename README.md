## indent

[![Build Status](https://github.com/wkhere/indent/workflows/Go/badge.svg?branch=master)](https://github.com/wkhere/indent/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/wkhere/indent/badge.svg?branch=master&kill_cache=1)](https://coveralls.io/github/wkhere/indent?branch=master)


`indent.Reader` will prepend each line
of underlying input with configured indentation.

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

