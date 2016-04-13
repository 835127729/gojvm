package jutil

import (
	"strings"
)

func Point2Slash(origin string) string {
	return strings.Replace(origin, ".", "/", -1)
}
