package main

import (
	"fmt"
	"io"
)

type Header struct {
	Key, Value string
}
type Status struct {
	Code   int
	Reason string
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}
	for _, h := range headers {
		_, err = fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}

	if _, err = fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)
	return err
}

type ErrWriter struct {
	io.Writer
	error
}

func (w *ErrWriter) Write(data []byte) (int, error) {
	if w.error != nil {
		return 0, w.error
	}
	var n int
	n, w.error = w.Writer.Write(data)
	return n, w.error
}

func WriteResponseOptimize(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &ErrWriter{Writer: w}

	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprint(ew, "\r\n")

	io.Copy(ew, body)

	return ew.error
}

