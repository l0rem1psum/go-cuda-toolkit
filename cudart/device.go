package cudart

/*
#cgo pkg-config: cudart-12.1

#include <cuda_runtime_api.h>
*/
import "C"

// https://docs.nvidia.com/cuda/archive/12.1.1/cuda-runtime-api/group__CUDART__DEVICE.html#group__CUDART__DEVICE_1g159587909ffa0791bbe4b40187a4c6bb
func CUDASetDevice(device int) error {
	return cudaErrorToGoError(C.cudaSetDevice(C.int(device)))
}
