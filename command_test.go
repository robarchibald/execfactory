package command

import (
	"testing"
)

func TestCommand(t *testing.T) {
	SetExec()
	cmd := Command("ls", "-l")
	runner := cmd.(*shellCmd)
	if runner.cmd.Args[0] != "ls" || runner.cmd.Args[1] != "-l" {
		t.Error("expected valid command")
	}
}

func TestPipeCommands(t *testing.T) {
	SetExec()
	c1 := Command("ls")
	c2 := Command("grep", "command")
	output := PipeCommands(c1, c2)
	if output != "command.go\ncommand_test.go" {
		t.Error("expected valid command", output)
	}
}
