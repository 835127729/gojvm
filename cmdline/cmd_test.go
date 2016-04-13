package cmdline

import (
	"fmt"
	"testing"
)

func TestClasspath(t *testing.T) {

	cmd := ParseCmd()

	if cmd.VersionFlag() {
		fmt.Println("version 0.0.1")
		return
	} else if cmd.HelpFlag() || cmd.Class() == "" {
		PrintUsage()
		return
	}
	fmt.Println(cmd.Class())
	fmt.Println(cmd)
}
