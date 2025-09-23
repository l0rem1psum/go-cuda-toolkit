package nppi

/*
#include <nppi_geometry_transforms.h>
*/
import "C"

import (
	"unsafe"

	"github.com/l0rem1psum/go-cuda-toolkit/npp/internal"
)

// ResizeSqrPixel

// Resize

// NppStatus nppiGetResizeTiledSourceOffset(NppiRect oSrcRectROI, NppiRect oDstRectROI, NppiPoint * pNewSrcRectOffset);
func GetResizeTiledSourceOffset(oSrcRectROI Rect, oDstRectROI Rect) (Point, error) {
	var pNewSrcRectOffset C.NppiPoint
	status := C.nppiGetResizeTiledSourceOffset(
		oSrcRectROI.asC(),
		oDstRectROI.asC(),
		&pNewSrcRectOffset,
	)
	return Point{
		X: int(pNewSrcRectOffset.x),
		Y: int(pNewSrcRectOffset.y),
	}, internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_8u_C1R(const Npp8u * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp8u * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_8u_C1R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_8u_C1R(
		(*C.Npp8u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_8u_C3R(const Npp8u * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp8u * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_8u_C3R(
		(*C.Npp8u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_8u_C4R(const Npp8u * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp8u * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_8u_C4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_8u_C4R(
		(*C.Npp8u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_8u_AC4R(const Npp8u * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp8u * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_8u_AC4R(
		(*C.Npp8u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_8u_P3R(const Npp8u * pSrc[3], int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp8u * pDst[3], int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_8u_P3R(
		(**C.Npp8u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(**C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_8u_P4R(const Npp8u * pSrc[4], int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp8u * pDst[4], int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_8u_P4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_8u_P4R(
		(**C.Npp8u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(**C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16u_C1R(const Npp16u * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16u * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16u_C1R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16u_C1R(
		(*C.Npp16u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16u_C3R(const Npp16u * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16u * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16u_C3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16u_C3R(
		(*C.Npp16u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16u_C4R(const Npp16u * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16u * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16u_C4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16u_C4R(
		(*C.Npp16u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16u_AC4R(const Npp16u * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16u * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16u_AC4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16u_AC4R(
		(*C.Npp16u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16u_P3R(const Npp16u * pSrc[3], int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16u * pDst[3], int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16u_P3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16u_P3R(
		(**C.Npp16u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(**C.Npp16u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16u_P4R(const Npp16u * pSrc[4], int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16u * pDst[4], int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16u_P4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16u_P4R(
		(**C.Npp16u)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(**C.Npp16u)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16s_C1R(const Npp16s * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16s * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16s_C1R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16s_C1R(
		(*C.Npp16s)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16s)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16s_C3R(const Npp16s * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16s * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16s_C3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16s_C3R(
		(*C.Npp16s)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16s)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16s_C4R(const Npp16s * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16s * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16s_C4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16s_C4R(
		(*C.Npp16s)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16s)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16s_AC4R(const Npp16s * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16s * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16s_AC4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16s_AC4R(
		(*C.Npp16s)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16s)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16s_P3R(const Npp16s * pSrc[3], int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16s * pDst[3], int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16s_P3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16s_P3R(
		(**C.Npp16s)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(**C.Npp16s)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16s_P4R(const Npp16s * pSrc[4], int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16s * pDst[4], int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16s_P4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16s_P4R(
		(**C.Npp16s)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(**C.Npp16s)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16f_C1R(const Npp16f * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16f * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16f_C1R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16f_C1R(
		(*C.Npp16f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16f_C3R(const Npp16f * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16f * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16f_C3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16f_C3R(
		(*C.Npp16f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_16f_C4R(const Npp16f * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp16f * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_16f_C4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_16f_C4R(
		(*C.Npp16f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp16f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_32f_C1R(const Npp32f * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp32f * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_32f_C1R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_32f_C1R(
		(*C.Npp32f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp32f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_32f_C3R(const Npp32f * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp32f * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_32f_C3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_32f_C3R(
		(*C.Npp32f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp32f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_32f_C4R(const Npp32f * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp32f * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_32f_C4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_32f_C4R(
		(*C.Npp32f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp32f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_32f_AC4R(const Npp32f * pSrc, int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp32f * pDst, int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_32f_AC4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_32f_AC4R(
		(*C.Npp32f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(*C.Npp32f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_32f_P3R(const Npp32f * pSrc[3], int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp32f * pDst[3], int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_32f_P3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_32f_P3R(
		(**C.Npp32f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(**C.Npp32f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiResize_32f_P4R(const Npp32f * pSrc[4], int nSrcStep, NppiSize oSrcSize, NppiRect oSrcRectROI, Npp32f * pDst[4], int nDstStep, NppiSize oDstSize, NppiRect oDstRectROI, int eInterpolation);
func Resize_32f_P4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcRectROI Rect, pDst unsafe.Pointer, nDstStep int, oDstSize Size, oDstRectROI Rect, eInterpolation InterpolationMode) error {
	status := C.nppiResize_32f_P4R(
		(**C.Npp32f)(pSrc),
		C.int(nSrcStep),
		oSrcSize.asC(),
		oSrcRectROI.asC(),
		(**C.Npp32f)(pDst),
		C.int(nDstStep),
		oDstSize.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// Affine Transforms

// NppStatus nppiWarpAffine_8u_C1R(const Npp8u * pSrc, NppiSize oSrcSize, int nSrcStep, NppiRect oSrcROI, Npp8u * pDst, int nDstStep, NppiRect oDstROI, const double aCoeffs[2][3], int eInterpolation);
func WarpAffine_8u_C1R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcROI Rect, pDst unsafe.Pointer, nDstStep int, oDstROI Rect, aCoeffs [2][3]float64, eInterpolation InterpolationMode) error {
	status := C.nppiWarpAffine_8u_C1R(
		(*C.Npp8u)(pSrc),
		oSrcSize.asC(),
		C.int(nSrcStep),
		oSrcROI.asC(),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstROI.asC(),
		(*[3]C.double)(unsafe.Pointer(&aCoeffs[0][0])),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiWarpAffine_8u_C3R(const Npp8u * pSrc, NppiSize oSrcSize, int nSrcStep, NppiRect oSrcROI, Npp8u * pDst, int nDstStep, NppiRect oDstROI, const double aCoeffs[2][3], int eInterpolation);
func WarpAffine_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcROI Rect, pDst unsafe.Pointer, nDstStep int, oDstROI Rect, aCoeffs [2][3]float64, eInterpolation InterpolationMode) error {
	status := C.nppiWarpAffine_8u_C3R(
		(*C.Npp8u)(pSrc),
		oSrcSize.asC(),
		C.int(nSrcStep),
		oSrcROI.asC(),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstROI.asC(),
		(*[3]C.double)(unsafe.Pointer(&aCoeffs[0][0])),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiWarpAffine_8u_C4R(const Npp8u * pSrc, NppiSize oSrcSize, int nSrcStep, NppiRect oSrcROI, Npp8u * pDst, int nDstStep, NppiRect oDstROI, const double aCoeffs[2][3], int eInterpolation);
func WarpAffine_8u_C4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcROI Rect, pDst unsafe.Pointer, nDstStep int, oDstROI Rect, aCoeffs [2][3]float64, eInterpolation InterpolationMode) error {
	status := C.nppiWarpAffine_8u_C4R(
		(*C.Npp8u)(pSrc),
		oSrcSize.asC(),
		C.int(nSrcStep),
		oSrcROI.asC(),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstROI.asC(),
		(*[3]C.double)(unsafe.Pointer(&aCoeffs[0][0])),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiWarpAffine_8u_AC4R(const Npp8u * pSrc, NppiSize oSrcSize, int nSrcStep, NppiRect oSrcROI, Npp8u * pDst, int nDstStep, NppiRect oDstROI, const double aCoeffs[2][3], int eInterpolation);
func WarpAffine_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcROI Rect, pDst unsafe.Pointer, nDstStep int, oDstROI Rect, aCoeffs [2][3]float64, eInterpolation InterpolationMode) error {
	status := C.nppiWarpAffine_8u_AC4R(
		(*C.Npp8u)(pSrc),
		oSrcSize.asC(),
		C.int(nSrcStep),
		oSrcROI.asC(),
		(*C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstROI.asC(),
		(*[3]C.double)(unsafe.Pointer(&aCoeffs[0][0])),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiWarpAffine_8u_P3R(const Npp8u * pSrc[3], NppiSize oSrcSize, int nSrcStep, NppiRect oSrcROI, Npp8u * pDst[3], int nDstStep, NppiRect oDstROI, const double aCoeffs[2][3], int eInterpolation);
func WarpAffine_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcROI Rect, pDst unsafe.Pointer, nDstStep int, oDstROI Rect, aCoeffs [2][3]float64, eInterpolation InterpolationMode) error {
	status := C.nppiWarpAffine_8u_P3R(
		(**C.Npp8u)(pSrc),
		oSrcSize.asC(),
		C.int(nSrcStep),
		oSrcROI.asC(),
		(**C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstROI.asC(),
		(*[3]C.double)(unsafe.Pointer(&aCoeffs[0][0])),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiWarpAffine_8u_P4R(const Npp8u * pSrc[4], NppiSize oSrcSize, int nSrcStep, NppiRect oSrcROI, Npp8u * pDst[4], int nDstStep, NppiRect oDstROI, const double aCoeffs[2][3], int eInterpolation);
func WarpAffine_8u_P4R(pSrc unsafe.Pointer, nSrcStep int, oSrcSize Size, oSrcROI Rect, pDst unsafe.Pointer, nDstStep int, oDstROI Rect, aCoeffs [2][3]float64, eInterpolation InterpolationMode) error {
	status := C.nppiWarpAffine_8u_P4R(
		(**C.Npp8u)(pSrc),
		oSrcSize.asC(),
		C.int(nSrcStep),
		oSrcROI.asC(),
		(**C.Npp8u)(pDst),
		C.int(nDstStep),
		oDstROI.asC(),
		(*[3]C.double)(unsafe.Pointer(&aCoeffs[0][0])),
		C.int(eInterpolation),
	)
	return internal.StatusToGoError(int(status))
}

// Affine Transform Batch

// NppStatus nppiWarpAffineBatchInit(NppiWarpAffineBatchCXR * pBatchList, unsigned int nBatchSize);
func WarpAffineBatchInit(pBatchList []WarpAffineBatchCXR) error {
	batchList := make([]C.NppiWarpAffineBatchCXR, len(pBatchList))
	for i := range pBatchList {
		batchList[i] = pBatchList[i].asC()
	}
	status := C.nppiWarpAffineBatchInit(
		(*C.NppiWarpAffineBatchCXR)(unsafe.Pointer(&pBatchList[0])),
		C.uint(len(pBatchList)),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiWarpAffineBatch_8u_C1R(NppiSize oSmallestSrcSize, NppiRect oSrcRectROI, NppiRect oDstRectROI, int eInterpolation, NppiWarpAffineBatchCXR * pBatchList, unsigned int nBatchSize);
func WarpAffineBatch_8u_C1R(oSmallestSrcSize Size, oSrcRectROI Rect, oDstRectROI Rect, eInterpolation InterpolationMode, pBatchList []WarpAffineBatchCXR) error {
	batchList := make([]C.NppiWarpAffineBatchCXR, len(pBatchList))
	for i := range pBatchList {
		batchList[i] = pBatchList[i].asC()
	}
	status := C.nppiWarpAffineBatch_8u_C1R(
		oSmallestSrcSize.asC(),
		oSrcRectROI.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
		(*C.NppiWarpAffineBatchCXR)(unsafe.Pointer(&pBatchList[0])),
		C.uint(len(pBatchList)),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiWarpAffineBatch_8u_C3R(NppiSize oSmallestSrcSize, NppiRect oSrcRectROI, NppiRect oDstRectROI, int eInterpolation, NppiWarpAffineBatchCXR * pBatchList, unsigned int nBatchSize);
func WarpAffineBatch_8u_C3R(oSmallestSrcSize Size, oSrcRectROI Rect, oDstRectROI Rect, eInterpolation InterpolationMode, pBatchList []WarpAffineBatchCXR) error {
	batchList := make([]C.NppiWarpAffineBatchCXR, len(pBatchList))
	for i := range pBatchList {
		batchList[i] = pBatchList[i].asC()
	}
	status := C.nppiWarpAffineBatch_8u_C3R(
		oSmallestSrcSize.asC(),
		oSrcRectROI.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
		(*C.NppiWarpAffineBatchCXR)(unsafe.Pointer(&pBatchList[0])),
		C.uint(len(pBatchList)),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiWarpAffineBatch_8u_C4R(NppiSize oSmallestSrcSize, NppiRect oSrcRectROI, NppiRect oDstRectROI, int eInterpolation, NppiWarpAffineBatchCXR * pBatchList, unsigned int nBatchSize);
func WarpAffineBatch_8u_C4R(oSmallestSrcSize Size, oSrcRectROI Rect, oDstRectROI Rect, eInterpolation InterpolationMode, pBatchList []WarpAffineBatchCXR) error {
	batchList := make([]C.NppiWarpAffineBatchCXR, len(pBatchList))
	for i := range pBatchList {
		batchList[i] = pBatchList[i].asC()
	}
	status := C.nppiWarpAffineBatch_8u_C4R(
		oSmallestSrcSize.asC(),
		oSrcRectROI.asC(),
		oDstRectROI.asC(),
		C.int(eInterpolation),
		(*C.NppiWarpAffineBatchCXR)(unsafe.Pointer(&pBatchList[0])),
		C.uint(len(pBatchList)),
	)
	return internal.StatusToGoError(int(status))
}
