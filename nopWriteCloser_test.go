package execfactory

import (
	"strings"
	"testing"
)

func TestNopWriteCloser(t *testing.T) {
	var buf strings.Builder
	w := newNopWriteCloser(&buf)
	w.Write([]byte("hello"))
	w.Close()
	if v := buf.String(); v != "hello" {
		t.Error("Expected buffer to be filled correctly", v)
	}
}
