## indent

This Go library gives you an `indent.Reader` which will prepend each line
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

