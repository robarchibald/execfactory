package execfactory

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"syscall"
)

// OS is an instance of an OS creator
var OS Creator = &osCreator{}

// Mock is an instance of a Mock Creator
var Mock Creator = &mockCreator{}

// Creator is the interface used to create either a mock or real os/exec Cmd
type Creator interface {
	Command(name string, arg ...string) Cmder
	CommandContext(ctx context.Context, name string, arg ...string) Cmder
}

// Cmder interface wraps the os/exec Cmd struct
type Cmder interface {
	Run() error
	Start() error
	String() string
	CombinedOutput() ([]byte, error)
	Output() ([]byte, error)
	StdinPipe() (io.WriteCloser, error)
	StderrPipe() (io.ReadCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	Wait() error
	cmdGetter
	cmdSetter
}

type cmdSetter interface {
	SetPath(path string)
	SetArgs(args []string)
	SetEnv(env []string)
	SetDir(path string)
	SetStdin(stdin io.Reader)
	SetStdout(stdout io.Writer)
	SetStderr(stderr io.Writer)
	SetExtraFiles(files []*os.File)
	SetSysProcAttr(attr *syscall.SysProcAttr)
	SetProcess(process *os.Process)
	SetProcessState(processState *os.ProcessState)
}

type cmdGetter interface {
	GetPath() string
	GetArgs() []string
	GetEnv() []string
	GetDir() string
	GetStdin() io.Reader
	GetStdout() io.Writer
	GetStderr() io.Writer
	GetExtraFiles() []*os.File
	GetSysProcAttr() *syscall.SysProcAttr
	GetProcess() *os.Process
	GetProcessState() *os.ProcessState
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
