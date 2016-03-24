package heap

import (
	"fmt"
	"gojvm/classfile"
	"strings"
)

type Class struct {
	AccessFlags
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	superClass        *Class
	interfaces        []*Class
	jClass            *Object // java.lang.Class instance
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	loader            *ClassLoader
	initStarted       bool
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
	return class
}

func (self *Class) ToString() {
	fmt.Println("Class:")
	fmt.Println("name:", self.name)
	fmt.Println("superClassName:", self.superClassName)
	fmt.Println("interfaceNames:", self.interfaceNames)
	self.constantPool.ToString()
}

func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

//获取域
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

//获取方法
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

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) GetClinitMethod() *Method {
	return self.getMethod("<clinit>", "()V", true)
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) GetMainMethod() *Method {
	return self.getMethod("main", "([Ljava/lang/String;)V", true)
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

// getters
func (self *Class) Name() string {
	return self.name
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
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

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}
