package classfile

import (
	"fmt"
	"gojvm/classpath"
	"gojvm/cmdline"
	"testing"
)

func TestClassfile(t *testing.T) {

	cmd := cmdline.ParseCmd()

	if cmd.VersionFlag() {
		fmt.Println("version 0.0.1")
		return
	} else if cmd.HelpFlag() || cmd.Class() == "" {
		cmdline.PrintUsage()
		return
	}
	cp := classpath.Parse(cmd.XjreOption(), cmd.CpOption())
	readClassFile(cp, cmd.Class())
}

func readClassFile(cp *classpath.Classpath, classname string) {
	data, _, err := cp.ReadClass(classname)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	cf, _ := Parse(data)
	printClassfile(cf)
	if cf.superClass != 0 {
		//readClassFile(cp, cf.SuperClassName())
	}
}

func printClassfile(cf *ClassFile) {
	fmt.Println("**********", cf.ClassName(), "***********")
	fmt.Println(cf.ClassName())
	fmt.Println("	minor version", cf.MinorVersion())
	fmt.Println("	major version", cf.MajorVersion())
	printClassFlags(cf.AccessFlags())
	printConstantPool(cf.constantPool)
	fmt.Println("{")
	printFields(cf.Fields())
	printMethods(cf.Methods())
	fmt.Println("}")
	fmt.Println("*********************")
}

const (
	ACC_PUBLIC       = 0x0001 // class field method
	ACC_PRIVATE      = 0x0002 //       field method
	ACC_PROTECTED    = 0x0004 //       field method
	ACC_STATIC       = 0x0008 //       field method
	ACC_FINAL        = 0x0010 // class field method
	ACC_SUPER        = 0x0020 // class
	ACC_SYNCHRONIZED = 0x0020 //             method
	ACC_VOLATILE     = 0x0040 //       field
	ACC_BRIDGE       = 0x0040 //             method
	ACC_TRANSIENT    = 0x0080 //       field
	ACC_VARARGS      = 0x0080 //             method
	ACC_NATIVE       = 0x0100 //             method
	ACC_INTERFACE    = 0x0200 // class
	ACC_ABSTRACT     = 0x0400 // class       method
	ACC_STRICT       = 0x0800 //             method
	ACC_SYNTHETIC    = 0x1000 // class field method
	ACC_ANNOTATION   = 0x2000 // class
	ACC_ENUM         = 0x4000 // class field
)

func printClassFlags(flag uint16) {
	fmt.Print("	flags:")
	if 0 != flag&ACC_PUBLIC {
		fmt.Print("ACC_PUBLIC,")
	}
	if 0 != flag&ACC_FINAL {
		fmt.Print("ACC_FINAL,")
	}
	if 0 != flag&ACC_SUPER {
		fmt.Print("ACC_SUPER,")
	}
	if 0 != flag&ACC_INTERFACE {
		fmt.Print("ACC_INTERFACE,")
	}
	if 0 != flag&ACC_ABSTRACT {
		fmt.Print("ACC_ABSTRACT,")
	}
	if 0 != flag&ACC_STRICT {
		fmt.Print("ACC_STRICT,")
	}
	if 0 != flag&ACC_SYNTHETIC {
		fmt.Print("ACC_SYNTHETIC,")
	}
	if 0 != flag&ACC_ANNOTATION {
		fmt.Print("ACC_ANNOTATION,")
	}
	if 0 != flag&ACC_ENUM {
		fmt.Print("ACC_ENUM,")
	}
	fmt.Println()
}

func printFieldFlags(flag uint16) {
	fmt.Print("    flags:")
	if 0 != flag&ACC_PUBLIC {
		fmt.Print("ACC_PUBLIC,")
	}
	if 0 != flag&ACC_PRIVATE {
		fmt.Print("ACC_PRIVATE,")
	}
	if 0 != flag&ACC_PROTECTED {
		fmt.Print("ACC_PROTECTED,")
	}
	if 0 != flag&ACC_STATIC {
		fmt.Print("ACC_STATIC,")
	}
	if 0 != flag&ACC_FINAL {
		fmt.Print("ACC_FINAL,")
	}
	if 0 != flag&ACC_VOLATILE {
		fmt.Print("ACC_VOLATILE,")
	}
	if 0 != flag&ACC_TRANSIENT {
		fmt.Print("ACC_TRANSIENT,")
	}
	if 0 != flag&ACC_SYNTHETIC {
		fmt.Print("ACC_SYNTHETIC,")
	}
	if 0 != flag&ACC_ENUM {
		fmt.Print("ACC_ENUM,")
	}
	fmt.Println()
}

