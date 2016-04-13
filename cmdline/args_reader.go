package cmdline

type ArgsReader struct {
	args []string
}

func (self *ArgsReader) hasMoreOptions() bool {
	return len(self.args) > 0 && (self.args[0] == "classpath" || self.args[0][0] == '-')
}

func (self *ArgsReader) removeFirst() string {
	first := self.args[0]
	self.args = self.args[1:]
	return first
}
