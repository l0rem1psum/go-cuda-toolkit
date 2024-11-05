package cudadrv

/*
#include <cuda.h>
*/
import "C"
import "unsafe"

func NewCUcontextFromC(c unsafe.Pointer) *CUcontext {
	return &CUcontext{c: (C.CUcontext)(c)}
}

func (c *CUcontext) C() unsafe.Pointer {
	return unsafe.Pointer(c.c)
}

func NewCUdeviceFromC(d unsafe.Pointer) *CUdevice {
	return &CUdevice{d: *(*C.CUdevice)(d)}
}

func (d *CUdevice) C() unsafe.Pointer {
	return unsafe.Pointer(&d.d)
}

func NewCUstreamFromC(s unsafe.Pointer) *CUstream {
	return &CUstream{s: (C.CUstream)(s)}
}

func (s *CUstream) C() unsafe.Pointer {
	return unsafe.Pointer(s.s)
}
