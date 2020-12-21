package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	digits := [][]string{
		{
			"▓▓▓▓▓",
			"▓   ▓",
			"▓   ▓",
			"▓   ▓",
			"▓▓▓▓▓",
		},
		{
			"▓▓▓  ",
			"  ▓  ",
			"  ▓  ",
			"  ▓  ",
			"▓▓▓▓▓",
		},
		{
			"▓▓▓▓▓",
			"    ▓",
			"▓▓▓▓▓",
			"▓    ",
			"▓▓▓▓▓",
		},
		{
			"▓▓▓▓▓",
			"    ▓",
			" ▓▓▓▓",
			"    ▓",
			"▓▓▓▓▓",
		},
		{
			"▓   ▓",
			"▓   ▓",
			"▓▓▓▓▓",
			"    ▓",
			"    ▓",
		},
		{
			"▓▓▓▓▓",
			"▓    ",
			"▓▓▓▓▓",
			"    ▓",
			"▓▓▓▓▓",
		},
		{
			"▓▓▓▓▓",
			"▓    ",
			"▓▓▓▓▓",
			"▓   ▓",
			"▓▓▓▓▓",
		},
		{
			"▓▓▓▓▓",
			"    ▓",
			"    ▓",
			"    ▓",
			"    ▓",
		},
		{
			"▓▓▓▓▓",
			"▓   ▓",
			"▓▓▓▓▓",
			"▓   ▓",
			"▓▓▓▓▓",
		},
		{
			"▓▓▓▓▓",
			"▓   ▓",
			"▓▓▓▓▓",
			"    ▓",
			"▓▓▓▓▓",
		},
	}

	space := []string{
		"   ",
		" ▓ ",
		"   ",
		" ▓ ",
		"   ",
	}

	for {
		//clear the screen
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		//get the current time (hours, minutes, seconds)
		t := time.Now()
		h, m, s := t.Hour(), t.Minute(), t.Second()

		//print the current time
		fmt.Println(h, ":", m, ":", s)
		fmt.Println()

		//print the retro clock
		for i := 0; i < 5; i++ {
			fmt.Print(digits[h/10][i])
			fmt.Print(" ")
			fmt.Print(digits[h%10][i])
			fmt.Print(space[i])
			fmt.Print(digits[m/10][i])
			fmt.Print(" ")
			fmt.Print(digits[m%10][i])
			fmt.Print(space[i])
			fmt.Print(digits[s/10][i])
			fmt.Print(" ")
			fmt.Print(digits[s%10][i])
			fmt.Println()
		}

		//sleep for 1 second
		time.Sleep(time.Second)
	}

}
