package cudart

/*
#include <cuda_runtime_api.h>
*/
import "C"

type CUDAStream struct {
	s C.cudaStream_t
}

func CUDAStreamCreate() (*CUDAStream, error) {
	var stream C.cudaStream_t
	ce := C.cudaStreamCreate(&stream)
	return &CUDAStream{stream}, cudaErrorToGoError(ce)
}

func CUDAStreamDestroy(stream *CUDAStream) error {
	return cudaErrorToGoError(C.cudaStreamDestroy(stream.s))
}

func CUDAStreamSynchronize(stream *CUDAStream) error {
	return cudaErrorToGoError(C.cudaStreamSynchronize(stream.s))
}
