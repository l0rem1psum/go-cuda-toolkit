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
