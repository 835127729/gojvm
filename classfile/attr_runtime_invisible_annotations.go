package classfile

//RuntimeInvisibleAnnotations
type RuntimeInvisibleAnnotationsAttribute struct {
	annotation []Annotation
}

func (self *RuntimeInvisibleAnnotationsAttribute) readInfo(reader *ClassReader) {
	numAnnotations := reader.readUint16()
	self.annotation = make([]Annotation, numAnnotations)
	for _, annotation := range self.annotation {
		annotation.typeIndex = reader.readUint16()
		numElementValuePairs := reader.readUint16()
		annotation.elementValuePairs = make([]ElementValuePair, numElementValuePairs)
		for _, elementValuePair := range annotation.elementValuePairs {
			elementValuePair.elementNameIndex = reader.readUint16()
			elementValuePair.value = readElementValue(reader)
		}
	}
}

func readElementValue(reader *ClassReader) ElementValue {
	tag := reader.readUint8()
	switch tag {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 's':
		constValueIndex := ConstValueIndex(reader.readUint16())
		return constValueIndex
	case 'e':
		enumConstValue := &EnumConstValue{
			typeNameIndex:  reader.readUint16(),
			constNameIndex: reader.readUint16(),
		}
		return enumConstValue
	case 'c':
		classInfoIndex := ClassInfoIndex(reader.readUint16())
		return classInfoIndex
	case '[':
		numValue := reader.readUint16()
		arrayValue := &ArrayValue{
			values: make([]ElementValue, numValue),
		}
		for i, _ := range arrayValue.values {
			arrayValue.values[i] = readElementValue(reader)
		}
		return arrayValue
	default:
		panic("RuntimeVisibleAnnotations error")
	}
}

type Annotation struct {
	typeIndex         uint16
	elementValuePairs []ElementValuePair
}

type ElementValuePair struct {
	elementNameIndex uint16
	value            ElementValue
}

//element_value
type ElementValue interface{}

type ConstValueIndex uint16

type EnumConstValue struct {
	typeNameIndex  uint16
	constNameIndex uint16
}

type ClassInfoIndex uint16

type ArrayValue struct {
	values []ElementValue
}
