package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		return b + 13
	} else if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
		return b - 13
	}
	return b
}

func (t *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = t.r.Read(b)

	if err == nil {
		for i := range b {
			b[i] = rot13(b[i])
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
