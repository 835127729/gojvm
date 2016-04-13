package main

import "fmt"
import "strings"
import "gojvm/classpath"
import "gojvm/cmdline"
import "gojvm/rtda/heap"

func main() {
	cmd := cmdline.ParseCmd()

	if cmd.VersionFlag() {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag() || cmd.Class() == "" {
		cmdline.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *cmdline.Cmd) {
	cp := classpath.Parse(cmd.XjreOption(), cmd.CpOption())
	classLoader := heap.NewClassLoader(cp, cmd.VerboseClassFlag())
	//heap.InitBootClassLoader(cp)

	className := strings.Replace(cmd.Class(), ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.VerboseClassFlag(), cmd.Args())
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.Class())
	}
}
