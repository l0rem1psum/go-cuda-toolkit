package cudadrv

/*
#include <cuda.h>
*/
import "C"
import "fmt"

type (
	CUresult int
)

const (
	ErrInvalidValue                = CUresult(C.CUDA_ERROR_INVALID_VALUE)
	ErrOutOfMemory                 = CUresult(C.CUDA_ERROR_OUT_OF_MEMORY)
	ErrNotInitialized              = CUresult(C.CUDA_ERROR_NOT_INITIALIZED)
	ErrDeinitialized               = CUresult(C.CUDA_ERROR_DEINITIALIZED)
	ErrProfilerDisabled            = CUresult(C.CUDA_ERROR_PROFILER_DISABLED)
	ErrProfilerNotInitialized      = CUresult(C.CUDA_ERROR_PROFILER_NOT_INITIALIZED)
	ErrProfilerAlreadyStarted      = CUresult(C.CUDA_ERROR_PROFILER_ALREADY_STARTED)
	ErrProfilerAlreadyStopped      = CUresult(C.CUDA_ERROR_PROFILER_ALREADY_STOPPED)
	ErrStubLibrary                 = CUresult(C.CUDA_ERROR_STUB_LIBRARY)
	ErrNoDevice                    = CUresult(C.CUDA_ERROR_NO_DEVICE)
	ErrInvalidDevice               = CUresult(C.CUDA_ERROR_INVALID_DEVICE)
	ErrDeviceNotLicensed           = CUresult(C.CUDA_ERROR_DEVICE_NOT_LICENSED)
	ErrInvalidImage                = CUresult(C.CUDA_ERROR_INVALID_IMAGE)
	ErrInvalidContext              = CUresult(C.CUDA_ERROR_INVALID_CONTEXT)
	ErrContextAlreadyCurrent       = CUresult(C.CUDA_ERROR_CONTEXT_ALREADY_CURRENT)
	ErrMapFailed                   = CUresult(C.CUDA_ERROR_MAP_FAILED)
	ErrUnmapFailed                 = CUresult(C.CUDA_ERROR_UNMAP_FAILED)
	ErrArrayIsMapped               = CUresult(C.CUDA_ERROR_ARRAY_IS_MAPPED)
	ErrAlreadyMapped               = CUresult(C.CUDA_ERROR_ALREADY_MAPPED)
	ErrNoBinaryForGPU              = CUresult(C.CUDA_ERROR_NO_BINARY_FOR_GPU)
	ErrAlreadyAcquired             = CUresult(C.CUDA_ERROR_ALREADY_ACQUIRED)
	ErrNotMapped                   = CUresult(C.CUDA_ERROR_NOT_MAPPED)
	ErrNotMappedAsArray            = CUresult(C.CUDA_ERROR_NOT_MAPPED_AS_ARRAY)
	ErrNotMappedAsPointer          = CUresult(C.CUDA_ERROR_NOT_MAPPED_AS_POINTER)
	ErrECCUncorrectable            = CUresult(C.CUDA_ERROR_ECC_UNCORRECTABLE)
	ErrUnsupportedLimit            = CUresult(C.CUDA_ERROR_UNSUPPORTED_LIMIT)
	ErrContextAlreadyInUse         = CUresult(C.CUDA_ERROR_CONTEXT_ALREADY_IN_USE)
	ErrPeerAccessUnsupported       = CUresult(C.CUDA_ERROR_PEER_ACCESS_UNSUPPORTED)
	ErrInvalidPTX                  = CUresult(C.CUDA_ERROR_INVALID_PTX)
	ErrInvalidGraphicsContext      = CUresult(C.CUDA_ERROR_INVALID_GRAPHICS_CONTEXT)
	ErrNVLinkUncorrectable         = CUresult(C.CUDA_ERROR_NVLINK_UNCORRECTABLE)
	ErrJITCompilerNotFound         = CUresult(C.CUDA_ERROR_JIT_COMPILER_NOT_FOUND)
	ErrUnsupportedPTXVersion       = CUresult(C.CUDA_ERROR_UNSUPPORTED_PTX_VERSION)
	ErrInvalidSource               = CUresult(C.CUDA_ERROR_INVALID_SOURCE)
	ErrFileNotFound                = CUresult(C.CUDA_ERROR_FILE_NOT_FOUND)
	ErrSharedObjectSymbol          = CUresult(C.CUDA_ERROR_SHARED_OBJECT_SYMBOL_NOT_FOUND)
	ErrSharedObjectInit            = CUresult(C.CUDA_ERROR_SHARED_OBJECT_INIT_FAILED)
	ErrOperatingSystem             = CUresult(C.CUDA_ERROR_OPERATING_SYSTEM)
	ErrInvalidHandle               = CUresult(C.CUDA_ERROR_INVALID_HANDLE)
	ErrIllegalState                = CUresult(C.CUDA_ERROR_ILLEGAL_STATE)
	ErrNotFound                    = CUresult(C.CUDA_ERROR_NOT_FOUND)
	ErrNotReady                    = CUresult(C.CUDA_ERROR_NOT_READY)
	ErrIllegalAddress              = CUresult(C.CUDA_ERROR_ILLEGAL_ADDRESS)
	ErrLaunchOutOfResources        = CUresult(C.CUDA_ERROR_LAUNCH_OUT_OF_RESOURCES)
	ErrLaunchTimeout               = CUresult(C.CUDA_ERROR_LAUNCH_TIMEOUT)
	ErrLaunchIncompatibleTexturing = CUresult(C.CUDA_ERROR_LAUNCH_INCOMPATIBLE_TEXTURING)
	ErrPeerAccessAlreadyEnabled    = CUresult(C.CUDA_ERROR_PEER_ACCESS_ALREADY_ENABLED)
	ErrPeerAccessNotEnabled        = CUresult(C.CUDA_ERROR_PEER_ACCESS_NOT_ENABLED)
	ErrPrimaryContextActive        = CUresult(C.CUDA_ERROR_PRIMARY_CONTEXT_ACTIVE)
	ErrContextIsDestroyed          = CUresult(C.CUDA_ERROR_CONTEXT_IS_DESTROYED)
	ErrAssert                      = CUresult(C.CUDA_ERROR_ASSERT)
	ErrTooManyPeers                = CUresult(C.CUDA_ERROR_TOO_MANY_PEERS)
	ErrHostMemoryAlreadyRegistered = CUresult(C.CUDA_ERROR_HOST_MEMORY_ALREADY_REGISTERED)
	ErrHostMemoryNotRegistered     = CUresult(C.CUDA_ERROR_HOST_MEMORY_NOT_REGISTERED)
	ErrHardwareStackError          = CUresult(C.CUDA_ERROR_HARDWARE_STACK_ERROR)
	ErrIllegalInstruction          = CUresult(C.CUDA_ERROR_ILLEGAL_INSTRUCTION)
	ErrMisalignedAddress           = CUresult(C.CUDA_ERROR_MISALIGNED_ADDRESS)
	ErrInvalidAddressSpace         = CUresult(C.CUDA_ERROR_INVALID_ADDRESS_SPACE)
	ErrInvalidPC                   = CUresult(C.CUDA_ERROR_INVALID_PC)
	ErrLaunchFailed                = CUresult(C.CUDA_ERROR_LAUNCH_FAILED)
	ErrCooperativeLaunchTooLarge   = CUresult(C.CUDA_ERROR_COOPERATIVE_LAUNCH_TOO_LARGE)
	ErrNotPermitted                = CUresult(C.CUDA_ERROR_NOT_PERMITTED)
	ErrNotSupported                = CUresult(C.CUDA_ERROR_NOT_SUPPORTED)
	ErrSystemNotReady              = CUresult(C.CUDA_ERROR_SYSTEM_NOT_READY)
	ErrSystemDriverMismatch        = CUresult(C.CUDA_ERROR_SYSTEM_DRIVER_MISMATCH)
	ErrCompatNotSupportedOnDevice  = CUresult(C.CUDA_ERROR_COMPAT_NOT_SUPPORTED_ON_DEVICE)
	ErrStreamCaptureUnsupported    = CUresult(C.CUDA_ERROR_STREAM_CAPTURE_UNSUPPORTED)
	ErrStreamCaptureInvalidated    = CUresult(C.CUDA_ERROR_STREAM_CAPTURE_INVALIDATED)
	ErrStreamCaptureMerge          = CUresult(C.CUDA_ERROR_STREAM_CAPTURE_MERGE)
	ErrStreamCaptureUnmatched      = CUresult(C.CUDA_ERROR_STREAM_CAPTURE_UNMATCHED)
	ErrStreamCaptureUnjoined       = CUresult(C.CUDA_ERROR_STREAM_CAPTURE_UNJOINED)
	ErrStreamCaptureIsolation      = CUresult(C.CUDA_ERROR_STREAM_CAPTURE_ISOLATION)
	ErrStreamCaptureImplicit       = CUresult(C.CUDA_ERROR_STREAM_CAPTURE_IMPLICIT)
	ErrCapturedEvent               = CUresult(C.CUDA_ERROR_CAPTURED_EVENT)
	ErrStreamCaptureWrongThread    = CUresult(C.CUDA_ERROR_STREAM_CAPTURE_WRONG_THREAD)
	ErrTimeout                     = CUresult(C.CUDA_ERROR_TIMEOUT)
	ErrGraphExecUpdateFailure      = CUresult(C.CUDA_ERROR_GRAPH_EXEC_UPDATE_FAILURE)
	ErrUnknown                     = CUresult(C.CUDA_ERROR_UNKNOWN)
)

func cuResultToGoError(cr C.CUresult) error {
	if cr == C.CUDA_SUCCESS {
		return nil
	}
	return CUresult(cr)
}

func (cr CUresult) Error() string {
	return fmt.Sprintf("CUDA error %d: %s (%s)", cr, CUGetErrorName(cr), CUGetErrorString(cr))
}

func CUGetErrorName(result CUresult) string {
	return cuGetErrorName(C.CUresult(result))
}

func cuGetErrorName(cResult C.CUresult) string {
	var cErrName *C.char
	if cRes := C.cuGetErrorName(cResult, &cErrName); cRes != C.CUDA_SUCCESS {
		return cuGetErrorName(cRes) // Technically won't cause stack overflow
	}
	return C.GoString(cErrName)
}

func CUGetErrorString(result CUresult) string {
	return cuGetErrorString(C.CUresult(result))
}

func cuGetErrorString(cResult C.CUresult) string {
	var cErrStr *C.char
	if cRes := C.cuGetErrorString(cResult, &cErrStr); cRes != C.CUDA_SUCCESS {
		return cuGetErrorString(cRes) // Technically won't cause stack overflow
	}
	return C.GoString(cErrStr)
}
