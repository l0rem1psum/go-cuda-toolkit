package cudart

/*
#include <cuda_runtime_api.h>
*/
import "C"
import "unsafe"

type CUDAMemcpyKind int

const (
	CUDAMemcpyHostToHost     = CUDAMemcpyKind(C.cudaMemcpyHostToHost)
	CUDAMemcpyHostToDevice   = CUDAMemcpyKind(C.cudaMemcpyHostToDevice)
	CUDAMemcpyDeviceToHost   = CUDAMemcpyKind(C.cudaMemcpyDeviceToHost)
	CUDAMemcpyDeviceToDevice = CUDAMemcpyKind(C.cudaMemcpyDeviceToDevice)
	CUDAMemcpyDefault        = CUDAMemcpyKind(C.cudaMemcpyDefault)
)

func CUDAMalloc(size uint64) (unsafe.Pointer, error) {
	var devicePtr unsafe.Pointer
	ce := C.cudaMalloc(&devicePtr, C.size_t(size))
	return devicePtr, cudaErrorToGoError(ce)
}

func CUDAFree(devicePtr unsafe.Pointer) error {
	return cudaErrorToGoError(C.cudaFree(devicePtr))
}

func CUDAMemcpy(dst unsafe.Pointer, src unsafe.Pointer, count uint64, kind CUDAMemcpyKind) error {
	return cudaErrorToGoError(C.cudaMemcpy(dst, src, C.size_t(count), C.enum_cudaMemcpyKind(kind)))
}

func CUDAMemcpyAsync(dst unsafe.Pointer, src unsafe.Pointer, count uint64, kind CUDAMemcpyKind, stream *CUDAStream) error {
	return cudaErrorToGoError(C.cudaMemcpyAsync(dst, src, C.size_t(count), C.enum_cudaMemcpyKind(kind), stream.s))
}
