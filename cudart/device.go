package cudart

/*
#include <cuda_runtime_api.h>
*/
import "C"

func CUDASetDevice(device int) error {
	return cudaErrorToGoError(C.cudaSetDevice(C.int(device)))
}
