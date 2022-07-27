package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"reflect"
	"testing"
)

func ExampleRun() {
	args := []string{"echo", "hello world"}
	Run(args, os.Stdin, os.Stdout, os.Stderr)
	args = []string{"ls"}
	Run(args, os.Stdin, os.Stdout, os.Stderr)
	// Output:
	// hello world
	// main.go
	// main_test.go
}

func TestCommand(t *testing.T) {
	for _, c := range casesCommand() {
		t.Run(c.name, func(t *testing.T) {
			command := Command(c.args)
			if !reflect.DeepEqual(command, c.expect) {
				t.Errorf("got '%v', expected '%v'", command, c.expect)
			}
		})
	}
}

type testCommand struct {
	name   string
	args   []string
	expect []string
}

func casesCommand() []testCommand {
	return []testCommand{
		{"no command", []string{"main"}, []string{"echo", "No command provided"}},
		{"simple", []string{"main", "ls"}, []string{"ls"}},
		{"longer", []string{"main", "echo", "one", "two"}, []string{"echo", "one", "two"}},
	}
}

func TestRun(t *testing.T) {
	for _, c := range casesRun() {
		t.Run(c.name, func(t *testing.T) {
			err := Run(c.command, c.in, c.out, c.errOut)
			if err != nil && c.expect != nil {
				if !(len(err.Error()) > 0 && len(c.expect.Error()) > 0) {
					t.Errorf("got '%v', expected '%v'", err, c.expect)
				}
			}
		})
	}
}

type testRun struct {
	name    string
	command []string
	in      io.Reader
	out     io.Writer
	errOut  io.Writer
	expect  error
}

func casesRun() []testRun {
	return []testRun{
		{"success", []string{"ls", "-l"}, newReader(), newWriter(), newWriter(), nil},
		{"unknown command", []string{"notacommand", "-l"}, newReader(), newWriter(), newWriter(), errors.New("error")},
		{"empty command", []string{}, newReader(), newWriter(), newWriter(), errors.New("error")},
	}
}

func newWriter() *bytes.Buffer {
	return &bytes.Buffer{}
}

func newReader() (reader io.Reader) {
	return
}
