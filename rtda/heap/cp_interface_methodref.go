package heap

import "gojvm/classfile"

type ConstantInterfaceMethodref struct {
	ConstantMemberref
	method *Method
}

func newConstantInterfaceMethodref(refInfo *classfile.ConstantInterfaceMethodrefInfo) *ConstantInterfaceMethodref {
	ref := &ConstantInterfaceMethodref{}
	ref.copy(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *ConstantInterfaceMethodref) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.4
func (self *ConstantInterfaceMethodref) resolveInterfaceMethodRef() {

}
