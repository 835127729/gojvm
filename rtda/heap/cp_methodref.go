package heap

import (
	"gojvm/classfile"
	_ "gojvm/jutil"
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

// jvms8 5.4.3.3
func (self *ConstantMethodref) ResolveMethod() *Method {
	c := bootLoader.LoadClass(self.className)
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	self.method = method
	return method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
