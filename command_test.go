package execfactory

import (
	"os"
	"os/exec"
	"reflect"
	"strings"
	"syscall"
	"testing"
)

func newCmd() *exec.Cmd {
	return exec.Command("ls", "command.go")
}

func TestCommand(t *testing.T) {
	factory := NewOSCreator()
	cmd := factory.Command("ls", "-l")
	if cmd.GetArgs()[0] != "ls" || cmd.GetArgs()[1] != "-l" {
		t.Error("expected valid command")
	}
}

func TestPipeCommands(t *testing.T) {
	factory := NewOSCreator()
	c1 := factory.Command("ls")
	c2 := factory.Command("grep", "command")
	output := PipeCommands(c1, c2)
	if output != "command.go\ncommand_test.go" {
		t.Error("expected valid command", output)
	}
}

func TestRealString(t *testing.T) {
	cmd := &osCmd{&exec.Cmd{Path: "hello", Args: []string{"there"}}}
	if v := cmd.String(); v != "hello" {
		t.Error("Expected string to work", v)
	}
}

func TestRealRun(t *testing.T) {
	method := "Run"
	cmd := &osCmd{newCmd()}
	if err := cmd.Run(); err != nil {
		t.Errorf("Expected %s to work: %v", method, err)
	}
}

func TestRealStart(t *testing.T) {
	method := "Start"
	cmd := &osCmd{newCmd()}
	if err := cmd.Start(); err != nil {
		t.Errorf("Expected %s to work: %v", method, err)
	}
}

func TestRealCombinedOutput(t *testing.T) {
	method := "CombinedOutput"
	cmd := &osCmd{newCmd()}
	if v, err := cmd.CombinedOutput(); err != nil || string(v) != "command.go\n" {
		t.Errorf("Expected %s to work with output %s: %v", method, string(v), err)
	}
}

func TestRealOutput(t *testing.T) {
	method := "Output"
	cmd := &osCmd{newCmd()}
	if v, err := cmd.Output(); err != nil || string(v) != "command.go\n" {
		t.Errorf("Expected %s to work with output %s: %v", method, string(v), err)
	}
}

func TestRealStdinPipe(t *testing.T) {
	method := "StdinPipe"
	cmd := &osCmd{newCmd()}
	if w, err := cmd.StdinPipe(); err != nil || w == nil {
		t.Errorf("Expected %s to work: %v", method, err)
	}
}

func TestRealStderrPipe(t *testing.T) {
	method := "StderrPipe"
	cmd := &osCmd{newCmd()}
	if r, err := cmd.StderrPipe(); err != nil || r == nil {
		t.Errorf("Expected %s to work: %v", method, err)
	}
}

func TestRealStdoutPipe(t *testing.T) {
	method := "StdoutPipe"
	cmd := &osCmd{newCmd()}
	if r, err := cmd.StdoutPipe(); err != nil || r == nil {
		t.Errorf("Expected %s to work: %v", method, err)
	}
}

func TestRealWait(t *testing.T) {
	method := "Wait"
	cmd := &osCmd{newCmd()}
	if err := cmd.Wait(); err == nil || err.Error() != "exec: not started" {
		t.Errorf("Expected %s to work: %v", method, err)
	}
}

func TestRealGetSetPath(t *testing.T) {
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetPath("path")
	if v := cmd.GetPath(); v != "path" {
		t.Error("Expected valid path", v)
	}
}

func TestRealGetSetArgs(t *testing.T) {
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetPath("path")
	if v := cmd.GetPath(); v != "path" {
		t.Error("Expected valid path", v)
	}
}

func TestRealGetSetEnv(t *testing.T) {
	env := []string{"1", "2"}
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetEnv(env)
	if v := cmd.GetEnv(); !reflect.DeepEqual(v, env) {
		t.Error("Expected valid env", v)
	}
}
func TestRealGetSetDir(t *testing.T) {
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetDir("dir")
	if v := cmd.GetDir(); v != "dir" {
		t.Error("Expected valid dir", v)
	}
}
func TestRealGetSetStdin(t *testing.T) {
	cmd := &osCmd{&exec.Cmd{}}
	r := strings.NewReader("string")
	cmd.SetStdin(r)
	if v := cmd.GetStdin(); v != r {
		t.Error("Expected valid stdin", v)
	}
}
func TestRealGetSetStdout(t *testing.T) {
	buf := &strings.Builder{}
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetStdout(buf)
	if v := cmd.GetStdout(); v != buf {
		t.Error("Expected valid stdout", buf)
	}
}
func TestRealGetSetStderr(t *testing.T) {
	buf := &strings.Builder{}
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetStderr(buf)
	if v := cmd.GetStderr(); v != buf {
		t.Error("Expected valid stderr", v)
	}
}
func TestRealGetSetExtraFiles(t *testing.T) {
	files := []*os.File{{}}
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetExtraFiles(files)
	if v := cmd.GetExtraFiles(); !reflect.DeepEqual(v, files) {
		t.Error("Expected valid extra files", v)
	}
}
func TestRealGetSetProcAttr(t *testing.T) {
	procAttr := &syscall.SysProcAttr{}
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetSysProcAttr(procAttr)
	if v := cmd.GetSysProcAttr(); v != procAttr {
		t.Error("Expected valid sysProcAttr", v)
	}
}
func TestRealGetSetProcess(t *testing.T) {
	proc := &os.Process{}
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetProcess(proc)
	if v := cmd.GetProcess(); v != proc {
		t.Error("Expected valid process", v)
	}
}
func TestRealGetSetProcessState(t *testing.T) {
	ps := &os.ProcessState{}
	cmd := &osCmd{&exec.Cmd{}}
	cmd.SetProcessState(ps)
	if v := cmd.GetProcessState(); v != ps {
		t.Error("Expected valid processState", v)
	}
}
