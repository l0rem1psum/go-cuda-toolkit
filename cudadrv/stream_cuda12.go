//go:build cuda12

package cudadrv

/*
#include <cuda.h>
*/
import "C"

// CUresult cuStreamGetId ( CUstream hStream, unsigned long long* streamId )
func CUStreamGetId(s *CUstream) (uint64, error) {
	var streamId C.ulonglong
	err := cuResultToGoError(C.cuStreamGetId(s.s, &streamId))
	return uint64(streamId), err
}
