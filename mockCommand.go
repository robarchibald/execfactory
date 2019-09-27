package execfactory

import (
	"context"
	"fmt"
	"io"
	"os"
	"syscall"
)

type mockCreator struct {
	instances []MockInstance
}

func (c *mockCreator) Command(name string, arg ...string) Cmder {
	var instance MockInstance
	if len(c.instances) > 0 {
		instance = c.instances[0]
		c.instances = c.instances[1:]
	}
	return &mockCmd{name: name, args: arg, mock: instance}
}

func (c *mockCreator) CommandContext(ctx context.Context, name string, arg ...string) Cmder {
	return &mockCmd{name: name, args: arg}
}

// MockInstance defines the returned values of the mock Cmd method calls
type MockInstance struct {
	RunErr            error
	StartErr          error
	CombinedOutputVal []byte
	CombinedOutputErr error
	OutputVal         []byte
	OutputErr         error
	StdinPipeVal      io.WriteCloser
	StdinPipeErr      error
	StderrPipeVal     io.ReadCloser
	StderrPipeErr     error
	StdoutPipeVal     io.ReadCloser
	StdoutPipeErr     error
	WaitErr           error
}

type mockCmd struct {
	Cmder
	mock          MockInstance
	name          string
	path          string
	args          []string
	env           []string
	dir           string
	stdin         io.Reader
	stdout        io.Writer
	stderr        io.Writer
	extraFiles    []*os.File
	sysProcAttr   *syscall.SysProcAttr
	process       *os.Process
	processState  *os.ProcessState
	methodsCalled []MethodCall
}

// MethodCall is used to keep track of all the methods called by the mock Cmd
type MethodCall struct {
	Name string
	Args []interface{}
}

func (c *mockCmd) String() string {
	return fmt.Sprintf("%s, %v", c.name, c.args)
}
func (c *mockCmd) Run() error {
	return c.mock.RunErr
}
func (c *mockCmd) Start() error {
	return c.mock.StartErr
}
func (c *mockCmd) CombinedOutput() ([]byte, error) {
	return c.mock.CombinedOutputVal, c.mock.CombinedOutputErr
}
func (c *mockCmd) Output() ([]byte, error) {
	return c.mock.OutputVal, c.mock.OutputErr
}
func (c *mockCmd) StdinPipe() (io.WriteCloser, error) {
	return c.mock.StdinPipeVal, c.mock.StdinPipeErr
}
func (c *mockCmd) StderrPipe() (io.ReadCloser, error) {
	return c.mock.StderrPipeVal, c.mock.StderrPipeErr
}
func (c *mockCmd) StdoutPipe() (io.ReadCloser, error) {
	return c.mock.StdoutPipeVal, c.mock.StdoutPipeErr
}
func (c *mockCmd) Wait() error {
	return c.mock.WaitErr
}

// Sets

func (c *mockCmd) SetPath(path string) {
	c.path = path
}
func (c *mockCmd) SetArgs(args []string) {
	c.args = args
}
func (c *mockCmd) SetEnv(env []string) {
	c.env = env
}
func (c *mockCmd) SetDir(path string) {
	c.dir = path
}
func (c *mockCmd) SetStdin(stdin io.Reader) {
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

func (c *mockCmd) GetPath() string {
	return c.path
}
func (c *mockCmd) GetArgs() []string {
	return c.args
}
func (c *mockCmd) GetEnv() []string {
	return c.env
}
func (c *mockCmd) GetDir() string {
	return c.dir
}
func (c *mockCmd) GetStdin() io.Reader {
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
func (c *mockCmd) GetMethodsCalled() []MethodCall {
	return c.methodsCalled
}
