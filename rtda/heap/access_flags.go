package heap

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

type AccessFlags struct {
	accessFlags uint16
}

func (self *AccessFlags) GetAccessFlags() uint16 {
	return self.accessFlags
}

func (self *AccessFlags) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *AccessFlags) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *AccessFlags) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *AccessFlags) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *AccessFlags) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *AccessFlags) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *AccessFlags) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *AccessFlags) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *AccessFlags) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *AccessFlags) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *AccessFlags) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *AccessFlags) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *AccessFlags) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *AccessFlags) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *AccessFlags) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}
func (self *AccessFlags) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *AccessFlags) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *AccessFlags) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
