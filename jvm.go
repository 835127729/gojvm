package main

import (
	"fmt"
	"gojvm/classpath"
	"gojvm/cmdline"
	"gojvm/jutil"
	"gojvm/options"
	"gojvm/rtda"
	"gojvm/rtda/heap"
	"os"
	"runtime/pprof"
)

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
	heap.InitBootLoader(cp)
	mainClassName := jutil.ReplaceAll(cmd.Class(), ".", "/")
	mainClass := heap.GetBootLoader().LoadClass(mainClassName)
	mainThread := createMainThread(mainClassName, cmd.Args())
	mainMethod := mainClass.GetMainMethod()
	fmt.Println(mainMethod.Name())
	mainThread.NewFrame(mainMethod)
}

func createMainThread(className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread()
	return mainThread
}
