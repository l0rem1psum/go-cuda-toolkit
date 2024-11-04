package npp

/*
#include <nppdefs.h>
*/
import "C"

import (
	"github.com/l0rem1psum/go-cuda-toolkit/cudart"
	"github.com/l0rem1psum/go-cuda-toolkit/npp/internal"
)

type Status = internal.Status

type StreamContext struct {
	Stream                            *cudart.CUDAStream
	CudaDeviceId                      int
	MultiProcessorCount               int
	MaxThreadsPerMultiProcessor       int
	MaxThreadsPerBlock                int
	SharedMemPerBlock                 uint64
	CudaDevAttrComputeCapabilityMajor int
	CudaDevAttrComputeCapabilityMinor int
	StreamFlags                       uint
}

func (sc *StreamContext) asC() *C.NppStreamContext {
	return &C.NppStreamContext{
		hStream:                            C.cudaStream_t(sc.Stream.C()),
		nCudaDeviceId:                      C.int(sc.CudaDeviceId),
		nMultiProcessorCount:               C.int(sc.MultiProcessorCount),
		nMaxThreadsPerMultiProcessor:       C.int(sc.MaxThreadsPerMultiProcessor),
		nMaxThreadsPerBlock:                C.int(sc.MaxThreadsPerBlock),
		nSharedMemPerBlock:                 C.ulong(sc.SharedMemPerBlock),
		nCudaDevAttrComputeCapabilityMajor: C.int(sc.CudaDevAttrComputeCapabilityMajor),
		nCudaDevAttrComputeCapabilityMinor: C.int(sc.CudaDevAttrComputeCapabilityMinor),
		nStreamFlags:                       C.uint(sc.StreamFlags),
	}
}
