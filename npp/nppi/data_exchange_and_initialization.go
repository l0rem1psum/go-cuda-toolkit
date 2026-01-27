package nppi

/*
#include <nppi_data_exchange_and_initialization.h>
*/
import "C"

import (
	"unsafe"

	"github.com/l0rem1psum/go-cuda-toolkit/npp/internal"
)

// NppStatus nppiConvert_8u32f_C3R(const Npp8u  * pSrc, int nSrcStep, Npp32f * pDst, int nDstStep, NppiSize oSizeROI);
func Convert_8u32f_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiConvert_8u32f_C3R(
		(*C.Npp8u)(pSrc),
		C.int(nSrcStep),
		(*C.Npp32f)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCopy_8u_C3P3R(const Npp8u *pSrc, int nSrcStep, Npp8u *const aDst[3], int nDstStep, NppiSize oSizeROI)
func Copy_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, aDst [3]unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiCopy_8u_C3P3R(
		(*C.Npp8u)(pSrc),
		C.int(nSrcStep),
		(**C.Npp8u)(unsafe.Pointer(&aDst[0])),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCopy_32f_C3P3R(const Npp32f *pSrc, int nSrcStep, Npp32f *const aDst[3], int nDstStep, NppiSize oSizeROI)
func Copy_32f_C3P3R(pSrc unsafe.Pointer, nSrcStep int, aDst [3]unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiCopy_32f_C3P3R(
		(*C.Npp32f)(pSrc),
		C.int(nSrcStep),
		(**C.Npp32f)(unsafe.Pointer(&aDst[0])),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiSwapChannels_8u_C3C4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI, const int aDstOrder[4], const Npp8u nValue);
func SwapChannels_8u_C3C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, aDstOrder [4]int32, nValue uint8) error {
	status := C.nppiSwapChannels_8u_C3C4R(
		(*C.Npp8u)(pSrc),
		C.int(nSrcStep),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		(*C.int)(unsafe.Pointer(&aDstOrder[0])),
		C.Npp8u(nValue),
	)
	return internal.StatusToGoError(int(status))
}
