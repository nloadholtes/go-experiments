package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"fmt"
)

func main() {
	cmd := exec.Command("ls", "-la")
	cmd.Env = append(os.Environ(),
		"FOO=duplicate_value", // ignored
		"FOO=actual_value",    // this value is used
	)

	var out strings.Builder
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OUTPUT:\n", out.String())
}
