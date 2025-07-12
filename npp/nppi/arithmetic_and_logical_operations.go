package nppi

import (
	"unsafe"

	"github.com/l0rem1psum/go-cuda-toolkit/npp/internal"
)

/*
#include <nppi_arithmetic_and_logical_operations.h>
*/
import "C"

// NppStatus nppiAddC_32f_C3R(const Npp32f *pSrc1, int nSrc1Step, const Npp32f aConstants[3], Npp32f *pDst, int nDstStep, NppiSize oSizeROI)
func AddC_32f_C3R(pSrc1 unsafe.Pointer, nSrc1Step int, aConstants [3]float32, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiAddC_32f_C3R(
		(*C.Npp32f)(pSrc1),
		C.int(nSrc1Step),
		(*C.Npp32f)(unsafe.Pointer(&aConstants[0])),
		(*C.Npp32f)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiAddC_32f_C3IR(const Npp32f aConstants[3], Npp32f *pSrcDst, int nSrcDstStep, NppiSize oSizeROI)
func AddC_32f_C3IR(aConstants [3]float32, pSrcDst unsafe.Pointer, nSrcDstStep int, oSizeROI Size) error {
	status := C.nppiAddC_32f_C3IR(
		(*C.Npp32f)(unsafe.Pointer(&aConstants[0])),
		(*C.Npp32f)(pSrcDst),
		C.int(nSrcDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiSubC_32f_C3R(const Npp32f *pSrc1, int nSrc1Step, const Npp32f aConstants[3], Npp32f *pDst, int nDstStep, NppiSize oSizeROI)
func SubC_32f_C3R(pSrc1 unsafe.Pointer, nSrc1Step int, aConstants [3]float32, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiSubC_32f_C3R(
		(*C.Npp32f)(pSrc1),
		C.int(nSrc1Step),
		(*C.Npp32f)(unsafe.Pointer(&aConstants[0])),
		(*C.Npp32f)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiSubC_32f_C3IR(const Npp32f aConstants[3], Npp32f *pSrcDst, int nSrcDstStep, NppiSize oSizeROI)
func SubC_32f_C3IR(aConstants [3]float32, pSrcDst unsafe.Pointer, nSrcDstStep int, oSizeROI Size) error {
	status := C.nppiSubC_32f_C3IR(
		(*C.Npp32f)(unsafe.Pointer(&aConstants[0])),
		(*C.Npp32f)(pSrcDst),
		C.int(nSrcDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiDivC_32f_C3R(const Npp32f *pSrc1, int nSrc1Step, const Npp32f aConstants[3], Npp32f *pDst, int nDstStep, NppiSize oSizeROI)
func DivC_32f_C3R(pSrc1 unsafe.Pointer, nSrc1Step int, aConstants [3]float32, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiDivC_32f_C3R(
		(*C.Npp32f)(pSrc1),
		C.int(nSrc1Step),
		(*C.Npp32f)(unsafe.Pointer(&aConstants[0])),
		(*C.Npp32f)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiDivC_32f_C3IR(const Npp32f aConstants[3], Npp32f *pSrcDst, int nSrcDstStep, NppiSize oSizeROI)
func DivC_32f_C3IR(aConstants [3]float32, pSrcDst unsafe.Pointer, nSrcDstStep int, oSizeROI Size) error {
	status := C.nppiDivC_32f_C3IR(
		(*C.Npp32f)(unsafe.Pointer(&aConstants[0])),
		(*C.Npp32f)(pSrcDst),
		C.int(nSrcDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}
