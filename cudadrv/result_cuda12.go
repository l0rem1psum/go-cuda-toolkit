//go:build cuda12

package cudadrv

/*
#include <cuda.h>
*/
import "C"

const (
	ErrDeviceUnavailable        = CUresult(C.CUDA_ERROR_DEVICE_UNAVAILABLE)
	ErrJITCompilationDisabled   = CUresult(C.CUDA_ERROR_JIT_COMPILATION_DISABLED)
	ErrUnsupportedExecAffinity  = CUresult(C.CUDA_ERROR_UNSUPPORTED_EXEC_AFFINITY)
	ErrUnsupportedDevSideSync   = CUresult(C.CUDA_ERROR_UNSUPPORTED_DEVSIDE_SYNC)
	ErrMPSConnectionFailed      = CUresult(C.CUDA_ERROR_MPS_CONNECTION_FAILED)
	ErrMPSRPCFailure            = CUresult(C.CUDA_ERROR_MPS_RPC_FAILURE)
	ErrMPSServerNotReady        = CUresult(C.CUDA_ERROR_MPS_SERVER_NOT_READY)
	ErrMPSMaxClientsReached     = CUresult(C.CUDA_ERROR_MPS_MAX_CLIENTS_REACHED)
	ErrMPSMaxConnectionsReached = CUresult(C.CUDA_ERROR_MPS_MAX_CONNECTIONS_REACHED)
	ErrMPSClientTerminated      = CUresult(C.CUDA_ERROR_MPS_CLIENT_TERMINATED)
	ErrCDPNotSupported          = CUresult(C.CUDA_ERROR_CDP_NOT_SUPPORTED)
	ErrCDPVersionMismatch       = CUresult(C.CUDA_ERROR_CDP_VERSION_MISMATCH)
	ErrExternalDevice           = CUresult(C.CUDA_ERROR_EXTERNAL_DEVICE)
	ErrInvalidClusterSize       = CUresult(C.CUDA_ERROR_INVALID_CLUSTER_SIZE)
)
