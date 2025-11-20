package cudart

/*
#include <cuda_runtime_api.h>

enum {
	cudaHostAllocDefaul_Go = cudaHostAllocDefault,
	cudaHostAllocPortable_Go = cudaHostAllocPortable,
	cudaHostAllocMapped_Go = cudaHostAllocMapped,
	cudaHostAllocWriteCombined_Go = cudaHostAllocWriteCombined,
};
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

type CUDAHostAllocFlag uint32

const (
	CUDAHostAllocDefault       = CUDAHostAllocFlag(C.cudaHostAllocDefaul_Go)
	CUDAHostAllocPortable      = CUDAHostAllocFlag(C.cudaHostAllocPortable_Go)
	CUDAHostAllocMapped        = CUDAHostAllocFlag(C.cudaHostAllocMapped_Go)
	CUDAHostAllocWriteCombined = CUDAHostAllocFlag(C.cudaHostAllocWriteCombined_Go)
)

func CUDAHostAlloc(size uint64, flags CUDAHostAllocFlag) (unsafe.Pointer, error) {
	var hostPtr unsafe.Pointer
	ce := C.cudaHostAlloc(&hostPtr, C.size_t(size), C.uint(flags))
	return hostPtr, cudaErrorToGoError(ce)
}

func CUDAMalloc(size uint64) (unsafe.Pointer, error) {
	var devicePtr unsafe.Pointer
	ce := C.cudaMalloc(&devicePtr, C.size_t(size))
	return devicePtr, cudaErrorToGoError(ce)
}

func CUDAMallocHost(size uint64) (unsafe.Pointer, error) {
	var hostPtr unsafe.Pointer
	ce := C.cudaMallocHost(&hostPtr, C.size_t(size))
	return hostPtr, cudaErrorToGoError(ce)
}

func CUDAFree(devicePtr unsafe.Pointer) error {
	return cudaErrorToGoError(C.cudaFree(devicePtr))
}

func CUDAFreeHost(hostPtr unsafe.Pointer) error {
	return cudaErrorToGoError(C.cudaFreeHost(hostPtr))
}

func CUDAMemGetInfo() (uint32, uint32, error) {
	var free C.size_t
	var total C.size_t
	ce := C.cudaMemGetInfo(&free, &total)
	return uint32(free), uint32(total), cudaErrorToGoError(ce)
}

func CUDAMemcpy(dst unsafe.Pointer, src unsafe.Pointer, count uint64, kind CUDAMemcpyKind) error {
	return cudaErrorToGoError(C.cudaMemcpy(dst, src, C.size_t(count), C.enum_cudaMemcpyKind(kind)))
}

func CUDAMemcpyAsync(dst unsafe.Pointer, src unsafe.Pointer, count uint64, kind CUDAMemcpyKind, stream *CUDAStream) error {
	return cudaErrorToGoError(C.cudaMemcpyAsync(dst, src, C.size_t(count), C.enum_cudaMemcpyKind(kind), stream.s))
}

func CUDAMemset(dst unsafe.Pointer, value int, count uint64) error {
	return cudaErrorToGoError(C.cudaMemset(dst, C.int(value), C.size_t(count)))
}