func printMethodFlags(flag uint16) {
	fmt.Print("    flags:")
	if 0 != flag&ACC_PUBLIC {
		fmt.Print("ACC_PUBLIC,")
	}
	if 0 != flag&ACC_PRIVATE {
		fmt.Print("ACC_PRIVATE,")
	}
	if 0 != flag&ACC_PROTECTED {
		fmt.Print("ACC_PROTECTED,")
	}
	if 0 != flag&ACC_STATIC {
		fmt.Print("ACC_STATIC,")
	}
	if 0 != flag&ACC_FINAL {
		fmt.Print("ACC_FINAL,")
	}
	if 0 != flag&ACC_SYNCHRONIZED {
		fmt.Print("ACC_SYNCHRONIZED,")
	}
	if 0 != flag&ACC_VOLATILE {
		fmt.Print("ACC_VOLATILE,")
	}
	if 0 != flag&ACC_BRIDGE {
		fmt.Print("ACC_BRIDGE,")
	}
	if 0 != flag&ACC_VARARGS {
		fmt.Print("ACC_VARARGS,")
	}
	if 0 != flag&ACC_NATIVE {
		fmt.Print("ACC_NATIVE,")
	}
	if 0 != flag&ACC_ABSTRACT {
		fmt.Print("ACC_ABSTRACT,")
	}
	if 0 != flag&ACC_STRICT {
		fmt.Print("ACC_STRICT,")
	}
	if 0 != flag&ACC_SYNTHETIC {
		fmt.Print("ACC_SYNTHETIC,")
	}
	fmt.Println()
}

/*
func printFlags(flag uint16) {
	fmt.Print("flags:")
	if 0 != flag&ACC_PUBLIC {
		fmt.Print("ACC_PUBLIC,")
	}
	if 0 != flag&ACC_PRIVATE {
		fmt.Print("ACC_PRIVATE,")
	}
	if 0 != flag&ACC_PROTECTED {
		fmt.Print("ACC_PROTECTED,")
	}
	if 0 != flag&ACC_STATIC {
		fmt.Print("ACC_STATIC,")
	}
	if 0 != flag&ACC_FINAL {
		fmt.Print("ACC_FINAL,")
	}
	if 0 != flag&ACC_SUPER {
		fmt.Print("ACC_SUPER,")
	}
	if 0 != flag&ACC_SYNCHRONIZED {
		fmt.Print("ACC_SYNCHRONIZED,")
	}
	if 0 != flag&ACC_VOLATILE {
		fmt.Print("ACC_VOLATILE,")
	}
	if 0 != flag&ACC_TRANSIENT {
		fmt.Print("ACC_TRANSIENT,")
	}
	if 0 != flag&ACC_VARARGS {
		fmt.Print("ACC_VARARGS,")
	}
	if 0 != flag&ACC_NATIVE {
		fmt.Print("ACC_NATIVE,")
	}
	if 0 != flag&ACC_INTERFACE {
		fmt.Print("ACC_INTERFACE,")
	}
	if 0 != flag&ACC_ABSTRACT {
		fmt.Print("ACC_ABSTRACT,")
	}
	if 0 != flag&ACC_STRICT {
		fmt.Print("ACC_STRICT,")
	}
	if 0 != flag&ACC_SYNTHETIC {
		fmt.Print("ACC_SYNTHETIC,")
	}
	if 0 != flag&ACC_ANNOTATION {
		fmt.Print("ACC_ANNOTATION,")
	}
	if 0 != flag&ACC_ENUM {
		fmt.Print("ACC_ENUM,")
	}
	fmt.Println()
}
*/

