package classfile

var (
	_attrDeprecated = &DeprecatedAttribute{}
	_attrSynthetic  = &SyntheticAttribute{}
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	// case "AnnotationDefault":
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{cp: cp}
	case "Deprecated":
		return _attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp: cp}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{cp: cp}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
	case "RuntimeVisibleAnnotations":
		return &RuntimeVisibleAnnotationsAttribute{}
	case "RuntimeInvisibleAnnotations":
		return &RuntimeInvisibleAnnotationsAttribute{}
	case "RuntimeVisibleParameterAnnotations":
		return &RuntimeVisibleParammeterAnnotationsAttribute{}
	case "AnnotationDefaulte":
		return &AnnotationDefaultAttribute{}
	// case "MethodParameters":
	// case "RuntimeInvisibleTypeAnnotations":

	// case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "SourceDebugExtension":
		return &SourceDebugExtensionAttribute{}
	case "StackMapTable":
		return &StackMapTableAttribute{cp: cp, attributeLength: attrLen}
	case "Synthetic":
		return _attrSynthetic
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
