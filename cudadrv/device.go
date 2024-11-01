package cudadrv

/*
#include <cuda.h>
*/
import "C"

type CUdevice struct {
	d C.CUdevice
}

// CUresult cuDeviceGet ( CUdevice* device, int  ordinal )
func CUDeviceGet(ordinal int) (CUdevice, error) {
	var device C.CUdevice
	err := cuResultToGoError(C.cuDeviceGet(&device, C.int(ordinal)))
	return CUdevice{device}, err
}

// CUresult cuDeviceGetCount ( int* count )
func CUDeviceGetCount() (int, error) {
	var count C.int
	err := cuResultToGoError(C.cuDeviceGetCount(&count))
	return int(count), err
}
