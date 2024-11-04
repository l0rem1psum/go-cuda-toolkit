package npp

/*
#include <nppdefs.h>
*/
import "C"
import "unsafe"

func (sc *StreamContext) AsC() unsafe.Pointer {
	return unsafe.Pointer(sc.asC())
}
