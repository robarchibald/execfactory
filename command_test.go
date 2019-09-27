package execfactory

import (
	"testing"
)

func TestCommand(t *testing.T) {
	exec := NewOSCreator()
	cmd := exec.Command("ls", "-l")
	if cmd.GetArgs()[0] != "ls" || cmd.GetArgs()[1] != "-l" {
		t.Error("expected valid command")
	}
}

func TestPipeCommands(t *testing.T) {
	exec := NewOSCreator()
	c1 := exec.Command("ls")
	c2 := exec.Command("grep", "command")
	output := PipeCommands(c1, c2)
	if output != "command.go\ncommand_test.go" {
		t.Error("expected valid command", output)
	}
}
