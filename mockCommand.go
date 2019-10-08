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
	return c.newMockCmd(name, arg...)
}

func (c *mockCreator) CommandContext(ctx context.Context, name string, arg ...string) Cmder {
	return c.newMockCmd(name, arg...)
}

func (c *mockCreator) newMockCmd(name string, arg ...string) *mockCmd {
	var instance MockInstance
	if len(c.instances) > 0 {
		instance = c.instances[0]
		c.instances = c.instances[1:]
	}
	return &mockCmd{path: name, args: arg, mock: instance}
}

// MockInstance defines the returned values of the mock Cmd method calls
type MockInstance struct {
	RunErr                 error
	StartErr               error
	CombinedOutputVal      []byte
	CombinedOutputErr      error
	OutputVal              []byte
	OutputErr              error
	SeparateOutputOut      []byte
	SeparateOutputErr      []byte
	SeparateOutputExitCode int
	StdinPipeVal           io.WriteCloser
	StdinPipeErr           error
	StderrPipeVal          io.ReadCloser
	StderrPipeErr          error
	StdoutPipeVal          io.ReadCloser
	StdoutPipeErr          error
	WaitErr                error
}

type mockCmd struct {
	Cmder
	mock          MockInstance
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
	methodsCalled []string
}

func (c *mockCmd) String() string {
	c.methodsCalled = append(c.methodsCalled, "String")
	return fmt.Sprintf("%s, %v", c.path, c.args)
}
func (c *mockCmd) Run() error {
	c.methodsCalled = append(c.methodsCalled, "Run")
	return c.mock.RunErr
}
func (c *mockCmd) Start() error {
	c.methodsCalled = append(c.methodsCalled, "Start")
	return c.mock.StartErr
}
func (c *mockCmd) CombinedOutput() ([]byte, error) {
	c.methodsCalled = append(c.methodsCalled, "CombinedOutput")
	return c.mock.CombinedOutputVal, c.mock.CombinedOutputErr
}
func (c *mockCmd) Output() ([]byte, error) {
	c.methodsCalled = append(c.methodsCalled, "Output")
	return c.mock.OutputVal, c.mock.OutputErr
}
func (c *mockCmd) SeparateOutput() ([]byte, []byte, int) {
	c.methodsCalled = append(c.methodsCalled, "SeparateOutput")
	return c.mock.SeparateOutputOut, c.mock.SeparateOutputErr, c.mock.SeparateOutputExitCode
}
func (c *mockCmd) StdinPipe() (io.WriteCloser, error) {
	c.methodsCalled = append(c.methodsCalled, "StdinPipe")
	return c.mock.StdinPipeVal, c.mock.StdinPipeErr
}
func (c *mockCmd) StderrPipe() (io.ReadCloser, error) {
	c.methodsCalled = append(c.methodsCalled, "StderrPipe")
	return c.mock.StderrPipeVal, c.mock.StderrPipeErr
}
func (c *mockCmd) StdoutPipe() (io.ReadCloser, error) {
	c.methodsCalled = append(c.methodsCalled, "StdoutPipe")
	return c.mock.StdoutPipeVal, c.mock.StdoutPipeErr
}
func (c *mockCmd) Wait() error {
	c.methodsCalled = append(c.methodsCalled, "Wait")
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
func (c *mockCmd) GetMethodsCalled() []string {
	return c.methodsCalled
}
