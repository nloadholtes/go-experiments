package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func CommandFromString(input string, alias_map map[string]string) (*exec.Cmd, error) {
	input_values := strings.Split(input, " ")
	return CommandFromStringSlice(input_values, alias_map)
}

func CommandFromStringSlice(input []string, alias_map map[string]string) (*exec.Cmd, error) {
	var command string = input[0]
	var values []string = input[1:]
	alias_command, ok := alias_map[command]
	if ok {
		alias_command_slice := strings.Split(alias_command, " ")
		command = alias_command_slice[0]
		values = append(alias_command_slice[1:], input[1:]...)
	}
	return exec.Command(command, values...), nil
}

func MapAliases() (map[string]string, error) {
	var output = make(map[string]string)
	cmd := exec.Command("/usr/bin/zsh", "-ic", "alias")
	cmd_output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	cmd_lines := strings.Split(string(cmd_output), "\n")
	for _, line := range cmd_lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(strings.ReplaceAll(line, "'", ""), "=", 2)
		output[parts[0]] = parts[1]
	}

	return output, nil
}

func main() {
	cmd_args := os.Args[1:]
	alias_map, _ := MapAliases()
	cmd, _ := CommandFromStringSlice(cmd_args, alias_map)
	// I know I needed env, but I don't remember why the FOO
	// was added/needed.
	cmd.Env = append(os.Environ(),
		"FOO=duplicate_value", // ignored
		"FOO=actual_value",    // this value is used
	)

	var out strings.Builder
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}
