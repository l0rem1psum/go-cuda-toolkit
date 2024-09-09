package nppi

/*
#include <nppi_color_conversion.h>
*/
import "C"

import (
	"unsafe"

	"github.com/l0rem1psum/go-cuda-toolkit/npp/internal"
)

// Color Model Conversion

// NppStatus nppiRGBToYUV_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYUV_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYUV_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYUV_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYUV_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func RGBToYUV_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYUV_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func RGBToYUV_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYUV_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV_8u_AC4P4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[4], int nDstStep, NppiSize oSizeROI);
func RGBToYUV_8u_AC4P4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYUV_8u_AC4P4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYUV_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToYUV_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYUV_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYUV_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToYUV_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYUV_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYUV_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func BGRToYUV_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYUV_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYUV_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func BGRToYUV_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYUV_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYUV_8u_AC4P4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[4], int nDstStep, NppiSize oSizeROI);
func BGRToYUV_8u_AC4P4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYUV_8u_AC4P4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToRGB_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUVToRGB_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUVToRGB_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToRGB_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUVToRGB_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUVToRGB_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToRGB_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YUVToRGB_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUVToRGB_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToRGB_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUVToRGB_8u_P3C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUVToRGB_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToRGBBatch_8u_C3R(const NppiImageDescriptor* pSrcBatchList, NppiImageDescriptor* pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YUVToRGBBatch_8u_C3R(pSrcBatchList ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	status := C.nppiYUVToRGBBatch_8u_C3R(
		pSrcBatchList.asCPtr(),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToRGBBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YUVToRGBBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUVToRGBBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToRGBBatch_8u_C3R_Advanced(const NppiImageDescriptor* pSrcBatchList, NppiImageDescriptor* pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YUVToRGBBatch_8u_C3R_Advanced(pSrcBatchList ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	status := C.nppiYUVToRGBBatch_8u_C3R_Advanced(
		pSrcBatchList.asCPtr(),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToRGBBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YUVToRGBBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUVToRGBBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToBGR_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUVToBGR_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUVToBGR_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToBGR_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUVToBGR_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUVToBGR_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToBGR_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YUVToBGR_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUVToBGR_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToBGR_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUVToBGR_8u_P3C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUVToBGR_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToBGRBatch_8u_C3R(const NppiImageDescriptor* pSrcBatchList, NppiImageDescriptor* pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YUVToBGRBatch_8u_C3R(pSrcBatchList ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	status := C.nppiYUVToBGRBatch_8u_C3R(
		pSrcBatchList.asCPtr(),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToBGRBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YUVToBGRBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUVToBGRBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToBGRBatch_8u_C3R_Advanced(const NppiImageDescriptor* pSrcBatchList, NppiImageDescriptor* pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YUVToBGRBatch_8u_C3R_Advanced(pSrcBatchList ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	status := C.nppiYUVToBGRBatch_8u_C3R_Advanced(
		pSrcBatchList.asCPtr(),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUVToBGRBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YUVToBGRBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUVToBGRBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV422_8u_C3C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYUV422_8u_C3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYUV422_8u_C3C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV422_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYUV422_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYUV422_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV422_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYUV422_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYUV422_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV422ToRGB_8u_C2C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUV422ToRGB_8u_C2C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV422ToRGB_8u_C2C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV422ToRGB_8u_P3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YUV422ToRGB_8u_P3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV422ToRGB_8u_P3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV422ToRGB_8u_P3C3R(const Npp8u* const pSrc[3], int rSrcStep[3], Npp8u* pDst, int nDstStep, NppiSize oSizeROI);
func YUV422ToRGB_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV422ToRGB_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV422ToRGB_8u_P3AC4R(const Npp8u* const pSrc[3], int rSrcStep[3], Npp8u* pDst, int nDstStep, NppiSize oSizeROI);
func YUV422ToRGB_8u_P3AC4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV422ToRGB_8u_P3AC4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV422ToRGBBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YUV422ToRGBBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUV422ToRGBBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV422ToRGBBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YUV422ToRGBBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUV422ToRGBBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV422ToBGRBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YUV422ToBGRBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUV422ToBGRBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV422ToBGRBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YUV422ToBGRBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUV422ToBGRBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV420_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYUV420_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYUV420_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYUV420_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYUV420_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYUV420_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToRGB_8u_P3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YUV420ToRGB_8u_P3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV420ToRGB_8u_P3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToRGB_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUV420ToRGB_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV420ToRGB_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToRGB_8u_P3C4R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUV420ToRGB_8u_P3C4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV420ToRGB_8u_P3C4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToRGB_8u_P3AC4R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUV420ToRGB_8u_P3AC4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV420ToRGB_8u_P3AC4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToRGBBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YUV420ToRGBBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUV420ToRGBBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToRGBBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YUV420ToRGBBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUV420ToRGBBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiNV12ToRGB_8u_P2C3R(const Npp8u * const pSrc[2], int rSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func NV12ToRGB_8u_P2C3R(pSrc unsafe.Pointer, rSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiNV12ToRGB_8u_P2C3R(
		(**C.uchar)(pSrc),
		C.int(rSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiNV12ToRGB_709HDTV_8u_P2C3R(const Npp8u * const pSrc[2], int rSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func NV12ToRGB_709HDTV_8u_P2C3R(pSrc unsafe.Pointer, rSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiNV12ToRGB_709HDTV_8u_P2C3R(
		(**C.uchar)(pSrc),
		C.int(rSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiNV12ToRGB_709CSC_8u_P2C3R(const Npp8u * const pSrc[2], int rSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func NV12ToRGB_709CSC_8u_P2C3R(pSrc unsafe.Pointer, rSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiNV12ToRGB_709CSC_8u_P2C3R(
		(**C.uchar)(pSrc),
		C.int(rSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiNV21ToRGB_8u_P2C4R(const Npp8u * const pSrc[2], int rSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func NV21ToRGB_8u_P2C4R(pSrc unsafe.Pointer, rSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiNV21ToRGB_8u_P2C4R(
		(**C.uchar)(pSrc),
		C.int(rSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYUV420_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYUV420_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYUV420_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToBGR_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUV420ToBGR_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV420ToBGR_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToBGR_8u_P3C4R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YUV420ToBGR_8u_P3C4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYUV420ToBGR_8u_P3C4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToBGRBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YUV420ToBGRBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUV420ToBGRBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYUV420ToBGRBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YUV420ToBGRBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYUV420ToBGRBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiNV12ToBGR_8u_P2C3R(const Npp8u * const pSrc[2], int rSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func NV12ToBGR_8u_P2C3R(pSrc unsafe.Pointer, rSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiNV12ToBGR_8u_P2C3R(
		(**C.uchar)(pSrc),
		C.int(rSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiNV12ToBGR_709HDTV_8u_P2C3R(const Npp8u * const pSrc[2], int rSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func NV12ToBGR_709HDTV_8u_P2C3R(pSrc unsafe.Pointer, rSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiNV12ToBGR_709HDTV_8u_P2C3R(
		(**C.uchar)(pSrc),
		C.int(rSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiNV12ToBGR_709CSC_8u_P2C3R(const Npp8u * const pSrc[2], int rSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func NV12ToBGR_709CSC_8u_P2C3R(pSrc unsafe.Pointer, rSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiNV12ToBGR_709CSC_8u_P2C3R(
		(**C.uchar)(pSrc),
		C.int(rSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiNV21ToBGR_8u_P2C4R(const Npp8u * const pSrc[2], int rSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func NV21ToBGR_8u_P2C4R(pSrc unsafe.Pointer, rSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiNV21ToBGR_8u_P2C4R(
		(**C.uchar)(pSrc),
		C.int(rSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYCbCr_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func RGBToYCbCr_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func RGBToYCbCr_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func RGBToYCbCr_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGB_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCrToRGB_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCrToRGB_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGB_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCrToRGB_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCrToRGB_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGB_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCrToRGB_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCrToRGB_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGB_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCrToRGB_8u_P3C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCrToRGB_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGB_8u_P3C4R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst , int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCbCrToRGB_8u_P3C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCbCrToRGB_8u_P3C4R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGBBatch_8u_C3R(const NppiImageDescriptor* pSrcBatchList, NppiImageDescriptor* pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YCbCrToRGBBatch_8u_C3R(pSrcBatchList ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	status := C.nppiYCbCrToRGBBatch_8u_C3R(
		pSrcBatchList.asCPtr(),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGBBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YCbCrToRGBBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCrToRGBBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGBBatch_8u_C3R_Advanced(const NppiImageDescriptor* pSrcBatchList, NppiImageDescriptor* pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YCbCrToRGBBatch_8u_C3R_Advanced(pSrcBatchList ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	status := C.nppiYCbCrToRGBBatch_8u_C3R_Advanced(
		pSrcBatchList.asCPtr(),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToRGBBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YCbCrToRGBBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCrToRGBBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToBGR_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCrToBGR_8u_P3C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCrToBGR_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToBGR_8u_P3C4R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCbCrToBGR_8u_P3C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCbCrToBGR_8u_P3C4R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToBGRBatch_8u_C3R(const NppiImageDescriptor* pSrcBatchList, NppiImageDescriptor* pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YCbCrToBGRBatch_8u_C3R(pSrcBatchList ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	status := C.nppiYCbCrToBGRBatch_8u_C3R(
		pSrcBatchList.asCPtr(),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToBGRBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YCbCrToBGRBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCrToBGRBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToBGRBatch_8u_C3R_Advanced(const NppiImageDescriptor* pSrcBatchList, NppiImageDescriptor* pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YCbCrToBGRBatch_8u_C3R_Advanced(pSrcBatchList ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	status := C.nppiYCbCrToBGRBatch_8u_C3R_Advanced(
		pSrcBatchList.asCPtr(),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToBGRBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YCbCrToBGRBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCrToBGRBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToBGR_709CSC_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCrToBGR_709CSC_8u_P3C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCrToBGR_709CSC_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCrToBGR_709CSC_8u_P3C4R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCbCrToBGR_709CSC_8u_P3C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCbCrToBGR_709CSC_8u_P3C4R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr422_8u_C3C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYCbCr422_8u_C3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr422_8u_C3C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr422_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr422_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr422_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr422_8u_P3C2R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYCbCr422_8u_P3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr422_8u_P3C2R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToRGB_8u_C2C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr422ToRGB_8u_C2C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToRGB_8u_C2C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToRGB_8u_C2P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr422ToRGB_8u_C2P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToRGB_8u_C2P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToRGB_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr422ToRGB_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToRGB_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToRGBBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YCbCr422ToRGBBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCr422ToRGBBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToRGBBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YCbCr422ToRGBBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCr422ToRGBBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCrCb422_8u_C3C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYCrCb422_8u_C3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCrCb422_8u_C3C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCrCb422_8u_P3C2R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYCrCb422_8u_P3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCrCb422_8u_P3C2R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCrCb422ToRGB_8u_C2C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCrCb422ToRGB_8u_C2C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCrCb422ToRGB_8u_C2C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCrCb422ToRGB_8u_C2P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCrCb422ToRGB_8u_C2P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCrCb422ToRGB_8u_C2P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr422_8u_C3C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToYCbCr422_8u_C3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr422_8u_C3C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr422_8u_AC4C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToYCbCr422_8u_AC4C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr422_8u_AC4C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr422_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr422_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr422_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr422_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr422_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr422_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToBGR_8u_C2C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr422ToBGR_8u_C2C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToBGR_8u_C2C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToBGR_8u_C2C4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCbCr422ToBGR_8u_C2C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCbCr422ToBGR_8u_C2C4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToBGR_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr422ToBGR_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToBGR_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToBGRBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YCbCr422ToBGRBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCr422ToBGRBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToBGRBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YCbCr422ToBGRBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCr422ToBGRBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToCbYCr422_8u_C3C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToCbYCr422_8u_C3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToCbYCr422_8u_C3C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToCbYCr422Gamma_8u_C3C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToCbYCr422Gamma_8u_C3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToCbYCr422Gamma_8u_C3C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCbYCr422ToRGB_8u_C2C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func CbYCr422ToRGB_8u_C2C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiCbYCr422ToRGB_8u_C2C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToCbYCr422_8u_AC4C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToCbYCr422_8u_AC4C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToCbYCr422_8u_AC4C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToCbYCr422_709HDTV_8u_C3C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToCbYCr422_709HDTV_8u_C3C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToCbYCr422_709HDTV_8u_C3C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToCbYCr422_709HDTV_8u_AC4C2R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToCbYCr422_709HDTV_8u_AC4C2R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToCbYCr422_709HDTV_8u_AC4C2R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCbYCr422ToBGR_8u_C2C4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func CbYCr422ToBGR_8u_C2C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiCbYCr422ToBGR_8u_C2C4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCbYCr422ToBGR_709HDTV_8u_C2C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func CbYCr422ToBGR_709HDTV_8u_C2C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiCbYCr422ToBGR_709HDTV_8u_C2C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCbYCr422ToBGR_709HDTV_8u_C2C4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func CbYCr422ToBGR_709HDTV_8u_C2C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiCbYCr422ToBGR_709HDTV_8u_C2C4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr420_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr420_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr420_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToRGB_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr420ToRGB_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr420ToRGB_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToRGBBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YCbCr420ToRGBBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCr420ToRGBBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToRGBBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YCbCr420ToRGBBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCr420ToRGBBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCrCb420_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYCrCb420_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCrCb420_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCrCb420ToRGB_8u_P3C4R(const Npp8u * const pSrc[3],int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCrCb420ToRGB_8u_P3C4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCrCb420ToRGB_8u_P3C4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr420_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr420_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr420_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr420_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr420_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr420_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr420_709CSC_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr420_709CSC_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr420_709CSC_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr420_709CSC_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr420_709CSC_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr420_709CSC_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr420_709HDTV_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr420_709HDTV_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr420_709HDTV_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCrCb420_709CSC_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCrCb420_709CSC_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCrCb420_709CSC_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCrCb420_709CSC_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCrCb420_709CSC_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCrCb420_709CSC_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToBGR_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr420ToBGR_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr420ToBGR_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToBGR_8u_P3C4R(const Npp8u * const pSrc[3], int rSrcStep[3],Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCbCr420ToBGR_8u_P3C4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCbCr420ToBGR_8u_P3C4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToBGRBatch_8u_P3C3R(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oSizeROI);
func YCbCr420ToBGRBatch_8u_P3C3R(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCr420ToBGRBatch_8u_P3C3R(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToBGRBatch_8u_P3C3R_Advanced(const NppiImageDescriptor * const pSrcBatchList[3], NppiImageDescriptor * pDstBatchList, int nBatchSize, NppiSize oMaxSizeROI);
func YCbCr420ToBGRBatch_8u_P3C3R_Advanced(pSrcBatchList [3]ImageDescriptor, pDstBatchList ImageDescriptor, nBatchSize int, oMaxSizeROI Size) error {
	srcBatchList := [3]*C.NppiImageDescriptor{pSrcBatchList[0].asCPtr(), pSrcBatchList[1].asCPtr(), pSrcBatchList[2].asCPtr()}
	status := C.nppiYCbCr420ToBGRBatch_8u_P3C3R_Advanced(
		(**C.NppiImageDescriptor)(&srcBatchList[0]),
		pDstBatchList.asCPtr(),
		C.int(nBatchSize),
		oMaxSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToBGR_709CSC_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr420ToBGR_709CSC_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr420ToBGR_709CSC_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToBGR_709HDTV_8u_P3C4R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCbCr420ToBGR_709HDTV_8u_P3C4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCbCr420ToBGR_709HDTV_8u_P3C4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCrCb420_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCrCb420_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCrCb420_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCrCb420_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCrCb420_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCrCb420_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr411_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr411_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr411_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr411_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr411_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr411_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr411_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr411_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr411_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr411_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int rDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr411_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, rDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr411_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&rDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func BGRToYCbCr_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr_8u_AC4P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func BGRToYCbCr_8u_AC4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr_8u_AC4P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr_8u_AC4P4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[4], int nDstStep, NppiSize oSizeROI);
func BGRToYCbCr_8u_AC4P4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr_8u_AC4P4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr411ToBGR_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr411ToBGR_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr411ToBGR_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr411ToBGR_8u_P3C4R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCbCr411ToBGR_8u_P3C4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCbCr411ToBGR_8u_P3C4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr411ToRGB_8u_P3C3R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr411ToRGB_8u_P3C3R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr411ToRGB_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr411ToRGB_8u_P3C4R(const Npp8u * const pSrc[3], int rSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI, Npp8u nAval);
func YCbCr411ToRGB_8u_P3C4R(pSrc unsafe.Pointer, rSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size, nAval uint8) error {
	status := C.nppiYCbCr411ToRGB_8u_P3C4R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&rSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
		C.uchar(nAval),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToXYZ_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToXYZ_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToXYZ_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToXYZ_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToXYZ_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToXYZ_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiXYZToRGB_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func XYZToRGB_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiXYZToRGB_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiXYZToRGB_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func XYZToRGB_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiXYZToRGB_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToLUV_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToLUV_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToLUV_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToLUV_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToLUV_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToLUV_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiLUVToRGB_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func LUVToRGB_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiLUVToRGB_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiLUVToRGB_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func LUVToRGB_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiLUVToRGB_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToLab_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToLab_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToLab_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiLabToBGR_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func LabToBGR_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiLabToBGR_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCC_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYCC_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCC_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCC_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToYCC_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCC_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCCToRGB_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCCToRGB_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCCToRGB_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCCToRGB_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCCToRGB_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCCToRGB_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCCKToCMYK_JPEG_601_8u_P4R(const Npp8u * pSrc[4], int nSrcStep, Npp8u * pDst[4], int nDstStep, NppiSize oSizeROI);
func YCCKToCMYK_JPEG_601_8u_P4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCCKToCMYK_JPEG_601_8u_P4R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCMYKOrYCCKToRGB_JPEG_8u_P4P3R(const Npp8u * pSrc[4], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func CMYKOrYCCKToRGB_JPEG_8u_P4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiCMYKOrYCCKToRGB_JPEG_8u_P4P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCMYKOrYCCKToRGB_JPEG_8u_P4C3R(const Npp8u * pSrc[4], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func CMYKOrYCCKToRGB_JPEG_8u_P4C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiCMYKOrYCCKToRGB_JPEG_8u_P4C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCMYKOrYCCKToBGR_JPEG_8u_P4P3R(const Npp8u * pSrc[4], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func CMYKOrYCCKToBGR_JPEG_8u_P4P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiCMYKOrYCCKToBGR_JPEG_8u_P4P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiCMYKOrYCCKToBGR_JPEG_8u_P4C3R(const Npp8u * pSrc[4], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func CMYKOrYCCKToBGR_JPEG_8u_P4C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiCMYKOrYCCKToBGR_JPEG_8u_P4C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToHLS_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToHLS_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToHLS_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToHLS_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToHLS_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToHLS_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHLSToRGB_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func HLSToRGB_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHLSToRGB_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHLSToRGB_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func HLSToRGB_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHLSToRGB_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToHLS_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToHLS_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToHLS_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToHLS_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func BGRToHLS_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToHLS_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToHLS_8u_AC4P4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[4], int nDstStep, NppiSize oSizeROI);
func BGRToHLS_8u_AC4P4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToHLS_8u_AC4P4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToHLS_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToHLS_8u_P3C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToHLS_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToHLS_8u_AP4C4R(const Npp8u * const pSrc[4], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func BGRToHLS_8u_AP4C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToHLS_8u_AP4C4R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToHLS_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func BGRToHLS_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToHLS_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToHLS_8u_AP4R(const Npp8u * const pSrc[4], int nSrcStep, Npp8u * pDst[4], int nDstStep, NppiSize oSizeROI);
func BGRToHLS_8u_AP4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToHLS_8u_AP4R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHLSToBGR_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func HLSToBGR_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHLSToBGR_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHLSToBGR_8u_AC4P4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[4], int nDstStep, NppiSize oSizeROI);
func HLSToBGR_8u_AC4P4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHLSToBGR_8u_AC4P4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHLSToBGR_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func HLSToBGR_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHLSToBGR_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHLSToBGR_8u_AP4R(const Npp8u * const pSrc[4], int nSrcStep, Npp8u * pDst[4], int nDstStep, NppiSize oSizeROI);
func HLSToBGR_8u_AP4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHLSToBGR_8u_AP4R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHLSToBGR_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func HLSToBGR_8u_P3C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHLSToBGR_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHLSToBGR_8u_AP4C4R(const Npp8u * const pSrc[4], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func HLSToBGR_8u_AP4C4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHLSToBGR_8u_AP4C4R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToHSV_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToHSV_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToHSV_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToHSV_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func RGBToHSV_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToHSV_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHSVToRGB_8u_C3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func HSVToRGB_8u_C3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHSVToRGB_8u_C3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiHSVToRGB_8u_AC4R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func HSVToRGB_8u_AC4R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiHSVToRGB_8u_AC4R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr420_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr420_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr420_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr422_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr422_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr422_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr411_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr411_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr411_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr444_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func RGBToYCbCr444_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr444_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr420_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr420_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr420_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr422_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr422_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr422_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr411_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr411_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr411_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr444_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func BGRToYCbCr444_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr444_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToRGB_JPEG_8u_P3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr420ToRGB_JPEG_8u_P3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr420ToRGB_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToRGB_JPEG_8u_P3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr422ToRGB_JPEG_8u_P3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToRGB_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr411ToRGB_JPEG_8u_P3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr411ToRGB_JPEG_8u_P3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr411ToRGB_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr444ToRGB_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr444ToRGB_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr444ToRGB_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToBGR_JPEG_8u_P3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr420ToBGR_JPEG_8u_P3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr420ToBGR_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToBGR_JPEG_8u_P3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr422ToBGR_JPEG_8u_P3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToBGR_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr411ToBGR_JPEG_8u_P3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr411ToBGR_JPEG_8u_P3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr411ToBGR_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr444ToBGR_JPEG_8u_P3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func YCbCr444ToBGR_JPEG_8u_P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr444ToBGR_JPEG_8u_P3R(
		(**C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr420_JPEG_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr420_JPEG_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr420_JPEG_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr422_JPEG_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr422_JPEG_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr422_JPEG_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr411_JPEG_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func RGBToYCbCr411_JPEG_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr411_JPEG_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiRGBToYCbCr444_JPEG_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func RGBToYCbCr444_JPEG_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiRGBToYCbCr444_JPEG_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr420_JPEG_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr420_JPEG_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr420_JPEG_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr422_JPEG_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr422_JPEG_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr422_JPEG_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr411_JPEG_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int aDstStep[3], NppiSize oSizeROI);
func BGRToYCbCr411_JPEG_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, aDstStep [3]int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr411_JPEG_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		(*C.int)(unsafe.Pointer(&aDstStep[0])),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiBGRToYCbCr444_JPEG_8u_C3P3R(const Npp8u * pSrc, int nSrcStep, Npp8u * pDst[3], int nDstStep, NppiSize oSizeROI);
func BGRToYCbCr444_JPEG_8u_C3P3R(pSrc unsafe.Pointer, nSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiBGRToYCbCr444_JPEG_8u_C3P3R(
		(*C.uchar)(pSrc),
		C.int(nSrcStep),
		(**C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToRGB_JPEG_8u_P3C3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr420ToRGB_JPEG_8u_P3C3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr420ToRGB_JPEG_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToRGB_JPEG_8u_P3C3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr422ToRGB_JPEG_8u_P3C3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToRGB_JPEG_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr411ToRGB_JPEG_8u_P3C3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr411ToRGB_JPEG_8u_P3C3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr411ToRGB_JPEG_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr444ToRGB_JPEG_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr444ToRGB_JPEG_8u_P3C3R(pSrc unsafe.Pointer, aSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr444ToRGB_JPEG_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(aSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr420ToBGR_JPEG_8u_P3C3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr420ToBGR_JPEG_8u_P3C3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr420ToBGR_JPEG_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr422ToBGR_JPEG_8u_P3C3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr422ToBGR_JPEG_8u_P3C3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr422ToBGR_JPEG_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr411ToBGR_JPEG_8u_P3C3R(const Npp8u * const pSrc[3], int aSrcStep[3], Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr411ToBGR_JPEG_8u_P3C3R(pSrc unsafe.Pointer, aSrcStep [3]int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr411ToBGR_JPEG_8u_P3C3R(
		(**C.uchar)(pSrc),
		(*C.int)(unsafe.Pointer(&aSrcStep[0])),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// NppStatus nppiYCbCr444ToBGR_JPEG_8u_P3C3R(const Npp8u * const pSrc[3], int nSrcStep, Npp8u * pDst, int nDstStep, NppiSize oSizeROI);
func YCbCr444ToBGR_JPEG_8u_P3C3R(pSrc unsafe.Pointer, aSrcStep int, pDst unsafe.Pointer, nDstStep int, oSizeROI Size) error {
	status := C.nppiYCbCr444ToBGR_JPEG_8u_P3C3R(
		(**C.uchar)(pSrc),
		C.int(aSrcStep),
		(*C.uchar)(pDst),
		C.int(nDstStep),
		oSizeROI.asC(),
	)
	return internal.StatusToGoError(int(status))
}

// ColorToGray Conversion
