package exec

import (
	"testing"
)

func TestCommand(t *testing.T) {
	cmd := Command("ls", "-l")
	if cmd.GetArgs()[0] != "ls" || cmd.GetArgs()[1] != "-l" {
		t.Error("expected valid command")
	}
}

func TestPipeCommands(t *testing.T) {
	c1 := Command("ls")
	c2 := Command("grep", "command")
	output := PipeCommands(c1, c2)
	if output != "command.go\ncommand_test.go" {
		t.Error("expected valid command", output)
	}
}
