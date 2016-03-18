package heap

import (
	cf "gojvm/classfile"
)

type ConstantClassref struct {
	name  string
	class *Class
}

func newConstantClassref(classInfo *cf.ConstantClassInfo) *ConstantClassref {
	return &ConstantClassref{
		name: classInfo.Name(),
	}
}

func (self *ConstantClassref) Class() *Class {
	if self.class == nil {
		self.resolve()
	}
	return self.class
}

// todo
func (self *ConstantClassref) resolve() {
	// load class
	self.class = bootLoader.LoadClass(self.name)
}
