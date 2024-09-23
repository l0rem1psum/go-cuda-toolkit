package internal

/*
#include <nppdefs.h>
*/
import "C"
import "fmt"

type Status int

const (
	ErrNotSupportedMode               = Status(C.NPP_NOT_SUPPORTED_MODE_ERROR)
	ErrInvalidHostPointer             = Status(C.NPP_INVALID_HOST_POINTER_ERROR)
	ErrInvalidDevicePointer           = Status(C.NPP_INVALID_DEVICE_POINTER_ERROR)
	ErrLutPaletteBitsize              = Status(C.NPP_LUT_PALETTE_BITSIZE_ERROR)
	ErrZcModeNotSupported             = Status(C.NPP_ZC_MODE_NOT_SUPPORTED_ERROR)
	ErrNotSufficientComputeCapability = Status(C.NPP_NOT_SUFFICIENT_COMPUTE_CAPABILITY)
	ErrTextureBind                    = Status(C.NPP_TEXTURE_BIND_ERROR)
	ErrWrongIntersectionRoi           = Status(C.NPP_WRONG_INTERSECTION_ROI_ERROR)
	ErrHaarClassifierPixelMatch       = Status(C.NPP_HAAR_CLASSIFIER_PIXEL_MATCH_ERROR)
	ErrMemfree                        = Status(C.NPP_MEMFREE_ERROR)
	ErrMemset                         = Status(C.NPP_MEMSET_ERROR)
	ErrMemcpy                         = Status(C.NPP_MEMCPY_ERROR)
	ErrAlignment                      = Status(C.NPP_ALIGNMENT_ERROR)
	ErrCudaKernelExecution            = Status(C.NPP_CUDA_KERNEL_EXECUTION_ERROR)
	ErrRoundModeNotSupported          = Status(C.NPP_ROUND_MODE_NOT_SUPPORTED_ERROR)
	ErrQualityIndex                   = Status(C.NPP_QUALITY_INDEX_ERROR)
	ErrResizeNoOperation              = Status(C.NPP_RESIZE_NO_OPERATION_ERROR)
	ErrOverflow                       = Status(C.NPP_OVERFLOW_ERROR)
	ErrNotEvenStep                    = Status(C.NPP_NOT_EVEN_STEP_ERROR)
	ErrHistogramNumberOfLevels        = Status(C.NPP_HISTOGRAM_NUMBER_OF_LEVELS_ERROR)
	ErrLutNumberOfLevels              = Status(C.NPP_LUT_NUMBER_OF_LEVELS_ERROR)
	ErrCorruptedData                  = Status(C.NPP_CORRUPTED_DATA_ERROR)
	ErrChannelOrder                   = Status(C.NPP_CHANNEL_ORDER_ERROR)
	ErrZeroMaskValue                  = Status(C.NPP_ZERO_MASK_VALUE_ERROR)
	ErrQuadrangle                     = Status(C.NPP_QUADRANGLE_ERROR)
	ErrRectangle                      = Status(C.NPP_RECTANGLE_ERROR)
	ErrCoefficient                    = Status(C.NPP_COEFFICIENT_ERROR)
	ErrNumberOfChannels               = Status(C.NPP_NUMBER_OF_CHANNELS_ERROR)
	ErrCoi                            = Status(C.NPP_COI_ERROR)
	ErrDivisor                        = Status(C.NPP_DIVISOR_ERROR)
	ErrChannel                        = Status(C.NPP_CHANNEL_ERROR)
	ErrStride                         = Status(C.NPP_STRIDE_ERROR)
	ErrAnchor                         = Status(C.NPP_ANCHOR_ERROR)
	ErrMaskSize                       = Status(C.NPP_MASK_SIZE_ERROR)
	ErrResizeFactor                   = Status(C.NPP_RESIZE_FACTOR_ERROR)
	ErrInterpolation                  = Status(C.NPP_INTERPOLATION_ERROR)
	ErrMirrorFlip                     = Status(C.NPP_MIRROR_FLIP_ERROR)
	ErrMoment00Zero                   = Status(C.NPP_MOMENT_00_ZERO_ERROR)
	ErrThresholdNegativeLevel         = Status(C.NPP_THRESHOLD_NEGATIVE_LEVEL_ERROR)
	ErrThreshold                      = Status(C.NPP_THRESHOLD_ERROR)
	ErrContextMatch                   = Status(C.NPP_CONTEXT_MATCH_ERROR)
	ErrFftFlag                        = Status(C.NPP_FFT_FLAG_ERROR)
	ErrFftOrder                       = Status(C.NPP_FFT_ORDER_ERROR)
	ErrStep                           = Status(C.NPP_STEP_ERROR)
	ErrScaleRange                     = Status(C.NPP_SCALE_RANGE_ERROR)
	ErrDataType                       = Status(C.NPP_DATA_TYPE_ERROR)
	ErrOutOffRange                    = Status(C.NPP_OUT_OFF_RANGE_ERROR)
	ErrDivideByZero                   = Status(C.NPP_DIVIDE_BY_ZERO_ERROR)
	ErrMemoryAllocation               = Status(C.NPP_MEMORY_ALLOCATION_ERR)
	ErrNullPointer                    = Status(C.NPP_NULL_POINTER_ERROR)
	ErrRange                          = Status(C.NPP_RANGE_ERROR)
	ErrSize                           = Status(C.NPP_SIZE_ERROR)
	ErrBadArgument                    = Status(C.NPP_BAD_ARGUMENT_ERROR)
	ErrNoMemory                       = Status(C.NPP_NO_MEMORY_ERROR)
	ErrNotImplemented                 = Status(C.NPP_NOT_IMPLEMENTED_ERROR)
	ErrError                          = Status(C.NPP_ERROR)
	ErrErrorReserved                  = Status(C.NPP_ERROR_RESERVED)
)

const (
	WarningNoOperation           = Status(C.NPP_NO_OPERATION_WARNING)
	WarningDivideByZero          = Status(C.NPP_DIVIDE_BY_ZERO_WARNING)
	WarningAffineQuadIncorrect   = Status(C.NPP_AFFINE_QUAD_INCORRECT_WARNING)
	WarningWrongIntersectionRoi  = Status(C.NPP_WRONG_INTERSECTION_ROI_WARNING)
	WarningWrongIntersectionQuad = Status(C.NPP_WRONG_INTERSECTION_QUAD_WARNING)
	WarningDoubleSize            = Status(C.NPP_DOUBLE_SIZE_WARNING)
	WarningMisalignedDstRoi      = Status(C.NPP_MISALIGNED_DST_ROI_WARNING)
)

func StatusToGoError(ns int) error {
	if ns == C.NPP_NO_ERROR {
		return nil
	}
	return Status(ns)
}

func (s Status) Error() string {
	return fmt.Sprintf("NPP error %d", s)
}
