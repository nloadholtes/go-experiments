package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func CommandFromString(input string) (*exec.Cmd, error) {
	input_values := strings.Split(input, " ")
	var command string = input_values[0]
	var values string = input_values[1]

	return exec.Command(command, values), nil
}

func main() {
	fmt.Println(len(os.Args), os.Args)
	cmd_args := os.Args[1:]
	cmd := exec.Command(cmd_args[0], cmd_args...)
	cmd.Env = append(os.Environ(),
		"FOO=duplicate_value", // ignored
		"FOO=actual_value",    // this value is used
	)

	var out strings.Builder
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OUTPUT:\n %s", out.String())
}
