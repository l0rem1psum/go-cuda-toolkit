package npp

/*
#include <nppdefs.h>
*/
import "C"

import (
	"unsafe"

	"github.com/l0rem1psum/go-cuda-toolkit/cudart"
	"github.com/l0rem1psum/go-cuda-toolkit/npp/internal"
)

type Status = internal.Status

type StreamContext struct {
	sc *C.NppStreamContext
}

func NewStreamContext(
	stream *cudart.CUDAStream,
	cudaDeviceId int,
	multiProcessorCount int,
	maxThreadsPerMultiProcessor int,
	maxThreadsPerBlock int,
	sharedMemPerBlock uint64,
	cudaDevAttrComputeCapabilityMajor int,
	cudaDevAttrComputeCapabilityMinor int,
	streamFlags uint,
) *StreamContext {
	sc := C.NppStreamContext{
		hStream:                            C.cudaStream_t(stream.C()),
		nCudaDeviceId:                      C.int(cudaDeviceId),
		nMultiProcessorCount:               C.int(multiProcessorCount),
		nMaxThreadsPerMultiProcessor:       C.int(maxThreadsPerMultiProcessor),
		nMaxThreadsPerBlock:                C.int(maxThreadsPerBlock),
		nSharedMemPerBlock:                 C.ulong(sharedMemPerBlock),
		nCudaDevAttrComputeCapabilityMajor: C.int(cudaDevAttrComputeCapabilityMajor),
		nCudaDevAttrComputeCapabilityMinor: C.int(cudaDevAttrComputeCapabilityMinor),
		nStreamFlags:                       C.uint(streamFlags),
		nReserved0:                         0,
	}
	return &StreamContext{&sc}
}

func (s *StreamContext) C() unsafe.Pointer {
	return unsafe.Pointer(s.sc)
}
