package npp

/*
#include <nppdefs.h>
*/
import "C"
import "unsafe"

func NewNppStreamContextFromC(c unsafe.Pointer) *StreamContext {
	return &StreamContext{(*C.NppStreamContext)(c)}
}

func (s *StreamContext) C() unsafe.Pointer {
	return unsafe.Pointer(s.sc)
}
