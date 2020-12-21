package main

import (
	"os"
)

const (
	terminal = iota
	browser
)

func getFrontend() int {
	frontend := terminal
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-b":
			frontend = browser
		default:
			frontend = terminal
		}
	}
	return frontend
}

func main() {
	frontend := getFrontend()

	if frontend == terminal {
		startTerminalFrontend()
	} else if frontend == browser {
		startBrowserFrontend()
	}
}
