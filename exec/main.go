package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	command := Command(os.Args)
	err := Run(command, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Printf("command '%v` failed: '%v'\n", command, err)
		os.Exit(1)
	}
}

func Command(args []string) (command []string) {
	if len(args) < 2 {
		command = []string{"echo", "No command provided"}
		return
	}
	return args[1:]
}

func Run(command []string, i io.Reader, o, e io.Writer) (err error) {
	if len(command) < 1 {
		return errors.New("no command provided")
	}
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdin = i
	cmd.Stdout = o
	cmd.Stderr = e
	err = cmd.Run()
	return
}
