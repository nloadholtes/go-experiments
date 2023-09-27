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
	return CommandFromStringSlice(input_values)
}

func CommandFromStringSlice(input []string) (*exec.Cmd, error) {
	var command string = input[0]
	var values []string = input[1:]
	return exec.Command(command, values...), nil
}

func main() {
	//fmt.Println(len(os.Args), os.Args)
	cmd_args := os.Args[1:]
	cmd, _ := CommandFromStringSlice(cmd_args)
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
