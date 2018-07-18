package shell

import (
	"errors"
	"io"
	"os/exec"
)

type Shell struct {
	Proc   *exec.Cmd
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
}

func NewShell(command string) (Shell, error) {
	var err error

	shell := Shell{}
	shell.Proc = exec.Command(command)
	if shell.Stdin, err = shell.Proc.StdinPipe(); err != nil {
		return shell, errors.New("could not get a pipe to stdin")
	}
	if shell.Stdout, err = shell.Proc.StdoutPipe(); err != nil {
		return shell, errors.New("could not get a pipe to stdout")
	}

	if err = shell.Proc.Start(); err != nil {
		return shell, errors.New("could not start process")
	}

	return shell, nil
}

func (shell Shell) Close() {
	shell.Stdin.Close()
	shell.Stdout.Close()
}

func (shell Shell) Run(command string) (string, error) {
	// TODO handle case when rlen is > 1024
	shell.Stdin.Write([]byte(command + "\n"))
	var buf []byte
	buf = make([]byte, 1024)
	rlen, err := shell.Stdout.Read(buf)
	if err != nil {
		return "", errors.New("Error reading from stdout:\n" + err.Error())
	}
	out := string(buf[:rlen])
	return out, nil
}
