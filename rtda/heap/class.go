package heap

import (
	"gojvm/classfile"
)

type Class struct {
	AccessFlags
	name           string
	constantPool   *ConstantPool
	superClassName string
	interfaceNames []string
	fields         []*Field
	methods        []*Method
	jClass         *Object // java.lang.Class instance
	superClass     *Class
	interfaces     []*Class
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
	return class
}

// getters
func (self *Class) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *Class) Name() string {
	return self.name
}
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) SourceFile() string {
	return self.sourceFile
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) Interfaces() []*Class {
	return self.interfaces
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}
func (self *Class) InitStarted() bool {
	return self.initStarted
}
func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) StartInit() {
	self.initStarted = true
}