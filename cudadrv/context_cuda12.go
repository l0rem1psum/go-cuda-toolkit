//go:build cuda12

package cudadrv

/*
#include <cuda.h>
*/
import "C"

const (
	CU_CTX_COREDUMP_ENABLE      = CUctx_flags(C.CU_CTX_COREDUMP_ENABLE)
	CU_CTX_USER_COREDUMP_ENABLE = CUctx_flags(C.CU_CTX_USER_COREDUMP_ENABLE)
	CU_CTX_SYNC_MEMOPS          = CUctx_flags(C.CU_CTX_SYNC_MEMOPS)
)

// CUresult cuCtxGetId ( CUcontext ctx, unsigned long long* ctxId )
func CUCtxGetId(ctx *CUcontext) (uint64, error) {
	var ctxId C.ulonglong
	err := cuResultToGoError(C.cuCtxGetId(ctx.c, &ctxId))
	return uint64(ctxId), err
}

// CUresult cuCtxSetFlags ( unsigned int  flags )
func CUCtxSetFlags(flags []CUctx_flags) error {
	return cuResultToGoError(C.cuCtxSetFlags(C.uint(combineFlags(flags))))
}
