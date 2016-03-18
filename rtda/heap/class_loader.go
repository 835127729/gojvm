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

func GetBootLoader() *ClassLoader {
	return bootLoader
}

func (self *ClassLoader) _init() {
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

func (self *ClassLoader) loadArrayClass(className string) *Class {
	class := &Class{
		AccessFlags: AccessFlags{ACC_PUBLIC}, // todo
		name:        className,
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[className] = class
	return class
}

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

func (self *ClassLoader) _loadClass(name string, data []byte) *Class {
	class := self.parseClassData(name, data)
	/*
		self.resolveSuperClass(class)
		self.resolveInterfaces(class)
		calcStaticFieldSlotIds(class)
		calcInstanceFieldSlotIds(class)
		createVtable(class)
		prepare(class)
	*/
	// todo
	//class.classLoader = self
	self.classMap[name] = class

	if _jlClassClass != nil {
		class.jClass = _jlClassClass.NewObject()
		class.jClass.extra = class
	}

	return class
}

func (self *ClassLoader) parseClassData(name string, data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		// todo
		//ClassFormatError,UnsupportedClassVersionError
		panic("failed to parse class file: " + name + "!" + err.Error())
	}

	return newClass(cf)
}
