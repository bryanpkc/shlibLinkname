package linkname1

import (
	_ "runtime"
	_ "unsafe"
)

type errorString string

//go:linkname errorString.Error runtime.errorString.Error
func (e errorString) Error() string

func Fail() {
	s := errorString("linkname2")
	println(s.Error())
}
