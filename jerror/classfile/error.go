package jerror

type Error struct{
	msg string
}

var(
	ClassFormatError := Error{"ClassFormatError"}
) 
