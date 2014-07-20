// Copyright 2014 Will Maier <wcmaier@m.aier.us>. All rights reserved.
// See LICENSE for licensing information.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

// The decrypt function uses gpg to open and decrypt a file.
// Stdout is captured for later parsing, but stderr is allowed to pass through.
func decrypt(path string) ([]byte, error) {
	cmd := exec.Command("gpg", "-q", "--batch", "-d", path)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

// The parse function splits a byte array on whitespace.
func parse(data []byte) ([]string, error) {
	return strings.Fields(string(data)), nil
}

// The run function runs a command in an environment.
// Stdout and stderr are preserved.
func run(command []string, env []string) (error) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	return cmd.Run()
}

func usage() {
	self := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "usage: %s PATH COMMAND [ARGS...]\n\n", self)
	fmt.Fprint(os.Stderr, "Set environment variables defined in encrypted file PATH and run COMMAND.\n\n")
	fmt.Fprintln(os.Stderr, "Arguments:")
	fmt.Fprintln(os.Stderr, "  PATH     path to a gpg-encrypted file that can be read with eg `gpg -d PATH`")
	fmt.Fprintln(os.Stderr, "  COMMAND  command to be invoked in the context of the environment defined in PATH")
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	path := args[0]
	command := args[1:]
	env := []string{}

	out, err := decrypt(path)
	if err != nil {
		log.Fatal(err)
	}
	fields, err := parse(out)
	if err != nil {
		log.Fatal(err)
	}

	// Construct a new environment where values from the decrypted file
	// supercede values from our own environment.
	env = append(env, fields...)
	env = append(env, os.Environ()...)

	err = run(command, env)
	if err != nil {
		log.Fatal(err)
	}
}
