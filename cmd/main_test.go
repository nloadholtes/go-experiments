package main

import (
	"testing"
)

func TestCommandFromString(t *testing.T) {
	cmd_obj, err := CommandFromString("ls -la")
	if err != nil {
		t.Fatal("CommandFromString errored out: ", err)
	}
	if cmd_obj.Path != "ls" {
		t.Fatal("cmd_obj is not ls. Saw: ", cmd_obj.Path)
	}
	if cmd_obj.Args[0] != "-la" {
		t.Fatal("wrong arguments for ls seen: ", cmd_obj.Args)
	}

}
