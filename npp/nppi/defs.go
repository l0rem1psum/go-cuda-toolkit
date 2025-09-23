package nppi

/*
#include <nppdefs.h>
*/
import "C"
import "unsafe"

type InterpolationMode int

const (
	InterpolationModeUndefined         = InterpolationMode(C.NPPI_INTER_UNDEFINED)
	InterpolationModeNearestNeighbor   = InterpolationMode(C.NPPI_INTER_NN)
	InterpolationModeLinear            = InterpolationMode(C.NPPI_INTER_LINEAR)
	InterpolationModeCubic             = InterpolationMode(C.NPPI_INTER_CUBIC)
	InterpolationModeCubic2PBSpline    = InterpolationMode(C.NPPI_INTER_CUBIC2P_BSPLINE)
	InterpolationModeCubic2PCatmullRom = InterpolationMode(C.NPPI_INTER_CUBIC2P_CATMULLROM)
	InterpolationModeCubic2PB05C03     = InterpolationMode(C.NPPI_INTER_CUBIC2P_B05C03)
	InterpolationModeSuper             = InterpolationMode(C.NPPI_INTER_SUPER)
	InterpolationModeLanczos           = InterpolationMode(C.NPPI_INTER_LANCZOS)
	InterpolationModeLanczos3Advanced  = InterpolationMode(C.NPPI_INTER_LANCZOS3_ADVANCED)
	InterpolationModeSmoothEdge        = InterpolationMode(C.NPPI_SMOOTH_EDGE)
)

type BayerGridPosition int

const (
	BayerGridPositionBGGR = BayerGridPosition(C.NPPI_BAYER_BGGR)
	BayerGridPositionRGGB = BayerGridPosition(C.NPPI_BAYER_RGGB)
	BayerGridPositionGBRG = BayerGridPosition(C.NPPI_BAYER_GBRG)
	BayerGridPositionGRBG = BayerGridPosition(C.NPPI_BAYER_GRBG)
)

type Point struct {
	X, Y int
}

type Size struct {
	Width, Height int
}

func (s Size) asC() C.NppiSize {
	return C.NppiSize{width: C.int(s.Width), height: C.int(s.Height)}
}

type Rect struct {
	X, Y, Width, Height int
}

func (r Rect) asC() C.NppiRect {
	return C.NppiRect{x: C.int(r.X), y: C.int(r.Y), width: C.int(r.Width), height: C.int(r.Height)}
}

type ImageDescriptor struct {
	Data unsafe.Pointer
	Step int
	Size Size
}

func (id ImageDescriptor) asCPtr() *C.NppiImageDescriptor {
	return &C.NppiImageDescriptor{
		pData: id.Data,
		nStep: C.int(id.Step),
		oSize: id.Size.asC(),
	}
}

type WarpAffineBatchCXR struct {
	Src     unsafe.Pointer
	SrcStep int
	Dst     unsafe.Pointer
	DstStep int
	Coeffs  unsafe.Pointer
}

func (wab WarpAffineBatchCXR) asC() C.NppiWarpAffineBatchCXR {
	return C.NppiWarpAffineBatchCXR{
		pSrc:     wab.Src,
		nSrcStep: C.int(wab.SrcStep),
		pDst:     wab.Dst,
		nDstStep: C.int(wab.DstStep),
		pCoeffs:  (*C.Npp64f)(wab.Coeffs),
	}
}
