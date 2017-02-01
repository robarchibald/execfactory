package command

import (
	"io"
	"reflect"
)

var mockType = reflect.TypeOf(MockShellCmd{})

type mockShell struct {
	cmd *MockShellCmd
}

func (c *mockShell) Command(name string, arg ...string) Cmder {
	cmd := reflect.New(mockType).Interface().(*MockShellCmd)
	cmd.Name = name
	cmd.Args = arg
	if c.cmd == nil {
		return cmd
	}
	cmd.RunErr = c.cmd.RunErr
	cmd.StartErr = c.cmd.StartErr
	cmd.CombinedOutputVal = c.cmd.CombinedOutputVal
	cmd.CombinedOutputErr = c.cmd.CombinedOutputErr
	cmd.OutputVal = c.cmd.OutputVal
	cmd.OutputErr = c.cmd.OutputErr
	cmd.workingDir = c.cmd.workingDir
	cmd.Stdin = c.cmd.Stdin
	cmd.Stdout = c.cmd.Stdout
	cmd.StderrPipeVal = c.cmd.StderrPipeVal
	cmd.StderrPipeErr = c.cmd.StderrPipeErr
	cmd.StdoutPipeVal = c.cmd.StdoutPipeVal
	cmd.StdoutPipeErr = c.cmd.StdoutPipeErr
	cmd.WaitErr = c.cmd.WaitErr
	return cmd
}

// MockShellCmd contains all the properties and methods to test a shell command
type MockShellCmd struct {
	Cmder
	Name              string
	Args              []string
	RunErr            error
	StartErr          error
	CombinedOutputVal []byte
	CombinedOutputErr error
	OutputVal         []byte
	OutputErr         error
	workingDir        string
	Stdin             io.ReadCloser
	Stdout            io.Writer
	StderrPipeVal     io.ReadCloser
	StderrPipeErr     error
	StdoutPipeVal     io.ReadCloser
	StdoutPipeErr     error
	WaitErr           error
}

// Run method
func (r *MockShellCmd) Run() error {
	return r.RunErr
}

// Start method
func (r *MockShellCmd) Start() error {
	return r.StartErr
}

// CombinedOutput method
func (r *MockShellCmd) CombinedOutput() ([]byte, error) {
	return r.CombinedOutputVal, r.CombinedOutputErr
}

// Output method
func (r *MockShellCmd) Output() ([]byte, error) {
	return r.OutputVal, r.OutputErr
}

// SetWorkingDir method
func (r *MockShellCmd) SetWorkingDir(path string) {
	r.workingDir = path
}

// GetWorkingDir method
func (r *MockShellCmd) GetWorkingDir() string {
	return r.workingDir
}

// SetStdin method
func (r *MockShellCmd) SetStdin(stdin io.ReadCloser) {
	r.Stdin = stdin
}

// SetStdout method
func (r *MockShellCmd) SetStdout(stdout io.Writer) {
	r.Stdout = stdout
}

// StderrPipe method
func (r *MockShellCmd) StderrPipe() (io.ReadCloser, error) {
	return r.StderrPipeVal, r.StderrPipeErr
}

// StdoutPipe method
func (r *MockShellCmd) StdoutPipe() (io.ReadCloser, error) {
	return r.StdoutPipeVal, r.StdoutPipeErr
}

// Wait method
func (r *MockShellCmd) Wait() error {
	return r.WaitErr
}
