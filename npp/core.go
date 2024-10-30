package npp

/*
#include <nppcore.h>
*/
import "C"

import (
	"unsafe"

	"github.com/l0rem1psum/go-cuda-toolkit/cudart"
	"github.com/l0rem1psum/go-cuda-toolkit/npp/internal"
)

// int nppGetGpuNumSMs(void);
func GetGpuNumSMs() int {
	return int(C.nppGetGpuNumSMs())
}

// int nppGetMaxThreadsPerBlock(void);
func GetMaxThreadsPerBlock() int {
	return int(C.nppGetMaxThreadsPerBlock())
}

// int nppGetMaxThreadsPerSM(void);
func GetMaxThreadsPerSM() int {
	return int(C.nppGetMaxThreadsPerSM())
}

// cudaStream_t nppGetStream(void);
func GetStream() *cudart.CUDAStream {
	return cudart.NewCUDAStreamFromC(unsafe.Pointer(C.nppGetStream()))
}

// NppStatus nppGetStreamContext(NppStreamContext * pNppStreamContext);
func GetStreamContext() (*StreamContext, error) {
	var sc C.NppStreamContext
	status := C.nppGetStreamContext(&sc)
	return &StreamContext{&sc}, internal.StatusToGoError(int(status))
}

// unsigned int nppGetStreamNumSMs(void);
func GetStreamNumSMs() uint {
	return uint(C.nppGetStreamNumSMs())
}

// unsigned int nppGetStreamMaxThreadsPerSM(void);
func GetStreamMaxThreadsPerSM() uint {
	return uint(C.nppGetStreamMaxThreadsPerSM())
}

// NppStatus nppSetStream(cudaStream_t hStream);
func SetStream(hStream *cudart.CUDAStream) error {
	return internal.StatusToGoError(int(C.nppSetStream(C.cudaStream_t(hStream.C()))))
}
