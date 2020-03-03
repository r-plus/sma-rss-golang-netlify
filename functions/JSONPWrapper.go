package main

import (
	"fmt"
	"io"
)

// JSONPWrapper unmershall to struct from jsonp
type JSONPWrapper struct {
	Prefix     string
	Underlying io.Reader

	gotPrefix bool
}

func (jpw *JSONPWrapper) Read(b []byte) (int, error) {
	if jpw.gotPrefix {
		return jpw.Underlying.Read(b)
	}

	prefix := make([]byte, len(jpw.Prefix))
	n, err := io.ReadFull(jpw.Underlying, prefix)
	if err != nil {
		return n, err
	}

	if string(prefix) != jpw.Prefix {
		return n, fmt.Errorf("JSONP prefix mismatch: expected %q, got %q",
			jpw.Prefix, prefix)
	}

	// read until the (; in general, this should just be one read
	char := make([]byte, 1)
	for char[0] != '(' {
		n, err = jpw.Underlying.Read(char)
		if n == 0 || err != nil {
			return n, err
		}
	}

	// We've now consumed the JSONP prefix.
	jpw.gotPrefix = true
	return jpw.Underlying.Read(b)
}
