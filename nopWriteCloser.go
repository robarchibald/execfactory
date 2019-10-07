package execfactory

import "io"

type nopWriteCloser struct {
	w io.Writer
	io.WriteCloser
}

func newNopWriteCloser(w io.Writer) io.WriteCloser {
	return &nopWriteCloser{w: w}
}

func (w *nopWriteCloser) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w *nopWriteCloser) Close() error {
	return nil
}
