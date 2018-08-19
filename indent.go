package indent

import (
	"bufio"
	"io"
)

// Reader implements prepending each line with indentation.
type Reader struct {
	indentb     []byte
	llr         *bufio.Reader
	head, data  []byte
	startIndent bool
	err         error
}

// NewReader returns a Reader which will prepend each line read from
// underlying reader with given indentation.
func NewReader(r io.Reader, indent string) *Reader {
	return &Reader{
		indentb:     []byte(indent),
		llr:         bufio.NewReader(r),
		startIndent: true,
	}
}

// Read reads data into p, prepending each line with indentation.
// It may return (0, nil) in the middle of processing, then the subsequent
// call will read more data.
// At EOF the count may be > 0.
func (r *Reader) Read(p []byte) (n int, err error) {

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
	// Let's read the next line or at least part of it.
	r.data, r.err = r.llr.ReadSlice('\n')

	// No data read, end this.
	if len(r.data) == 0 {
		return 0, r.err
	}

	// Some data read. If we were at the beginning of the line,
	// set new indent to be filled before data.
	if r.startIndent {
		r.head = r.indentb
		r.startIndent = false
	}

	switch r.err {
	case nil:
		// Data contains full line, so after flushing it,
		// new indent should occur. Mark this for future calls.
		r.startIndent = true

	case bufio.ErrBufferFull:
		// Data doesn't have full line but we can ignore this,
		// more data will be read in future calls.
		r.err = nil
	}

	// Any other error is copied to r.err and first we should
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