func printConstantPool(cp ConstantPool) {
	fmt.Println("Constant pool:")
	cpCount := len(cp)
	for i := 1; i < cpCount; i++ {
		cpInfo := cp[i]
		fmt.Print("	#", i, " = ", cpInfo.toString())

		switch cpInfo.(type) {
		case *ConstantUtf8Info:
			utf8Info := cpInfo.(*ConstantUtf8Info)
			fmt.Print("	", utf8Info.Str())
		case *ConstantIntegerInfo:
			intInfo := cpInfo.(*ConstantIntegerInfo)
			fmt.Print("	", intInfo.Value())
		case *ConstantFloatInfo:
			floatInfo := cpInfo.(*ConstantFloatInfo)
			fmt.Print("	", floatInfo.Value())
		case *ConstantLongInfo:
			longInfo := cpInfo.(*ConstantLongInfo)
			fmt.Print("	", longInfo.Value())
			i++
		case *ConstantDoubleInfo:
			doubleInfo := cpInfo.(*ConstantDoubleInfo)
			fmt.Print("	", doubleInfo.Value())
			i++
		case *ConstantStringInfo:
			stringInfo := cpInfo.(*ConstantStringInfo)
			fmt.Print("	//", stringInfo.String())
		case *ConstantNameAndTypeInfo:
			nameAndTypeInfo := cpInfo.(*ConstantNameAndTypeInfo)
			fmt.Print("	#", nameAndTypeInfo.nameIndex, ",#", nameAndTypeInfo.descriptorIndex)
			fmt.Print("	//", cp.getUtf8(nameAndTypeInfo.nameIndex), ";", cp.getUtf8(nameAndTypeInfo.descriptorIndex))
		case *ConstantClassInfo:
			classInfo := cpInfo.(*ConstantClassInfo)
			fmt.Print("	#", classInfo.nameIndex)
			fmt.Print("	//", cp.getUtf8(classInfo.nameIndex))
		case *ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*ConstantFieldrefInfo)
			fmt.Print("	#", fieldrefInfo.classIndex, ",#", fieldrefInfo.nameAndTypeIndex)
			nameAndType := cp.getConstantInfo(fieldrefInfo.nameAndTypeIndex).(*ConstantNameAndTypeInfo)
			fmt.Print("	//", cp.getClassName(fieldrefInfo.classIndex), ";", cp.getUtf8(nameAndType.nameIndex), ";", cp.getUtf8(nameAndType.descriptorIndex))
		case *ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*ConstantMethodrefInfo)
			fmt.Print("	#", methodrefInfo.classIndex, ",#", methodrefInfo.nameAndTypeIndex)
			nameAndType := cp.getConstantInfo(methodrefInfo.nameAndTypeIndex).(*ConstantNameAndTypeInfo)
			fmt.Print(" //", cp.getClassName(methodrefInfo.classIndex), ";", cp.getUtf8(nameAndType.nameIndex), ";", cp.getUtf8(nameAndType.descriptorIndex))
		case *ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*ConstantInterfaceMethodrefInfo)
			fmt.Print("	#", methodrefInfo.classIndex, ",#", methodrefInfo.nameAndTypeIndex)
		default:
			// todo
		}
		fmt.Println()
	}
}

func printFields(fields []*MemberInfo) {
	for _, field := range fields {
		fmt.Println(" **filed**", field.Name())
		fmt.Println("    descriptor:", field.Descriptor())
		printFieldFlags(field.accessFlags)
		fmt.Println()
	}
}

func printMethods(methods []*MemberInfo) {
	for _, method := range methods {
		fmt.Println(" **method**", method.Name())
		fmt.Println("    descriptor:", method.Descriptor())
		printMethodFlags(method.accessFlags)
		printCode(method.CodeAttribute())
	}
}

