package command

import (
	"io"
	"os/exec"
)

type shell struct{}

func (c *shell) Command(name string, arg ...string) Cmder {
	return &shellCmd{cmd: exec.Command(name, arg...)}
}

type shellCmd struct {
	cmd        *exec.Cmd
	workingDir string
}

func (r *shellCmd) Run() error {
	return r.cmd.Run()
}

func (r *shellCmd) Start() error {
	return r.cmd.Start()
}

func (r *shellCmd) CombinedOutput() ([]byte, error) {
	return r.cmd.CombinedOutput()
}

func (r *shellCmd) Output() ([]byte, error) {
	return r.cmd.Output()
}

func (r *shellCmd) SetWorkingDir(path string) {
	r.cmd.Dir = path
}

func (r *shellCmd) SetStdin(stdin io.ReadCloser) {
	r.cmd.Stdin = stdin
}

func (r *shellCmd) SetStdout(stdout io.Writer) {
	r.cmd.Stdout = stdout
}

func (r *shellCmd) StderrPipe() (io.ReadCloser, error) {
	return r.cmd.StderrPipe()
}

func (r *shellCmd) StdoutPipe() (io.ReadCloser, error) {
	return r.cmd.StdoutPipe()
}

func (r *shellCmd) Wait() error {
	return r.cmd.Wait()
}
