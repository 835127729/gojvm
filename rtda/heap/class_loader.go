package heap

import (
	"fmt"
	"gojvm/classfile"
	"gojvm/classpath"
	"gojvm/options"
)

const (
	jlObjectClassName       = "java/lang/Object"
	jlClassClassName        = "java/lang/Class"
	jlStringClassName       = "java/lang/String"
	jlThreadClassName       = "java/lang/Thread"
	jlCloneableClassName    = "java/lang/Cloneable"
	ioSerializableClassName = "java/io/Serializable"
)

var (
	bootLoader           *ClassLoader // bootstrap class loader
	_jlObjectClass       *Class
	_jlClassClass        *Class
	_jlStringClass       *Class
	_jlThreadClass       *Class
	_jlCloneableClass    *Class
	_ioSerializableClass *Class
)

type ClassLoader struct {
	classPath *classpath.ClassPath
	classMap  map[string]*Class
}

func InitBootLoader(cp *classpath.ClassPath) {
	bootLoader = &ClassLoader{
		classPath: cp,
		classMap:  map[string]*Class{},
	}
	bootLoader._init()
}

func (self *ClassLoader) _init() {
	self.loadBasicClasses()
	self.loadPrimitiveClasses()
	self.loadPrimitiveArrayClasses()
}

//加载一些基础类
func (self *ClassLoader) loadBasicClasses() {
	_jlObjectClass = self.LoadClass(jlObjectClassName)
	_jlClassClass = self.LoadClass(jlClassClassName)
	for _, class := range self.classMap {
		if class.jClass == nil {
			class.jClass = _jlClassClass.NewObject()
		}
	}
	_jlCloneableClass = self.LoadClass(jlCloneableClassName)
	_ioSerializableClass = self.LoadClass(ioSerializableClassName)
	_jlThreadClass = self.LoadClass(jlThreadClassName)
	_jlStringClass = self.LoadClass(jlStringClassName)
}

//加载基本类型
func (self *ClassLoader) loadPrimitiveClasses() {
	for _, primitiveType := range PrimitiveTypes {
		self.LoadClass(primitiveType.WrapperClassName)
	}
}

//加载基本类型数组
func (self *ClassLoader) loadPrimitiveArrayClasses() {
	for _, primitiveType := range PrimitiveTypes {
		class := &Class{
			AccessFlags: AccessFlags{ACC_PUBLIC},
			name:        primitiveType.ArrayClassName,
			superClass:  self.LoadClass("java/lang/Object"),
			interfaces: []*Class{
				self.LoadClass("java/lang/Cloneable"),
				self.LoadClass("java/io/Serializable"),
			},
		}
		class.jClass = _jlClassClass.NewObject()
		self.classMap[primitiveType.ArrayClassName] = class
	}
}

func (self *ClassLoader) LoadClass(className string) *Class {
	if class, ok := self.classMap[className]; ok {
		return class
	}
	if className[0] != '[' {
		return self.reallyLoadClass(className)
	} else {
		// array class
		return self.loadArrayClass(className)
	}
}

//加载数组
func (self *ClassLoader) loadArrayClass(className string) *Class {
	componentClass := self.LoadClass(getComponentClassName(className))
	class := &Class{
		AccessFlags: AccessFlags{componentClass.AccessFlags.GetAccessFlags()},
		name:        className,
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	class.jClass = _jlClassClass.NewObject()
	self.classMap[className] = class
	return class
}

//加载非数组
func (self *ClassLoader) reallyLoadClass(name string) *Class {
	cpEntry, data := self.readClassData(name)
	class := self._loadClass(name, data)

	if options.VerboseClass {
		fmt.Printf("[Loaded %s from %s]\n", name, cpEntry)
	}

	return class
}

func (self *ClassLoader) readClassData(name string) (classpath.Entry, []byte) {
	cpEntry, classData, err := self.classPath.ReadClass(name)
	//todo
	if err != nil {
		panic("classNotFoundException")
	}

	return cpEntry, classData
}

//步骤
func (self *ClassLoader) _loadClass(name string, data []byte) *Class {
	class := self.parseClassData(name, data)
	link(class)
	prepare(class)
	//createVtable(class)
	return class
}

//加载
func (self *ClassLoader) parseClassData(name string, data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		// todo
		//ClassFormatError,UnsupportedClassVersionError
		panic("failed to parse class file: " + name + "!" + err.Error())
	}

	class := newClass(cf)
	self.classMap[name] = class
	if _jlClassClass != nil {
		class.jClass = _jlClassClass.NewObject()
	}
	return class
}

//链接
func link(class *Class) {
	resolveSuperClass(class)
	resolveInterfaces(class)
}

// jvms 5.4.3.1
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = bootLoader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = bootLoader.LoadClass(interfaceName)
		}
	}
}

//准备
func prepare(class *Class) {
	calcStaticFieldSlotIds(class)
	calcInstanceFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(bootLoader, goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}

//getter
func GetBootLoader() *ClassLoader {
	return bootLoader
}
