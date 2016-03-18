package heap

import (
	"gojvm/classfile"
)

type ConstantMethodref struct {
	ConstantMemberref
	method *Method
}

func newConstantMethodref(refInfo *classfile.ConstantMethodrefInfo) *ConstantMethodref {
	ref := &ConstantMethodref{}
	ref.copy(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *ConstantMethodref) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
func (self *ConstantMethodref) resolveMethodRef() {

}
