package main

import (
	"testing"
)

func TestCommandFromString(t *testing.T) {
	cmd_obj, err := CommandFromString("ls -la")
	if err != nil {
		t.Fatal("CommandFromString errored out: ", err)
	}
	// I really do not like this, I did not specify that path yet it derives it?
	if cmd_obj.Path != "/bin/ls" {
		t.Fatal("cmd_obj is not ls. Saw: ", cmd_obj.Path)
	}
	if cmd_obj.Args[1] != "-la" {
		t.Fatal("wrong arguments for ls seen: ", cmd_obj.Args)
	}

}
