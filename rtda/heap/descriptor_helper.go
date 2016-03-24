package heap

// [XXX -> [XXX
// LXXX; -> XXX
// I -> int ...
func getClassName(descriptor string) string {
	switch descriptor[0] {
	case '[': // array
		return descriptor
	case 'L': // object
		return descriptor[1 : len(descriptor)-1]
	default: // primirive
		return getPrimitiveType(descriptor)
	}
}
