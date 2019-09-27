package execfactory

import (
	"os"
	osexec "os/exec"
	"time"
)

// An ExitError reports an unsuccessful exit by a command.
type ExitError struct {
	*os.ProcessState

	// Stderr holds a subset of the standard error output from the
	// Cmd.Output method if standard error was not otherwise being
	// collected.
	//
	// If the error output is long, Stderr may contain only a prefix
	// and suffix of the output, with the middle replaced with
	// text about the number of omitted bytes.
	//
	// Stderr is provided for debugging, for inclusion in error messages.
	// Users with other needs should redirect Cmd.Stderr as needed.
	Stderr []byte
	err    *osexec.ExitError
}

func convertError(err error) error {
	if ee, ok := err.(*osexec.ExitError); ok {
		return &ExitError{ee.ProcessState, ee.Stderr, ee}
	}
	return err
}

func (e *ExitError) Error() string {
	return e.err.Error()
}

// ExitCode returns the exit code of the exited process, or -1
// if the process hasn't exited or was terminated by a signal.
func (e *ExitError) ExitCode() int {
	return e.err.ExitCode()
}

// Exited reports whether the program has exited.
func (e *ExitError) Exited() bool {
	return e.err.Exited()
}

// Pid returns the process id of the exited process.
func (e *ExitError) Pid() int {
	return e.err.Pid()
}
func (e *ExitError) String() string {
	return e.err.String()
}

// Success reports whether the program exited successfully,
// such as with exit status 0 on Unix.
func (e *ExitError) Success() bool {
	return e.err.Success()
}

// Sys returns system-dependent exit information about
// the process. Convert it to the appropriate underlying
// type, such as syscall.WaitStatus on Unix, to access its contents.
func (e *ExitError) Sys() interface{} {
	return e.err.Sys()
}

// SysUsage returns system-dependent resource usage information about
// the exited process. Convert it to the appropriate underlying
// type, such as *syscall.Rusage on Unix, to access its contents.
// (On Unix, *syscall.Rusage matches struct rusage as defined in the
// getrusage(2) manual page.)
func (e *ExitError) SysUsage() interface{} {
	return e.err.SysUsage()
}

// SystemTime returns the system CPU time of the exited process and its children.
func (e *ExitError) SystemTime() time.Duration {
	return e.err.SystemTime()
}

// UserTime returns the user CPU time of the exited process and its children.
func (e *ExitError) UserTime() time.Duration {
	return e.err.UserTime()
}
