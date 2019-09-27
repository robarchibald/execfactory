package execfactory

import (
	"testing"
)

func TestCommand(t *testing.T) {
	cmd := OS.Command("ls", "-l")
	if cmd.GetArgs()[0] != "ls" || cmd.GetArgs()[1] != "-l" {
		t.Error("expected valid command")
	}
}

func TestPipeCommands(t *testing.T) {
	c1 := OS.Command("ls")
	c2 := OS.Command("grep", "command")
	output := PipeCommands(c1, c2)
	if output != "command.go\ncommand_test.go" {
		t.Error("expected valid command", output)
	}
}
