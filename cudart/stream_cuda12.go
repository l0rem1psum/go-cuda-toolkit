//go:build cuda12

package cudart

/*
#include <cuda_runtime_api.h>
*/
import "C"

func CUDAStreamGetID(stream *CUDAStream) (uint64, error) {
	var id C.ulonglong
	ce := C.cudaStreamGetId(stream.s, &id)
	return uint64(id), cudaErrorToGoError(ce)
}
