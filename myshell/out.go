package main

import (
	"fmt"
	"os"
)

type out struct {
}

func (*out) Write(b []byte) (n int, err error) {
	fmt.Println("Write:")
	os.Stdout.Write(b)
	return 0, nil
}
