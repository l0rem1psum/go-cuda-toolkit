package cudadrv

/*
#include <cuda.h>
*/
import "C"

type CUcontext struct {
	c C.CUcontext
}

type CUctx_flags uint

const (
	CU_CTX_SCHED_AUTO           = CUctx_flags(C.CU_CTX_SCHED_AUTO)
	CU_CTX_SCHED_SPIN           = CUctx_flags(C.CU_CTX_SCHED_SPIN)
	CU_CTX_SCHED_YIELD          = CUctx_flags(C.CU_CTX_SCHED_YIELD)
	CU_CTX_SCHED_BLOCKING_SYNC  = CUctx_flags(C.CU_CTX_SCHED_BLOCKING_SYNC)
	CU_CTX_BLOCKING_SYNC        = CUctx_flags(C.CU_CTX_BLOCKING_SYNC)
	CU_CTX_SCHED_MASK           = CUctx_flags(C.CU_CTX_SCHED_MASK)
	CU_CTX_MAP_HOST             = CUctx_flags(C.CU_CTX_MAP_HOST)
	CU_CTX_LMEM_RESIZE_TO_MAX   = CUctx_flags(C.CU_CTX_LMEM_RESIZE_TO_MAX)
	CU_CTX_COREDUMP_ENABLE      = CUctx_flags(C.CU_CTX_COREDUMP_ENABLE)
	CU_CTX_USER_COREDUMP_ENABLE = CUctx_flags(C.CU_CTX_USER_COREDUMP_ENABLE)
	CU_CTX_SYNC_MEMOPS          = CUctx_flags(C.CU_CTX_SYNC_MEMOPS)
	CU_CTX_FLAGS_MASK           = CUctx_flags(C.CU_CTX_FLAGS_MASK)
)

// CUresult cuCtxCreate ( CUcontext* pctx, unsigned int  flags, CUdevice dev )
func CUCtxCreate(flags []CUctx_flags, dev *CUdevice) (*CUcontext, error) {
	var ctx C.CUcontext
	err := cuResultToGoError(C.cuCtxCreate(&ctx, C.uint(combineFlags(flags)), dev.d))
	return &CUcontext{ctx}, err
}

// CUresult cuCtxDestroy ( CUcontext ctx )
func CUCtxDestroy(ctx *CUcontext) error {
	return cuResultToGoError(C.cuCtxDestroy(ctx.c))
}

// CUresult cuCtxGetCurrent ( CUcontext* pctx )
func CUCtxGetCurrent() (*CUcontext, error) {
	var ctx C.CUcontext
	err := cuResultToGoError(C.cuCtxGetCurrent(&ctx))
	return &CUcontext{ctx}, err
}

// CUresult cuCtxGetDevice ( CUdevice* device )
func CUCtxGetDevice() (*CUdevice, error) {
	var device C.CUdevice
	err := cuResultToGoError(C.cuCtxGetDevice(&device))
	return &CUdevice{device}, err
}

// CUresult cuCtxGetFlags ( unsigned int* flags )
func CUCtxGetFlags() (uint, error) {
	var flags C.uint
	err := cuResultToGoError(C.cuCtxGetFlags(&flags))
	return uint(flags), err
}

// CUresult cuCtxGetId ( CUcontext ctx, unsigned long long* ctxId )
func CUCtxGetId(ctx *CUcontext) (uint64, error) {
	var ctxId C.ulonglong
	err := cuResultToGoError(C.cuCtxGetId(ctx.c, &ctxId))
	return uint64(ctxId), err
}

// CUresult cuCtxPopCurrent ( CUcontext* pctx )
func CUCtxPopCurrent() (*CUcontext, error) {
	var ctx C.CUcontext
	err := cuResultToGoError(C.cuCtxPopCurrent(&ctx))
	return &CUcontext{ctx}, err
}

// CUresult cuCtxPushCurrent ( CUcontext ctx )
func CUCtxPushCurrent(ctx *CUcontext) error {
	return cuResultToGoError(C.cuCtxPushCurrent(ctx.c))
}

// CUresult cuCtxSetCurrent ( CUcontext ctx )
func CUCtxSetCurrent(ctx *CUcontext) error {
	return cuResultToGoError(C.cuCtxSetCurrent(ctx.c))
}

// CUresult cuCtxSetFlags ( unsigned int  flags )
func CUCtxSetFlags(flags []CUctx_flags) error {
	return cuResultToGoError(C.cuCtxSetFlags(C.uint(combineFlags(flags))))
}

// CUresult cuCtxSynchronize ( void )
func CUCtxSynchronize() error {
	return cuResultToGoError(C.cuCtxSynchronize())
}
