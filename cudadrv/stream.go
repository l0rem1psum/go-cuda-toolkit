package cudadrv

/*
#include <cuda.h>
*/
import "C"

type CUstream struct {
	s C.CUstream
}

type CUstream_flags uint

const (
	CU_STREAM_DEFAULT      = CUstream_flags(C.CU_STREAM_DEFAULT)
	CU_STREAM_NON_BLOCKING = CUstream_flags(C.CU_STREAM_NON_BLOCKING)
)

// CUresult cuStreamCreate ( CUstream* phStream, unsigned int  Flags )
func CUStreamCreate(flags []CUstream_flags) (*CUstream, error) {
	var s C.CUstream
	err := cuResultToGoError(C.cuStreamCreate(&s, C.uint(combineFlags(flags))))
	return &CUstream{s}, err
}

// CUresult cuStreamCreateWithPriority ( CUstream* phStream, unsigned int  flags, int  priority )
func CUStreamCreateWithPriority(flags []CUstream_flags, priority int) (*CUstream, error) {
	var s C.CUstream
	err := cuResultToGoError(C.cuStreamCreateWithPriority(&s, C.uint(combineFlags(flags)), C.int(priority)))
	return &CUstream{s}, err
}

// CUresult cuStreamDestroy ( CUstream hStream )
func CUStreamDestroy(s *CUstream) error {
	return cuResultToGoError(C.cuStreamDestroy(s.s))
}

// CUresult cuStreamGetCtx ( CUstream hStream, CUcontext* pctx )
func CUStreamGetCtx(s *CUstream) (*CUcontext, error) {
	var ctx C.CUcontext
	err := cuResultToGoError(C.cuStreamGetCtx(s.s, &ctx))
	return &CUcontext{ctx}, err
}

// CUresult cuStreamGetFlags ( CUstream hStream, unsigned int* flags )
func CUStreamGetFlags(s *CUstream) (uint, error) {
	var flags C.uint
	err := cuResultToGoError(C.cuStreamGetFlags(s.s, &flags))
	return uint(flags), err
}

// CUresult cuStreamGetPriority ( CUstream hStream, int* priority )
func CUStreamGetPriority(s *CUstream) (int, error) {
	var priority C.int
	err := cuResultToGoError(C.cuStreamGetPriority(s.s, &priority))
	return int(priority), err
}

// CUresult cuStreamQuery ( CUstream hStream )
func CUStreamQuery(s *CUstream) error {
	return cuResultToGoError(C.cuStreamQuery(s.s))
}

// CUresult cuStreamSynchronize ( CUstream hStream )
func CUStreamSynchronize(s *CUstream) error {
	return cuResultToGoError(C.cuStreamSynchronize(s.s))
}
