package cudart

/*
#cgo pkg-config: cudart-12.1

#include <cuda_runtime_api.h>
*/
import "C"
import "unsafe"

// https://docs.nvidia.com/cuda/archive/12.1.1/cuda-runtime-api/group__CUDART__TYPES.html#group__CUDART__TYPES_1g18fa99055ee694244a270e4d5101e95b
type CUDAMemcpyKind int

const (
	CUDAMemcpyHostToHost     = CUDAMemcpyKind(C.cudaMemcpyHostToHost)
	CUDAMemcpyHostToDevice   = CUDAMemcpyKind(C.cudaMemcpyHostToDevice)
	CUDAMemcpyDeviceToHost   = CUDAMemcpyKind(C.cudaMemcpyDeviceToHost)
	CUDAMemcpyDeviceToDevice = CUDAMemcpyKind(C.cudaMemcpyDeviceToDevice)
	CUDAMemcpyDefault        = CUDAMemcpyKind(C.cudaMemcpyDefault)
)

// https://docs.nvidia.com/cuda/archive/12.1.1/cuda-runtime-api/group__CUDART__MEMORY.html#group__CUDART__MEMORY_1g37d37965bfb4803b6d4e59ff26856356
func CUDAMalloc(size uint64) (unsafe.Pointer, error) {
	var devicePtr unsafe.Pointer
	ce := C.cudaMalloc(&devicePtr, C.size_t(size))
	return devicePtr, cudaErrorToGoError(ce)
}

// https://docs.nvidia.com/cuda/archive/12.1.1/cuda-runtime-api/group__CUDART__MEMORY.html#group__CUDART__MEMORY_1ga042655cbbf3408f01061652a075e094
func CUDAFree(devicePtr unsafe.Pointer) error {
	return cudaErrorToGoError(C.cudaFree(devicePtr))
}

// https://docs.nvidia.com/cuda/archive/12.1.1/cuda-runtime-api/group__CUDART__MEMORY.html#group__CUDART__MEMORY_1gc263dbe6574220cc776b45438fc351e8
func CUDAMemcpy(dst unsafe.Pointer, src unsafe.Pointer, count uint64, kind CUDAMemcpyKind) error {
	return cudaErrorToGoError(C.cudaMemcpy(dst, src, C.size_t(count), C.enum_cudaMemcpyKind(kind)))
}
