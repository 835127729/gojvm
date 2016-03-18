package main

import (
	_ "fmt"
	"gojvm/classfile"
	"gojvm/classpath"
	"gojvm/cmdline"
	"gojvm/options"
	"os"
	"runtime/pprof"
	"testing"
)

func Test_Classfile(t *testing.T) {
	cmd, err := cmdline.ParseCommand(os.Args)
	if err != nil {
		cmdline.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *cmdline.Command) {
	Xcpuprofile := cmd.Options().Xcpuprofile
	if Xcpuprofile != "" {
		f, err := os.Create(Xcpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	options.InitOptions(cmd.Options())

	cp := classpath.Parse(cmd.Options().Classpath())
	_, classData, _ := cp.ReadClass("HelloWorld")
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic("failed to parse class file: " + "!" + err.Error())
	}
	cf.ToString()
}
