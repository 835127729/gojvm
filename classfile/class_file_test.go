package cmdline

import (
	"os"
	"testing"
)

func Test_ParseOptions(t *testing.T) {
	osArgs := os.Args
	t.Log("原命令:", osArgs)
	argReader := &ArgReader{osArgs[2:]}
	cmd := &Command{
		options: parseOptions(argReader),
		class:   argReader.removeFirst(),
		args:    argReader.args,
	}
	t.Log("解析后:")
	t.Log("classpath:", cmd.options.Classpath())
	t.Log("VerboseClass:", cmd.options.VerboseClass())
	t.Log("Xcpuprofile:", cmd.options.Xcpuprofile)
	t.Log("Xss:", cmd.options.Xss)
	t.Log("XuseJavaHome:", cmd.options.XuseJavaHome)
}
