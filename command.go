package execfactory

import (
	"context"
	"io"
	"os"
	"os/exec"
	osexec "os/exec"
	"syscall"
)

type osCreator struct{}

func (c *osCreator) Command(name string, arg ...string) Cmder {
	return &osCmd{cmd: osexec.Command(name, arg...)}
}

func (c *osCreator) CommandContext(ctx context.Context, name string, arg ...string) Cmder {
	return &osCmd{cmd: osexec.CommandContext(ctx, name, arg...)}
}

type osCmd struct {
	cmd        *exec.Cmd
	workingDir string
}

func (c *osCmd) String() string {
	return c.cmd.String()
}
func (c *osCmd) Run() error {
	return convertError(c.cmd.Run())
}
func (c *osCmd) Start() error {
	return convertError(c.cmd.Start())
}
func (c *osCmd) CombinedOutput() ([]byte, error) {
	out, err := c.cmd.CombinedOutput()
	return out, convertError(err)
}
func (c *osCmd) Output() ([]byte, error) {
	out, err := c.cmd.Output()
	return out, convertError(err)
}
func (c *osCmd) StdinPipe() (io.WriteCloser, error) {
	return c.cmd.StdinPipe()
}
func (c *osCmd) StderrPipe() (io.ReadCloser, error) {
	return c.cmd.StderrPipe()
}
func (c *osCmd) StdoutPipe() (io.ReadCloser, error) {
	return c.cmd.StdoutPipe()
}
func (c *osCmd) Wait() error {
	return convertError(c.cmd.Wait())
}

// Sets

func (c *osCmd) SetPath(path string) {
	c.cmd.Path = path
}
func (c *osCmd) SetArgs(args []string) {
	c.cmd.Args = args
}
func (c *osCmd) SetEnv(env []string) {
	c.cmd.Env = env
}
func (c *osCmd) SetDir(path string) {
	c.cmd.Dir = path
}
func (c *osCmd) SetStdin(stdin io.Reader) {
	c.cmd.Stdin = stdin
}
func (c *osCmd) SetStdout(stdout io.Writer) {
	c.cmd.Stdout = stdout
}
func (c *osCmd) SetStderr(stderr io.Writer) {
	c.cmd.Stderr = stderr
}
func (c *osCmd) SetExtraFiles(files []*os.File) {
	c.cmd.ExtraFiles = files
}
func (c *osCmd) SetSysProcAttr(attr *syscall.SysProcAttr) {
	c.cmd.SysProcAttr = attr
}
func (c *osCmd) SetProcess(process *os.Process) {
	c.cmd.Process = process
}
func (c *osCmd) SetProcessState(processState *os.ProcessState) {
	c.cmd.ProcessState = processState
}

// Gets

func (c *osCmd) GetPath() string {
	return c.cmd.Path
}
func (c *osCmd) GetArgs() []string {
	return c.cmd.Args
}
func (c *osCmd) GetEnv() []string {
	return c.cmd.Env
}
func (c *osCmd) GetDir() string {
	return c.cmd.Dir
}
func (c *osCmd) GetStdin() io.Reader {
	return c.cmd.Stdin
}
func (c *osCmd) GetStdout() io.Writer {
	return c.cmd.Stdout
}
func (c *osCmd) GetStderr() io.Writer {
	return c.cmd.Stderr
}
func (c *osCmd) GetExtraFiles() []*os.File {
	return c.cmd.ExtraFiles
}
func (c *osCmd) GetSysProcAttr() *syscall.SysProcAttr {
	return c.cmd.SysProcAttr
}
func (c *osCmd) GetProcess() *os.Process {
	return c.cmd.Process
}
func (c *osCmd) GetProcessState() *os.ProcessState {
	return c.cmd.ProcessState
}
