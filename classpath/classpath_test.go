package classpath

import (
	"fmt"
	"gojvm/cmdline"
	"testing"
)

func TestClasspath(t *testing.T) {

	cmd := cmdline.ParseCmd()

	if cmd.VersionFlag() {
		fmt.Println("version 0.0.1")
		return
	} else if cmd.HelpFlag() || cmd.Class() == "" {
		cmdline.PrintUsage()
		return
	}
	cp := Parse(cmd.XjreOption(), cmd.CpOption())
	fmt.Println(cp.String())
}
