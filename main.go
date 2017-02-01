package command

import (
	"bytes"
	"io"
	"strings"
)

type creator interface {
	Command(name string, arg ...string) Cmder
}

// Cmder interface wraps the os/exec Cmd struct
type Cmder interface {
	Run() error
	Start() error
	CombinedOutput() ([]byte, error)
	Output() ([]byte, error)
	SetWorkingDir(path string)
	SetStdin(stdin io.ReadCloser)
	SetStdout(stdout io.Writer)
	StderrPipe() (io.ReadCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	Wait() error
}

var run creator = &shell{}

// AsMock will run all commands as a mock object
func SetMock(mock *MockShellCmd) {
	run = &mockShell{cmd: mock}
}

// AsExec will run all commends as os/exec
func SetExec() {
	run = &shell{}
}

// PipeCommands will pipe two commands together and return the output
func PipeCommands(r1 Cmder, r2 Cmder) string {
	var buf bytes.Buffer
	stdin, _ := r1.StdoutPipe()
	r2.SetStdin(stdin)
	r2.SetStdout(&buf)
	r2.Start()
	r1.Run()
	r2.Wait()
	out := buf.String()
	if strings.HasSuffix(out, "\n") {
		out = out[0 : len(out)-1]
	}
	if strings.HasPrefix(out, "(stdin)= ") {
		out = out[9:len(out)]
	}
	return out
}

// Command mirrors the os/exec Command(name string, arg ...string) *Cmd method
func Command(name string, arg ...string) Cmder {
	return run.Command(name, arg...)
}
