package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func version1(args []string) {
	var sep string
	if len(args) > 0 {
		for _, arg := range args {
			fmt.Print(sep + arg)
			sep = " "
		}
		fmt.Println()
	}
}

func version2(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func main() {
	startTime := time.Now()
	version2(os.Args[1:])
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))
}
