package npp

/*
#include <nppdefs.h>
*/
import "C"

import (
	"unsafe"

	"github.com/l0rem1psum/go-cuda-toolkit/npp/internal"
)

type Status = internal.Status

type StreamContext struct {
	sc *C.NppStreamContext
}

func (s *StreamContext) C() unsafe.Pointer {
	return unsafe.Pointer(s.sc)
}
