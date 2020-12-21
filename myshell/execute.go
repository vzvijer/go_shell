package main

import (
	"io"
	"os"
	"os/exec"
	"os/user"
)

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Unknown"
	}
	return hostname
}

func getUsername() string {
	user, err := user.Current()
	if err != nil {
		return "Unknown"
	}
	return user.Username
}

func execCommand(command string, args []string, w io.Writer) {
	//output := &out{}
	if command == "exit" {
		os.Exit(0)
	} else {
		command := exec.Command(command, args...)
		command.Stdin = os.Stdin
		command.Stdout = w
		command.Stderr = w
		command.Run()
	}
}
