package indent

import (
	"io"
	"os"
	"strings"
	"testing"
)

func init() {
	os.Chdir("testdata")
}

var tcBasic = []struct {
	input, want string
	nwant       int64
}{
	{"", "", 0},
	{"\n", "XX\n", 3},
	{"x", "XXx", 3},
	{"x\n", "XXx\n", 4},
	{"\nx", "XX\nXXx", 6},
	{"\n\n\n", "XX\nXX\nXX\n", 9},
	{"\taa\n", "XX\taa\n", 6},
	{"a:\n  b: 42\n", "XXa:\nXX  b: 42\n", 15},
	{"a:\n  b: 42", "XXa:\nXX  b: 42", 14},
}

func TestBasic(t *testing.T) {

	for i, tc := range tcBasic {
		b := strings.NewReader(tc.input)
		b2 := new(strings.Builder)
		r := NewReader(b, "XX")
		n, err := io.Copy(b2, r)

		if err != nil {
			t.Errorf("tc[%d] unexpected error: %s", i, err)
		}

		have := b2.String()
		if have != tc.want {
			t.Errorf("tc[%d] mismatch\nhave: %q\nwant: %q", i, have, tc.want)
		}
		if n != tc.nwant {
			t.Errorf("tc[%d] n mismatch\nhave: %d, want: %d", i, n, tc.nwant)
		}
	}
}

func TestSmallBuffer(t *testing.T) {

	p := make([]byte, 1)

	for i, tc := range tcBasic {
		b := strings.NewReader(tc.input)
		b2 := new(strings.Builder)
		r := NewReader(b, "XX")

		var n int64
		var err error
		for {
			var n1 int
			n1, err = r.Read(p)
			b2.Write(p[:n1])
			n += int64(n1)
			if err != nil {
				break
			}
		}

		if err != nil && err != io.EOF {
			t.Errorf("tc[%d] unexpected error: %s", i, err)
		}

		have := b2.String()
		if have != tc.want {
			t.Errorf("tc[%d] mismatch\nhave: %q\nwant: %q", i, have, tc.want)
		}
		if n != tc.nwant {
			t.Errorf("tc[%d] n mismatch\nhave: %d, want: %d", i, n, tc.nwant)
		}
	}
}

func TestFiles(t *testing.T) {
	files := []string{
		"tree01.yml",
		"tree01sub1.yml",
		"tree01sub2.yml",
		"tree01sub3.yml",
		"tree01sub4.yml",
		"tree01sub5.yml",
		"tree01sub6.yml",
		"tree01sub7.yml",
		"tree02.yml",
		"tree02sub8.yml",
		"big4k+1.yml",
		"big8k+1.yml",
		"big16k+1.yml",
		"badbin01",
		"badbin1k",
		"badbin4k+1",
		"badbin8k+1",
		"badbin16k+1",
	}
	for i, fn := range files {
		err := readAndDiscardFile(fn)
		if err != nil {
			t.Errorf("tc[%d] error: %s", i, err)
		}
	}
}

func BenchmarkBasic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range tcBasic {
			b := strings.NewReader(tc.input)
			r := NewReader(b, "XX")
			io.Copy(io.Discard, r)
		}
	}
}

func BenchmarkFiles(b *testing.B) {
	files := []string{
		"tree01.yml",
		"tree01sub1.yml",
		"tree01sub2.yml",
		"tree01sub3.yml",
		"tree01sub4.yml",
		"tree01sub5.yml",
		"tree01sub6.yml",
		"tree01sub7.yml",
		"tree02.yml",
		"tree02sub8.yml",
	}
	for n := 0; n < b.N; n++ {
		for _, fn := range files {
			readAndDiscardFile(fn)
		}
	}
}

func BenchmarkBig(b *testing.B) {
	files := []string{
		"big4k+1.yml",
		"big8k+1.yml",
		"big16k+1.yml",
	}
	for n := 0; n < b.N; n++ {
		for _, fn := range files {
			readAndDiscardFile(fn)
		}
	}
}

func BenchmarkBad(b *testing.B) {
	files := []string{
		"badbin01",
		"badbin1k",
	}
	for n := 0; n < b.N; n++ {
		for _, fn := range files {
			readAndDiscardFile(fn)
		}
	}
}

func BenchmarkBadBig(b *testing.B) {
	files := []string{
		"badbin4k+1",
		"badbin8k+1",
		"badbin16k+1",
	}
	for n := 0; n < b.N; n++ {
		for _, fn := range files {
			readAndDiscardFile(fn)
		}
	}
}

func readAndDiscardFile(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := NewReader(f, "  ")
	_, err = io.Copy(io.Discard, r)
	return err
}
