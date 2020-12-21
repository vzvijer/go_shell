package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func checkError(err error, ) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	const help = `Usage:    grep [OPTION]... PATTERN [FILE]...
	Search for PATTERN in each FILE.
	Example: grep -i 'hello world' menu.h main.c`

	files := []os.File{}
	if len(os.Args) < 2 {
		fmt.Print(help)
	} else if len(os.Args) < 3 {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		files, err := ioutil.ReadDir(wd)
		if err
		for fn := range 
	}
	pattern := os.Args[1]
	file := os.Args[2]
}
