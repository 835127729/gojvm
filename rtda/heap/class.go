package heap

import (
	"fmt"
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
	//class.sourceFile = getSourceFile(cf)
	return class
}

func (self *Class) ToString() {
	fmt.Println("Class:")
	fmt.Println("name:", self.name)
	fmt.Println("superClassName:", self.superClassName)
	fmt.Println("interfaceNames:", self.interfaceNames)
	self.constantPool.ToString()
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) IsSuperClassOf(son *Class) bool {
	for son.superClass != nil {
		if son.superClass == self {
			return true
		}
	}
	return false
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for k := self; k != nil; k = k.superClass {
		for _, field := range k.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	// todo
	return nil
}

func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for k := self; k != nil; k = k.superClass {
		for _, method := range k.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	// todo
	return nil
}

func (self *Class) GetMainMethod() *Method {
	return self.getMethod("main", "([Ljava/lang/String;)V", true)
}

// getters
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

func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) Interfaces() []*Class {
	return self.interfaces
}
