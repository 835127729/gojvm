package heap

var primitiveTypes = []*PrimitiveType{
	&PrimitiveType{"V", "[V", "void", "java/lang/Void"},
	&PrimitiveType{"Z", "[Z", "boolean", "java/lang/Boolean"},
	&PrimitiveType{"B", "[B", "byte", "java/lang/Byte"},
	&PrimitiveType{"C", "[C", "char", "java/lang/Character"},
	&PrimitiveType{"S", "[S", "short", "java/lang/Short"},
	&PrimitiveType{"I", "[I", "int", "java/lang/Integer"},
	&PrimitiveType{"J", "[J", "long", "java/lang/Long"},
	&PrimitiveType{"F", "[F", "float", "java/lang/Float"},
	&PrimitiveType{"D", "[D", "double", "java/lang/Double"},
}

// type jboolean bool
// type jbyte int8
// type jchar uint16
// type jshort int16
// type jint int32
// type jlong int64
// type jfloat float32
// type jdouble float64

type PrimitiveType struct {
	Descriptor       string
	ArrayClassName   string
	Name             string
	WrapperClassName string
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
//将类名转化成对应数组类名
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
//从数组类名中提取组件类名
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

// [XXX => [XXX
// int  => I
// XXX  => LXXX;
//将类名转化成描述符
func toDescriptor(className string) string {
	if className[0] == '[' {
		// array
		return className
	}
	for _, primitiveType := range primitiveTypes {
		if className == primitiveType.Name {
			// primitive
			return primitiveType.Descriptor
		}
	}
	// object
	return "L" + className + ";"
}

// [XXX  => [XXX
// LXXX; => XXX
// I     => int
//获取类名
func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		// array
		return descriptor
	}
	if descriptor[0] == 'L' {
		// object
		return descriptor[1 : len(descriptor)-1]
	}
	for _, primitiveType := range primitiveTypes {
		if descriptor == primitiveType.Descriptor {
			// primitive
			return primitiveType.Name
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
