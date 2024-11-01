package cudadrv

/*
#include <cuda.h>
*/
import "C"

type CUstream struct {
	s C.CUstream
}

// CUresult cuStreamCreate ( CUstream* phStream, unsigned int  Flags )
func CUStreamCreate(flags uint) (*CUstream, error) {
	var s C.CUstream
	err := cuResultToGoError(C.cuStreamCreate(&s, C.uint(flags)))
	return &CUstream{s}, err
}

// CUresult cuStreamCreateWithPriority ( CUstream* phStream, unsigned int  flags, int  priority )
func CUStreamCreateWithPriority(flags uint, priority int) (*CUstream, error) {
	var s C.CUstream
	err := cuResultToGoError(C.cuStreamCreateWithPriority(&s, C.uint(flags), C.int(priority)))
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

// CUresult cuStreamGetId ( CUstream hStream, unsigned long long* streamId )
func CUStreamGetId(s *CUstream) (uint64, error) {
	var streamId C.ulonglong
	err := cuResultToGoError(C.cuStreamGetId(s.s, &streamId))
	return uint64(streamId), err
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
