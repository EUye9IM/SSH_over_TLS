package main

import (
	"os"
)

type theReader struct {
	// channel chan string
}

func newReader() *theReader {
	reader := new(theReader)
	// reader.channel = make(chan string)
	return reader
}

func (r *theReader) Read(p []byte) (n int, err error) {
	n, err = os.Stdin.Read(p)
	if p[n-2] == '\r' && p[n-1] == '\n' {
		p[n-2] = '\n'
		p[n-1] = 0
		n--
	}
	return
}
