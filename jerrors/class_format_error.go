package jerrors

type ClassFormatError struct {
	cause string
}

type UnsupportedClassVersionError struct {
	ClassFormatError
}
