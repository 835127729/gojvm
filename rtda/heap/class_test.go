package heap

import (
	"gojvm/classpath"
	"gojvm/cmdline"
	"gojvm/jutil"
	"gojvm/options"
	"os"
	"runtime/pprof"
	"testing"
)

func Test_Class(t *testing.T) {
	cmd, err := cmdline.ParseCommand(os.Args[1:])
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
	InitBootLoader(cp)
	mainClassName := jutil.ReplaceAll(cmd.Class(), ".", "/")
	testClass := bootLoader.LoadClass(mainClassName)
	testClass.ToString()
}
