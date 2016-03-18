package heap

import (
	cf "gojvm/classfile"
)

type ConstantMemberref struct {
	className  string
	name       string
	descriptor string
}

func (self *ConstantMemberref) copy(refInfo *cf.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}
