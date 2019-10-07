package execfactory

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"syscall"
	"testing"
)

func TestNewMockCmd(t *testing.T) {
	creator := &mockCreator{}
	cmd := creator.newMockCmd("go", "build")
	if args := cmd.GetArgs(); cmd.GetPath() != "go" || len(args) != 1 || args[0] != "build" {
		t.Error("Expected correct mock command", cmd)
	}

	// with MockInstance
	mock := MockInstance{RunErr: errors.New("test")}
	creator = &mockCreator{instances: []MockInstance{mock}}
	cmd = creator.newMockCmd("go", "build")
	if mCmd := cmd; mCmd.mock.StartErr != mock.StartErr {
		t.Error("Expected correct mock command", cmd)
	}

	// used up MockInstances, so now nil
	cmd = creator.newMockCmd("go", "build")
	if mCmd := cmd; mCmd.mock.StartErr != nil {
		t.Error("Expected correct mock command", cmd)
	}
}

func TestMockString(t *testing.T) {
	cmd := &mockCmd{path: "hello", args: []string{"there"}}
	if v := cmd.String(); v != "hello, [there]" {
		t.Error("Expected string to work", v)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != "String" {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockRun(t *testing.T) {
	method := "Run"
	expectedErr := fmt.Errorf("%s error", method)
	cmd := &mockCmd{mock: MockInstance{RunErr: expectedErr}}
	if err := cmd.Run(); err != expectedErr {
		t.Errorf("Expected %s to work: %v", method, err)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != method {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockStart(t *testing.T) {
	method := "Start"
	expectedErr := fmt.Errorf("%s error", method)
	cmd := &mockCmd{mock: MockInstance{StartErr: expectedErr}}
	if err := cmd.Start(); err != expectedErr {
		t.Errorf("Expected %s to work: %v", method, err)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != method {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockCombinedOutput(t *testing.T) {
	method := "CombinedOutput"
	expectedErr := fmt.Errorf("%s error", method)
	cmd := &mockCmd{mock: MockInstance{CombinedOutputErr: expectedErr, CombinedOutputVal: []byte(method)}}
	if v, err := cmd.CombinedOutput(); err != expectedErr || string(v) != "CombinedOutput" {
		t.Errorf("Expected %s to work: %v", method, err)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != method {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockOutput(t *testing.T) {
	method := "Output"
	expectedErr := fmt.Errorf("%s error", method)
	cmd := &mockCmd{mock: MockInstance{OutputErr: expectedErr, OutputVal: []byte(method)}}
	if v, err := cmd.Output(); err != expectedErr || string(v) != "Output" {
		t.Errorf("Expected %s to work: %v", method, err)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != method {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockStdinPipe(t *testing.T) {
	var buf strings.Builder
	expectedW := newNopWriteCloser(&buf)
	method := "StdinPipe"
	expectedErr := fmt.Errorf("%s error", method)
	cmd := &mockCmd{mock: MockInstance{StdinPipeErr: expectedErr, StdinPipeVal: expectedW}}
	if w, err := cmd.StdinPipe(); err != expectedErr || w != expectedW {
		t.Errorf("Expected %s to work: %v", method, err)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != method {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockStderrPipe(t *testing.T) {
	expectedR := ioutil.NopCloser(strings.NewReader("hello"))
	method := "StderrPipe"
	expectedErr := fmt.Errorf("%s error", method)
	cmd := &mockCmd{mock: MockInstance{StderrPipeErr: expectedErr, StderrPipeVal: expectedR}}
	if r, err := cmd.StderrPipe(); err != expectedErr || r != expectedR {
		t.Errorf("Expected %s to work: %v", method, err)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != method {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockStdoutPipe(t *testing.T) {
	expectedR := ioutil.NopCloser(strings.NewReader("hello"))
	method := "StdoutPipe"
	expectedErr := fmt.Errorf("%s error", method)
	cmd := &mockCmd{mock: MockInstance{StdoutPipeErr: expectedErr, StdoutPipeVal: expectedR}}
	if r, err := cmd.StdoutPipe(); err != expectedErr || r != expectedR {
		t.Errorf("Expected %s to work: %v", method, err)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != method {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockWait(t *testing.T) {
	method := "Wait"
	expectedErr := fmt.Errorf("%s error", method)
	cmd := &mockCmd{mock: MockInstance{WaitErr: expectedErr}}
	if err := cmd.Wait(); err != expectedErr {
		t.Errorf("Expected %s to work: %v", method, err)
	}
	if m := cmd.methodsCalled; len(m) != 1 || m[0] != method {
		t.Error("Expected methodsCalled to have been filled", m)
	}
}

func TestMockGetSetPath(t *testing.T) {
	cmd := &mockCmd{}
	cmd.SetPath("path")
	if v := cmd.GetPath(); v != "path" {
		t.Error("Expected valid path", v)
	}
}

func TestMockGetSetArgs(t *testing.T) {
	cmd := &mockCmd{}
	cmd.SetPath("path")
	if v := cmd.GetPath(); v != "path" {
		t.Error("Expected valid path", v)
	}
}

func TestMockGetSetEnv(t *testing.T) {
	env := []string{"1", "2"}
	cmd := &mockCmd{}
	cmd.SetEnv(env)
	if v := cmd.GetEnv(); !reflect.DeepEqual(v, env) {
		t.Error("Expected valid env", v)
	}
}
func TestMockGetSetDir(t *testing.T) {
	cmd := &mockCmd{}
	cmd.SetDir("dir")
	if v := cmd.GetDir(); v != "dir" {
		t.Error("Expected valid dir", v)
	}
}
func TestMockGetSetStdin(t *testing.T) {
	cmd := &mockCmd{}
	r := strings.NewReader("string")
	cmd.SetStdin(r)
	if v := cmd.GetStdin(); v != r {
		t.Error("Expected valid stdin", v)
	}
}
func TestMockGetSetStdout(t *testing.T) {
	buf := &strings.Builder{}
	cmd := &mockCmd{}
	cmd.SetStdout(buf)
	if v := cmd.GetStdout(); v != buf {
		t.Error("Expected valid stdout", buf)
	}
}
func TestMockGetSetStderr(t *testing.T) {
	buf := &strings.Builder{}
	cmd := &mockCmd{}
	cmd.SetStderr(buf)
	if v := cmd.GetStderr(); v != buf {
		t.Error("Expected valid stderr", v)
	}
}
func TestMockGetSetExtraFiles(t *testing.T) {
	files := []*os.File{{}}
	cmd := &mockCmd{}
	cmd.SetExtraFiles(files)
	if v := cmd.GetExtraFiles(); !reflect.DeepEqual(v, files) {
		t.Error("Expected valid extra files", v)
	}
}
func TestMockGetSetProcAttr(t *testing.T) {
	procAttr := &syscall.SysProcAttr{}
	cmd := &mockCmd{}
	cmd.SetSysProcAttr(procAttr)
	if v := cmd.GetSysProcAttr(); v != procAttr {
		t.Error("Expected valid sysProcAttr", v)
	}
}
func TestMockGetSetProcess(t *testing.T) {
	proc := &os.Process{}
	cmd := &mockCmd{}
	cmd.SetProcess(proc)
	if v := cmd.GetProcess(); v != proc {
		t.Error("Expected valid process", v)
	}
}
func TestMockGetSetProcessState(t *testing.T) {
	ps := &os.ProcessState{}
	cmd := &mockCmd{}
	cmd.SetProcessState(ps)
	if v := cmd.GetProcessState(); v != ps {
		t.Error("Expected valid processState", v)
	}
}
