package cmdline

import "fmt"
import "os"

// java [-options] class [args...]
type Cmd struct {
	helpFlag         bool
	versionFlag      bool
	verboseClassFlag bool
	verboseInstFlag  bool
	cpOption         string
	xjreOption       string
	class            string
	args             []string
}

func ParseCmd() *Cmd {
	return parseCmd()
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	margs := os.Args[1:]
	argsReader := ArgsReader{
		args: margs,
	}
	for argsReader.hasMoreOptions() {
		optionName := argsReader.removeFirst()
		switch optionName {
		case "-help", "?":
			cmd.helpFlag = true
		case "-version":
			cmd.versionFlag = true
		case "-verbose", "-verbose:class":
			cmd.verboseClassFlag = true
		case "-verbose:inst":
			cmd.verboseInstFlag = true
		case "-cp", "-classpath", "classpath":
			cmd.cpOption = argsReader.removeFirst()
		case "-Xjre":
			cmd.xjreOption = argsReader.removeFirst()
		}
	}
	cmd.class = argsReader.removeFirst()
	cmd.args = argsReader.args
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}

//getter
func (self *Cmd) HelpFlag() bool {
	return self.helpFlag
}

func (self *Cmd) VersionFlag() bool {
	return self.versionFlag
}

func (self *Cmd) VerboseClassFlag() bool {
	return self.verboseClassFlag
}

func (self *Cmd) VerboseInstFlag() bool {
	return self.verboseInstFlag
}

func (self *Cmd) CpOption() string {
	return self.cpOption
}

func (self *Cmd) XjreOption() string {
	return self.xjreOption
}

func (self *Cmd) Class() string {
	return self.class
}

func (self *Cmd) Args() []string {
	return self.args
}
