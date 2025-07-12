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

func (h *CUDAIPCMemHandle) Handle() []byte {
	return C.GoBytes(unsafe.Pointer(&h.handle.reserved[0]), 64) // CUDA_IPC_HANDLE_SIZE is 64 bytes
}
