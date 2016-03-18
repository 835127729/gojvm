package options

import (
	"os"
	"path/filepath"

	"gojvm/cmdline"
)

var (
	VerboseClass    bool
	ThreadStackSize uint
	AbsJavaHome     string // /path/to/jre
	AbsJreLib       string // /path/to/jre/lib
)

func InitOptions(cmdOptions *cmdline.Options) {
	VerboseClass = cmdOptions.VerboseClass()
	ThreadStackSize = uint(cmdOptions.Xss)
	initJavaHome(cmdOptions.XuseJavaHome)
}

func initJavaHome(useOsEnv bool) {
	//todo
	//jh := "./jre"
	jh := "C:/Program Files/Java/jdk1.8.0_60/jre"
	if useOsEnv {
		jh = os.Getenv("JAVA_HOME")
		if jh == "" {
			panic("$JAVA_HOME not set!")
		}
	}

	if absJh, err := filepath.Abs(jh); err == nil {
		AbsJavaHome = absJh
		AbsJreLib = filepath.Join(absJh, "lib")
	} else {
		panic(err)
	}
}
