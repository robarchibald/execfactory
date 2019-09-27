package exec

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

// MockCommand generates a testable Cmder interface
func MockCommand(name string, arg ...string) Cmder {
	return &mockCmd{name: name, args: arg}
}

type mockCmd struct {
	Cmder
	RunErr            error
	StartErr          error
	CombinedOutputVal []byte
	CombinedOutputErr error
	OutputVal         []byte
	OutputErr         error
	workingdir        string
	stdinPipeVal      io.WriteCloser
	stdinPipeErr      error
	StderrPipeVal     io.ReadCloser
	StderrPipeErr     error
	StdoutPipeVal     io.ReadCloser
	StdoutPipeErr     error
	WaitErr           error

	name         string
	path         string
	args         []string
	env          []string
	dir          string
	stdin        io.Reader
	stdout       io.Writer
	stderr       io.Writer
	extraFiles   []*os.File
	sysProcAttr  *syscall.SysProcAttr
	process      *os.Process
	processState *os.ProcessState
}

func (c *mockCmd) String() string {
	return fmt.Sprintf("%s, %v", c.name, c.args)
}
func (c *mockCmd) Run() error {
	return c.RunErr
}
func (c *mockCmd) Start() error {
	return c.StartErr
}
func (c *mockCmd) CombinedOutput() ([]byte, error) {
	return c.CombinedOutputVal, c.CombinedOutputErr
}
func (c *mockCmd) Output() ([]byte, error) {
	return c.OutputVal, c.OutputErr
}
func (c *mockCmd) stdinPipe() (io.WriteCloser, error) {
	return c.stdinPipeVal, c.stdinPipeErr
}
func (c *mockCmd) StderrPipe() (io.ReadCloser, error) {
	return c.StderrPipeVal, c.StderrPipeErr
}
func (c *mockCmd) StdoutPipe() (io.ReadCloser, error) {
	return c.StdoutPipeVal, c.StdoutPipeErr
}
func (c *mockCmd) Wait() error {
	return c.WaitErr
}

// Sets

func (c *mockCmd) Setpath(path string) {
	c.path = path
}
func (c *mockCmd) Setargs(args []string) {
	c.args = args
}
func (c *mockCmd) Setenv(env []string) {
	c.env = env
}
func (c *mockCmd) Setdir(path string) {
	c.dir = path
}
func (c *mockCmd) Setstdin(stdin io.Reader) {
	c.stdin = stdin
}
func (c *mockCmd) SetStdout(stdout io.Writer) {
	c.stdout = stdout
}
func (c *mockCmd) SetStderr(stderr io.Writer) {
	c.stderr = stderr
}
func (c *mockCmd) SetExtraFiles(files []*os.File) {
	c.extraFiles = files
}
func (c *mockCmd) SetSysProcAttr(attr *syscall.SysProcAttr) {
	c.sysProcAttr = attr
}
func (c *mockCmd) SetProcess(process *os.Process) {
	c.process = process
}
func (c *mockCmd) SetProcessState(processState *os.ProcessState) {
	c.processState = processState
}

// Gets

func (c *mockCmd) Getpath() string {
	return c.path
}
func (c *mockCmd) Getargs() []string {
	return c.args
}
func (c *mockCmd) Getenv() []string {
	return c.env
}
func (c *mockCmd) Getdir() string {
	return c.dir
}
func (c *mockCmd) Getstdin() io.Reader {
	return c.stdin
}
func (c *mockCmd) GetStdout() io.Writer {
	return c.stdout
}
func (c *mockCmd) GetStderr() io.Writer {
	return c.stderr
}
func (c *mockCmd) GetExtraFiles() []*os.File {
	return c.extraFiles
}
func (c *mockCmd) GetSysProcAttr() *syscall.SysProcAttr {
	return c.sysProcAttr
}
func (c *mockCmd) GetProcess() *os.Process {
	return c.process
}
func (c *mockCmd) GetProcessState() *os.ProcessState {
	return c.processState
}
