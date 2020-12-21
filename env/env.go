package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const evLength int = 25
	envVars := map[string]string{}
	var envVarName string

	if len(os.Args) > 1 {
		envVarName = strings.ToUpper(os.Args[1])
	}

	for _, envVar := range os.Environ() {
		splited := strings.Split(envVar, "=")
		ev, evv := splited[0], splited[1]
		envVars[ev] = evv
	}

	if envVarName != "" {
		for ev, evv := range envVars {
			if strings.Index(ev, envVarName) == 0 {
				fmt.Println(fmt.Sprintf("%s = %s", ev, evv))
				return
			}
		}
		fmt.Println(fmt.Sprintf("Environment variable %s does not exist.", envVarName))
	} else {
		for ev, evv := range envVars {
			if len(ev) > evLength {
				ev = ev[:evLength-3] + "..."
			}
			fmtString := fmt.Sprintf("%%%ds", evLength)
			fmt.Println(fmt.Sprintf(fmtString, ev), evv)
		}
	}
}