func printCode(codeAttribute *CodeAttribute) {
	if codeAttribute == nil {
		return
	}
	fmt.Println("    code:")
	fmt.Println("      stack=", codeAttribute.MaxStack(), ",locals=", codeAttribute.MaxLocals())
	count := len(codeAttribute.Code())
	for i := 0; i < count; i++ {
		code := codeAttribute.Code()[i]
		fmt.Print("		", i, ":")
		switch code {
		case 0x00:
			fmt.Println("nop")
		case 0x01:
			fmt.Println("aconst_null")
		case 0x02:
			fmt.Println("iconst_m1")
		case 0x03:
			fmt.Println("iconst_0")
		case 0x04:
			fmt.Println("iconst_1")
		case 0x05:
			fmt.Println("iconst_2")
		case 0x06:
			fmt.Println("iconst_3")
		case 0x07:
			fmt.Println("iconst_4")
		case 0x08:
			fmt.Println("iconst_5")
		case 0x09:
			fmt.Println("lconst_0")
		case 0x0a:
			fmt.Println("lconst_1")
		case 0x0b:
			fmt.Println("fconst_0")
		case 0x0c:
			fmt.Println("fconst_1")
		case 0x0d:
			fmt.Println("fconst_2")
		case 0x0e:
			fmt.Println("dconst_0")
		case 0x0f:
			fmt.Println("dconst_1")
		case 0x10:
			fmt.Println("bipush")
			i++
		case 0x11:
			fmt.Println("sipush")
			i += 2
		case 0x12:
			fmt.Println("ldc")
			i++
		case 0x13:
			fmt.Println("ldc_w")
			i += 2
		case 0x14:
			fmt.Println("ldc2_2")
			i += 2
		case 0x15:
			fmt.Println("iload")
			i++
		case 0x16:
			fmt.Println("lload")
			i++
		case 0x17:
			fmt.Println("fload")
			i++
		case 0x18:
			fmt.Println("dload")
			i++
		case 0x19:
			fmt.Println("aload")
			i++
		case 0x1a:
			fmt.Println("iload_0")
		case 0x1b:
			fmt.Println("iload_1")
		case 0x1c:
			fmt.Println("iload_2")
		case 0x1d:
			fmt.Println("iload_3")
		case 0x1e:
			fmt.Println("lload_0")
		case 0x1f:
			fmt.Println("lload_1")
		case 0x20:
			fmt.Println("lload_2")
		case 0x21:
			fmt.Println("lload_3")
		case 0x22:
			fmt.Println("fload_0")
		case 0x23:
			fmt.Println("fload_1")
		case 0x24:
			fmt.Println("fload_2")
		case 0x25:
			fmt.Println("fload_3")
		case 0x26:
			fmt.Println("dload_0")
		case 0x27:
			fmt.Println("dload_1")
		case 0x28:
			fmt.Println("dload_2")
		case 0x29:
			fmt.Println("dload_3")
		case 0x2a:
			fmt.Println("aload_0")
		case 0x2b:
			fmt.Println("aload_1")
		case 0x2c:
			fmt.Println("aload_2")
		case 0x2d:
			fmt.Println("aload_3")
		case 0x2e:
			fmt.Println("iaload")
		case 0x2f:
			fmt.Println("laload")
		case 0x30:
			fmt.Println("faload")
		case 0x31:
			fmt.Println("daload")
		case 0x32:
			fmt.Println("aaload")
		case 0x33:
			fmt.Println("baload")
		case 0x34:
			fmt.Println("caaload")
		case 0x35:
			fmt.Println("saload")
		case 0x36:
			fmt.Println("istore")
			i++
		case 0x37:
			fmt.Println("lstore")
			i++
		case 0x38:
			fmt.Println("fstore")
			i++
		case 0x39:
			fmt.Println("dstore")
			i++
		case 0x3a:
			fmt.Println("astore")
			i++
		case 0x3b:
			fmt.Println("istore_0")
		case 0x3c:
			fmt.Println("istore_1")
		case 0x3d:
			fmt.Println("istore_2")
		case 0x3e:
			fmt.Println("istore_3")
		case 0x3f:
			fmt.Println("lstore_0")
		case 0x40:
			fmt.Println("lstore_1")
		case 0x41:
			fmt.Println("lstore_2")
		case 0x42:
			fmt.Println("lstore_3")
		case 0x43:
			fmt.Println("fstore_0")
		case 0x44:
			fmt.Println("fstore_1")
		case 0x45:
			fmt.Println("fstore_2")
		case 0x46:
			fmt.Println("fstore_3")
		case 0x47:
			fmt.Println("dstore_0")
		case 0x48:
			fmt.Println("dstore_1")
		case 0x49:
			fmt.Println("dstore_2")
		case 0x4a:
			fmt.Println("dstore_3")
		case 0x4b:
			fmt.Println("astore_0")
		case 0x4c:
			fmt.Println("astore_1")
		case 0x4d:
			fmt.Println("astore_2")
		case 0x4e:
			fmt.Println("astore_3")
		case 0x4f:
			fmt.Println("iastore")
		case 0x50:
			fmt.Println("lastore")
		case 0x51:
			fmt.Println("fastore")
		case 0x52:
			fmt.Println("dastore")
		case 0x53:
			fmt.Println("aastore")
		case 0x54:
			fmt.Println("bastore")
		case 0x55:
			fmt.Println("castore")
		case 0x56:
			fmt.Println("sastore")
		case 0x57:
			fmt.Println("pop")
		case 0x58:
			fmt.Println("pop2")
		case 0x59:
			fmt.Println("dup")
		case 0x5a:
			fmt.Println("dup_x1")
		case 0x5b:
			fmt.Println("dup_x2")
		case 0x5c:
			fmt.Println("dup2")
		case 0x5d:
			fmt.Println("dup2_x1")
		case 0x5e:
			fmt.Println("dup2_x2")
		case 0x5f:
			fmt.Println("swap")
		case 0x60:
			fmt.Println("iadd")
		case 0x61:
			fmt.Println("ladd")
		case 0x62:
			fmt.Println("fadd")
		case 0x63:
			fmt.Println("dadd")
		case 0x64:
			fmt.Println("isub")
		case 0x65:
			fmt.Println("lsub")
		case 0x66:
			fmt.Println("fsub")
		case 0x67:
			fmt.Println("dsub")
		case 0x68:
			fmt.Println("imul")
		case 0x69:
			fmt.Println("lmul")
		case 0x6a:
			fmt.Println("fmul")
		case 0x6b:
			fmt.Println("dmul")
		case 0x6c:
			fmt.Println("idiv")
		case 0x6d:
			fmt.Println("ldiv")
		case 0x6e:
			fmt.Println("fdiv")
		case 0x6f:
			fmt.Println("ddiv")
		case 0x70:
			fmt.Println("irem")
		case 0x71:
			fmt.Println("lrem")
		case 0x72:
			fmt.Println("frem")
		case 0x73:
			fmt.Println("drem")
		case 0x74:
			fmt.Println("ineg")
		case 0x75:
			fmt.Println("lneg")
		case 0x76:
			fmt.Println("fneg")
		case 0x77:
			fmt.Println("dneg")
		case 0x78:
			fmt.Println("ishl")
		case 0x79:
			fmt.Println("lshl")
		case 0x7a:
			fmt.Println("ishr")
		case 0x7b:
			fmt.Println("lshr")
		case 0x7c:
			fmt.Println("iushr")
		case 0x7d:
			fmt.Println("lushr")
		case 0x7e:
			fmt.Println("iand")
		case 0x7f:
			fmt.Println("land")
		case 0x80:
			fmt.Println("ior")
		case 0x81:
			fmt.Println("lor")
		case 0x82:
			fmt.Println("ixor")
		case 0x83:
			fmt.Println("lxor")
		case 0x84:
			fmt.Println("iinc")
			i += 2
		case 0x85:
			fmt.Println("i2l")
		case 0x86:
			fmt.Println("i2f")
		case 0x87:
			fmt.Println("i2d")
		case 0x88:
			fmt.Println("l2i")
		case 0x89:
			fmt.Println("l2f")
		case 0x8a:
			fmt.Println("l2d")
		case 0x8b:
			fmt.Println("f2i")
		case 0x8c:
			fmt.Println("f2l")
		case 0x8d:
			fmt.Println("f2d")
		case 0x8e:
			fmt.Println("d2i")
		case 0x8f:
			fmt.Println("d2l")
		case 0x90:
			fmt.Println("d2f")
		case 0x91:
			fmt.Println("i2b")
		case 0x92:
			fmt.Println("i2c")
		case 0x93:
			fmt.Println("i2s")
		case 0x94:
			fmt.Println("lcmp")
		case 0x95:
			fmt.Println("fcmpl")
		case 0x96:
			fmt.Println("fcmpg")
		case 0x97:
			fmt.Println("dcmpl")
		case 0x98:
			fmt.Println("fcmpg")
		case 0x99:
			fmt.Println("ifeq")
			i += 2
		case 0x9a:
			fmt.Println("ifne")
			i += 2
		case 0x9b:
			fmt.Println("iflt")
			i += 2
		case 0x9c:
			fmt.Println("ifge")
			i += 2
		case 0x9d:
			fmt.Println("ifgt")
			i += 2
		case 0x9e:
			fmt.Println("ifle")
			i += 2
		case 0x9f:
			fmt.Println("if_icmpeq")
			i += 2
		case 0xa0:
			fmt.Println("if_icmpne")
			i += 2
		case 0xa1:
			fmt.Println("if_icmplt")
			i += 2
		case 0xa2:
			fmt.Println("if_icmpge")
			i += 2
		case 0xa3:
			fmt.Println("if_icmpgt")
			i += 2
		case 0xa4:
			fmt.Println("if_icmple")
			i += 2
		case 0xa5:
			fmt.Println("if_acmpeq")
			i += 2
		case 0xa6:
			fmt.Println("if_acmpne")
			i += 2
		case 0xa7:
			fmt.Println("goto")
			i += 2
		// case 0xa8:
		// 	return &JSR{}
		// case 0xa9:
		// 	return &RET{}
		case 0xaa:
			fmt.Println("tableswitch")
			j := 0
			for ; j < 3; j++ {
				if codeAttribute.Code()[i+j] != 0x00 {
					break
				}
			}
			i = i + j + 4
			low := (codeAttribute.Code()[i+1] << 24) | (codeAttribute.Code()[i+2] << 16) | (codeAttribute.Code()[i+3] << 8) | codeAttribute.Code()[i+4]
			high := (codeAttribute.Code()[i+5] << 24) | (codeAttribute.Code()[i+6] << 16) | (codeAttribute.Code()[i+7] << 8) | codeAttribute.Code()[i+8]
			i = i + 8
			offset := high - low + 1
			i = i + int(offset)
		case 0xab:
			fmt.Println("lookupwitch")
			j := 0
			for ; j < 3; j++ {
				if codeAttribute.Code()[i+j] != 0x00 {
					break
				}
			}
			i = i + j + 4
			offset := (codeAttribute.Code()[i+1] << 24) | (codeAttribute.Code()[i+2] << 16) | (codeAttribute.Code()[i+3] << 8) | codeAttribute.Code()[i+4]
			i = i + 4
			i = i + int(offset)
		case 0xac:
			fmt.Println("ireturn")
		case 0xad:
			fmt.Println("lreturn")
		case 0xae:
			fmt.Println("freturn")
		case 0xaf:
			fmt.Println("dreturn")
		case 0xb0:
			fmt.Println("areturn")
		case 0xb1:
			fmt.Println("return")
		case 0xb2:
			fmt.Println("getstatic")
			i += 2
		case 0xb3:
			fmt.Println("putstatic")
			i += 2
		case 0xb4:
			fmt.Println("getfield")
			i += 2
		case 0xb5:
			fmt.Println("putfield")
			i += 2
		case 0xb6:
			fmt.Println("invokevirtual")
			i += 2
		case 0xb7:
			fmt.Println("invokespecial")
			i += 2
		case 0xb8:
			fmt.Println("invokestatic")
			i += 2
		case 0xb9:
			fmt.Println("invokeinterface")
			i += 2
		// case 0xba:
		// 	return &INVOKE_DYNAMIC{}
		case 0xbb:
			fmt.Println("new")
			i += 2
		case 0xbc:
			fmt.Println("newarray")
			i++
		case 0xbd:
			fmt.Println("anewarray")
			i += 2
		case 0xbe:
			fmt.Println("arraylength")
		case 0xbf:
			fmt.Println("athrow")
		case 0xc0:
			fmt.Println("checkcast")
			i += 2
		case 0xc1:
			fmt.Println("instanceof")
			i += 2
		case 0xc2:
			fmt.Println("monitorenter")
		case 0xc3:
			fmt.Println("monitorexit")
		case 0xc4:
			fmt.Println("wide")
			//todo
		case 0xc5:
			fmt.Println("multianewarray")
			i += 3
		case 0xc6:
			fmt.Println("ifnull")
			i += 2
		case 0xc7:
			fmt.Println("ifnonnull")
			i += 2
		case 0xc8:
			fmt.Println("goto_w")
			i += 4
		// case 0xc9:
		// 	return &JSR_W{}
		// case 0xca: breakpoint
		case 0xfe:
			fmt.Println("invokenative")
			i += 2
		// case 0xff: impdep2
		default:
			panic(fmt.Errorf("Unsupported opcode: 0x%x!", code))
		}
	}
	printExceptionsAttribute(codeAttribute.exceptionTable, codeAttribute.cp)
	printLineNumberTableAttribute(codeAttribute.LineNumberTableAttribute())
	printLocalVariableTableAttribute(codeAttribute.LocalVariableTableAttribute())
	printStackMapTable(codeAttribute.StackMapTableAttribute())
}

