package cudart

/*
#include <cuda_runtime_api.h>
*/
import "C"

func CUDAGetDevice() (int, error) {
	var device C.int
	err := cudaErrorToGoError(C.cudaGetDevice(&device))
	return int(device), err
}

func CUDAGetDeviceCount() (int, error) {
	var count C.int
	err := cudaErrorToGoError(C.cudaGetDeviceCount(&count))
	return int(count), err
}

func CUDASetDevice(device int) error {
	return cudaErrorToGoError(C.cudaSetDevice(C.int(device)))
}

func CUDADeviceReset() error {
	return cudaErrorToGoError(C.cudaDeviceReset())
}

func CUDADeviceSynchronize() error {
	return cudaErrorToGoError(C.cudaDeviceSynchronize())
}
