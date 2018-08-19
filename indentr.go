package indentr

import (
	"bufio"
	"io"
)

type IndentR struct {
	indent     []byte
	llr        *bufio.Reader
	head, data []byte
	err        error
}

func NewIndentR(r io.Reader, spaces string) *IndentR {
	return &IndentR{
		indent: []byte(spaces),
		llr:    bufio.NewReader(r),
	}
}

func (r *IndentR) Read(p []byte) (n int, err error) {

	// First, return saved indent, or its chunk if p is smaller.
	if len(r.head) > 0 {
		n = copy(p, r.head)
		r.head = r.head[n:]
		return
	}

	// Then return saved data, or its chunk if p is smaller.
	// When all saved data is flushed this way, return also
	// possible error recorded with that data.
	if len(r.data) > 0 {
		n = copy(p, r.data)
		r.data = r.data[n:]
		err = errorIfFlushed(len(r.data), r.err)
		return
	}

	// Now the previous indent and line are flushed.
	// Let's read the next line.
	r.data, r.err = r.llr.ReadBytes('\n')

	// No data read, end this.
	if len(r.data) == 0 {
		return 0, r.err
	}

	// Data read. It's either line ending with LF,
	// or a data at the EOF without LF.
	// In both cases we want to indent it.
	r.head = r.indent

	// Any error is copied to r.err and first we should
	// start returning data alread gathered in r.head/r.line.
	// It can be done in future calls, now we can safely
	// return 0, nil.
	return
}

func errorIfFlushed(l int, err error) error {
	if l > 0 {
		return nil
	}
	return err
}
