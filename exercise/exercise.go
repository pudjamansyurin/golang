package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct {
	char byte
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = r.char
	}
	return 1, nil
}

func main() {
	reader.Validate(MyReader{char: 'A'})
}
