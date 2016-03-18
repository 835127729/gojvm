package main

import (
	_ "fmt"
	"gojvm/classfile"
	"gojvm/classpath"
	"gojvm/cmdline"
	"gojvm/options"
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
	_, classData, _ := cp.ReadClass("HelloWorld")
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic("failed to parse class file: " + "!" + err.Error())
	}
	cf.ToString()
	//heap.InitBootLoader(cp)
	/*
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
		mainThread := createMainThread(mainClassName, cmd.Args())
		interpreter.Loop(mainThread)
		interpreter.KeepAlive()
	*/
}

/*
func createMainThread(className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil)
	bootMethod := heap.BootstrapMethod()
	bootArgs := []interface{}{className, args}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
*/
