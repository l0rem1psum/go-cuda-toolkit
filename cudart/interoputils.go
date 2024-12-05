package cudart

/*
#include <cuda_runtime_api.h>
*/
import "C"
import "unsafe"

func NewCUDAStreamFromC(c unsafe.Pointer) *CUDAStream {
	return &CUDAStream{C.cudaStream_t(c)}
}

func (s *CUDAStream) C() unsafe.Pointer {
	if s == nil {
		return nil
	}
	return unsafe.Pointer(s.s)
}
