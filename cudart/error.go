package cudart

/*
#include <cuda_runtime_api.h>
*/
import "C"
import "fmt"

type CUDAError int

const (
	ErrInvalidValue        = CUDAError(C.cudaErrorInvalidValue)
	ErrMemoryAllocation    = CUDAError(C.cudaErrorMemoryAllocation)
	ErrInitializationError = CUDAError(C.cudaErrorInitializationError)
	ErrCudartUnloading     = CUDAError(C.cudaErrorCudartUnloading)
)

func cudaErrorToGoError(ce C.cudaError_t) error {
	if ce == C.cudaSuccess {
		return nil
	}
	return CUDAError(ce)
}

func (ce CUDAError) Error() string {
	cErrName := C.cudaGetErrorName(C.cudaError_t(ce))
	cErrStr := C.cudaGetErrorString(C.cudaError_t(ce))
	return fmt.Sprintf("CUDA error %d: %s (%s)", ce, C.GoString(cErrName), C.GoString(cErrStr))
}
