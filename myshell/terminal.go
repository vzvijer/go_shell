package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
)

func printPrompt(out io.Writer) {
	user, _ := user.Current()
	host, _ := os.Hostname()
	fmt.Fprintf(out, "%v@%v -> ", user.Username, host)
}

func startTerminalFrontend() {
	scanner := bufio.NewScanner(os.Stdin)
	//output := &out{}
	//output.Write([]byte("lsdjklfds"))
	for {
		printPrompt(os.Stdout)
		scanner.Scan()
		input := scanner.Text()
		inputFields := strings.Fields(input)
		command := inputFields[0]
		args := inputFields[1:]
		execCommand(command, args, os.Stdout)
	}
}
