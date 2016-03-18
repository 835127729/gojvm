package heap

import "gojvm/classfile"

type ConstantInterfaceMethodref struct {
	ConstantMemberref
	method *Method
}

func newConstantInterfaceMethodref(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *ConstantInterfaceMethodref {
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
	d := self.cp.class
	c := self.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

// todo
func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
