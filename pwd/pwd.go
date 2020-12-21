package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	optionL := false
	optionK := false

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(os.Args) > 1 {
		options := os.Args[1]
		if options[0] == '-' {
			options = options[1:]

			for _, option := range options {
				if option == 'L' {
					optionL = true
				} else if option == 'K' {
					optionK = true
				} else {
					fmt.Println(fmt.Sprintf("pwd: -%c: invalid option", option))
					fmt.Println("pwd: usage: pwd [-LP]")
					return
				}
			}
		}
	}
	if optionL {
		//todo
	}
	if optionK {
		//todo
	}
	fmt.Println(wd)
}