func printExceptionsAttribute(table []*ExceptionTableEntry, cp ConstantPool) {
	if table == nil || len(table) == 0 {
		return
	}
	fmt.Println("     Exception table:")
	fmt.Println("     	from    to	target	type")
	for _, entry := range table {
		fmt.Println("     	", entry.startPc, "	", entry.endPc, "	", entry.handlerPc, "	", cp.getClassName(entry.catchType))
	}
}

func printLineNumberTableAttribute(table *LineNumberTableAttribute) {
	if table == nil || len(table.lineNumberTable) == 0 {
		return
	}
	fmt.Println("     LineNumberTable:")
	for _, entry := range table.lineNumberTable {
		fmt.Println("		line", entry.lineNumber, ":", entry.startPc)
	}
}

func printLocalVariableTableAttribute(table *LocalVariableTableAttribute) {
	if table == nil || len(table.localVariableTable) == 0 {
		return
	}
	fmt.Println("     LocalVariableTable:")
	fmt.Println("     	Start	Length	Slot	Name	Signature:")
	for _, entry := range table.localVariableTable {
		fmt.Println("     	", entry.startPc, "	", entry.length, "	", entry.index, "	", entry.Name(), "	", entry.Descriptor())
	}
}

func printStackMapTable(table *StackMapTableAttribute) {
	if table == nil || len(table.entries) == 0 {
		return
	}
	fmt.Println("     StackMapTable: number_of_entries = :", len(table.entries))
	for _, stackMapFrame := range table.entries {
		switch stackMapFrame.(type) {
		case *SameFrame:
			sameFrame := stackMapFrame.(*SameFrame)
			fmt.Println("		frametype =", sameFrame.frameType, "/* same */")
		case *SameLocals1StackItemFrame:
			sameLocals1StackItemFrame := stackMapFrame.(*SameLocals1StackItemFrame)
			fmt.Println("		frametype =", sameLocals1StackItemFrame.frameType, "/* same_locals_1_stack_item */")
			fmt.Print("			stack = [")
			printVerificationTypeInfos(sameLocals1StackItemFrame.stack)
			fmt.Println("]")
		case *SameLocals1StackItemFrameExtended:
			sameLocals1StackItemFrameExtended := stackMapFrame.(*SameLocals1StackItemFrameExtended)
			fmt.Println("		frametype =", sameLocals1StackItemFrameExtended.frameType, "/* same_locals_1_stack_item_extended */")
			fmt.Println("			offset = ", sameLocals1StackItemFrameExtended.offset())
			fmt.Print("			stack = [")
			printVerificationTypeInfos(sameLocals1StackItemFrameExtended.stack)
			fmt.Println("]")
		case *ChopFrame:
			chopFrame := stackMapFrame.(*ChopFrame)
			fmt.Println("		frametype =", chopFrame.frameType, "/* chop */")
			fmt.Println("			offset = ", chopFrame.offset())
		case *SameFrameExtended:
			sameFrameExtended := stackMapFrame.(*SameFrameExtended)
			fmt.Println("		frametype =", sameFrameExtended.frameType, "/* chop */")
			fmt.Println("			offset = ", sameFrameExtended.offset())
		case *AppendFrame:
			appendFrame := stackMapFrame.(*AppendFrame)
			fmt.Println("		frametype =", appendFrame.frameType, "/* chop */")
			fmt.Println("			offset = ", appendFrame.offset())
			fmt.Print("			locals = [")
			printVerificationTypeInfos(appendFrame.locals)
			fmt.Println("]")
		case *FullFrame:
			fullFrame := stackMapFrame.(*FullFrame)
			fmt.Println("		frametype =", fullFrame.frameType, "/* chop */")
			fmt.Println("			offset = ", fullFrame.offset())
			fmt.Print("			stack = [")
			printVerificationTypeInfos(fullFrame.stack)
			fmt.Println("]")
			fmt.Print("			locals= [")
			printVerificationTypeInfos(fullFrame.locals)
			fmt.Println("]")
		}

	}
}

func printVerificationTypeInfos(infos []VerificationTypeInfo) {
	for _, info := range infos {
		fmt.Print(info.name(), ",")
	}
}
