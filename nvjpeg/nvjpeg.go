package nvjpeg

/*
#include <nvjpeg.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const MAX_COMPONENT = 4

type Status int

const (
	ErrNotInitialized             = Status(C.NVJPEG_STATUS_NOT_INITIALIZED)
	ErrInvalidParameter           = Status(C.NVJPEG_STATUS_INVALID_PARAMETER)
	ErrBadJpeg                    = Status(C.NVJPEG_STATUS_BAD_JPEG)
	ErrJpegNotSupported           = Status(C.NVJPEG_STATUS_JPEG_NOT_SUPPORTED)
	ErrAllocatorFailure           = Status(C.NVJPEG_STATUS_ALLOCATOR_FAILURE)
	ErrExecutionFailed            = Status(C.NVJPEG_STATUS_EXECUTION_FAILED)
	ErrArchMismatch               = Status(C.NVJPEG_STATUS_ARCH_MISMATCH)
	ErrInternalError              = Status(C.NVJPEG_STATUS_INTERNAL_ERROR)
	ErrImplementationNotSupported = Status(C.NVJPEG_STATUS_IMPLEMENTATION_NOT_SUPPORTED)
	ErrIncompleteBitstream        = Status(C.NVJPEG_STATUS_INCOMPLETE_BITSTREAM)
)

func statusToGoError(status C.nvjpegStatus_t) error {
	if status == C.NVJPEG_STATUS_SUCCESS {
		return nil
	}
	return Status(status)
}

func (s Status) Error() string {
	return fmt.Sprintf("nvjpeg error %d", s)
}

type ExifOrientation int

const (
	ExifOrientationUnknown        = ExifOrientation(C.NVJPEG_ORIENTATION_UNKNOWN)
	ExifOrientationNormal         = ExifOrientation(C.NVJPEG_ORIENTATION_NORMAL)
	ExifOrientationFlipHorizontal = ExifOrientation(C.NVJPEG_ORIENTATION_FLIP_HORIZONTAL)
	ExifOrientationRotate180      = ExifOrientation(C.NVJPEG_ORIENTATION_ROTATE_180)
	ExifOrientationFlipVertical   = ExifOrientation(C.NVJPEG_ORIENTATION_FLIP_VERTICAL)
	ExifOrientationTranspose      = ExifOrientation(C.NVJPEG_ORIENTATION_TRANSPOSE)
	ExifOrientationRotate90       = ExifOrientation(C.NVJPEG_ORIENTATION_ROTATE_90)
	ExifOrientationTransverse     = ExifOrientation(C.NVJPEG_ORIENTATION_TRANSVERSE)
	ExifOrientationRotate270      = ExifOrientation(C.NVJPEG_ORIENTATION_ROTATE_270)
)

type ChromaSubsampling int

const (
	ChromaSubsampling444     = ChromaSubsampling(C.NVJPEG_CSS_444)
	ChromaSubsampling422     = ChromaSubsampling(C.NVJPEG_CSS_422)
	ChromaSubsampling420     = ChromaSubsampling(C.NVJPEG_CSS_420)
	ChromaSubsampling440     = ChromaSubsampling(C.NVJPEG_CSS_440)
	ChromaSubsampling411     = ChromaSubsampling(C.NVJPEG_CSS_411)
	ChromaSubsampling410     = ChromaSubsampling(C.NVJPEG_CSS_410)
	ChromaSubsamplingGray    = ChromaSubsampling(C.NVJPEG_CSS_GRAY)
	ChromaSubsampling410V    = ChromaSubsampling(C.NVJPEG_CSS_410V)
	ChromaSubsamplingUnknown = ChromaSubsampling(C.NVJPEG_CSS_UNKNOWN)
)

type OutputFormat int

const (
	OutputFormatUnchanged     = OutputFormat(C.NVJPEG_OUTPUT_UNCHANGED)
	OutputFormatYUV           = OutputFormat(C.NVJPEG_OUTPUT_YUV)
	OutputFormatY             = OutputFormat(C.NVJPEG_OUTPUT_Y)
	OutputFormatRGB           = OutputFormat(C.NVJPEG_OUTPUT_RGB)
	OutputFormatBGR           = OutputFormat(C.NVJPEG_OUTPUT_BGR)
	OutputFormatRGBI          = OutputFormat(C.NVJPEG_OUTPUT_RGBI)
	OutputFormatBGRI          = OutputFormat(C.NVJPEG_OUTPUT_BGRI)
	OutputFormatUnchangedIU16 = OutputFormat(C.NVJPEG_OUTPUT_UNCHANGEDI_U16)
	OutputFormatMax           = OutputFormat(C.NVJPEG_OUTPUT_FORMAT_MAX)
)

type InputFormat int

const (
	InputFormatRGB  = InputFormat(C.NVJPEG_INPUT_RGB)
	InputFormatBGR  = InputFormat(C.NVJPEG_INPUT_BGR)
	InputFormatRGBI = InputFormat(C.NVJPEG_INPUT_RGBI)
	InputFormatBGRI = InputFormat(C.NVJPEG_INPUT_BGRI)
)

type Backend int

const (
	BackendDefault         = Backend(C.NVJPEG_BACKEND_DEFAULT)
	BackendHybrid          = Backend(C.NVJPEG_BACKEND_HYBRID)
	BackendGPUHybrid       = Backend(C.NVJPEG_BACKEND_GPU_HYBRID)
	BackendHardware        = Backend(C.NVJPEG_BACKEND_HARDWARE)
	BackendGPUHybridDevice = Backend(C.NVJPEG_BACKEND_GPU_HYBRID_DEVICE)
	BackendHardwareDevice  = Backend(C.NVJPEG_BACKEND_HARDWARE_DEVICE)
	BackendLosslessJPEG    = Backend(C.NVJPEG_BACKEND_LOSSLESS_JPEG)
)

type JpegEncoding int

const (
	JpegEncodingUnknown               = JpegEncoding(C.NVJPEG_ENCODING_UNKNOWN)
	JpegEncodingBaselineDCT           = JpegEncoding(C.NVJPEG_ENCODING_BASELINE_DCT)
	JpegEncodingExtendedSequentialDCT = JpegEncoding(C.NVJPEG_ENCODING_EXTENDED_SEQUENTIAL_DCT_HUFFMAN)
	JpegEncodingProgressiveDCTHuffman = JpegEncoding(C.NVJPEG_ENCODING_PROGRESSIVE_DCT_HUFFMAN)
	JpegEncodingLosslessHuffman       = JpegEncoding(C.NVJPEG_ENCODING_LOSSLESS_HUFFMAN)
)

type ScaleFactor int

const (
	ScaleFactorNone = ScaleFactor(C.NVJPEG_SCALE_NONE)
	ScaleFactor1By2 = ScaleFactor(C.NVJPEG_SCALE_1_BY_2)
	ScaleFactor1By4 = ScaleFactor(C.NVJPEG_SCALE_1_BY_4)
	ScaleFactor1By8 = ScaleFactor(C.NVJPEG_SCALE_1_BY_8)
)

type Image struct {
	Channel [MAX_COMPONENT]unsafe.Pointer
	Pitch   [MAX_COMPONENT]uint
}

func (i *Image) asC() *C.nvjpegImage_t {
	var cImage C.nvjpegImage_t
	for j := 0; j < MAX_COMPONENT; j++ {
		if i.Channel[j] != nil {
			cImage.channel[j] = (*C.uchar)(i.Channel[j])
		}
		cImage.pitch[j] = C.size_t(i.Pitch[j])
	}
	return &cImage
}

type Handle struct{ h C.nvjpegHandle_t }

type JpegState struct{ s C.nvjpegJpegState_t }

// nvjpegStatus_t NVJPEGAPI nvjpegCreateSimple(nvjpegHandle_t *handle);
func CreateSimple() (*Handle, error) {
	var handle C.nvjpegHandle_t
	if err := statusToGoError(C.nvjpegCreateSimple(&handle)); err != nil {
		return nil, err
	}
	return &Handle{h: handle}, nil
}

// nvjpegStatus_t NVJPEGAPI nvjpegDestroy(nvjpegHandle_t handle);
func Destroy(handle *Handle) error {
	return statusToGoError(C.nvjpegDestroy(handle.h))
}

type EncoderState struct{ es C.nvjpegEncoderState_t }

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderStateCreate(nvjpegHandle_t handle, nvjpegEncoderState_t *encoder_state, cudaStream_t stream);
func EncoderStateCreate(handle *Handle) (*EncoderState, error) {
	var encoderState C.nvjpegEncoderState_t
	if err := statusToGoError(C.nvjpegEncoderStateCreate(handle.h, &encoderState, nil)); err != nil {
		return nil, err
	}
	return &EncoderState{es: encoderState}, nil
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderStateDestroy(nvjpegEncoderState_t encoder_state);
func EncoderStateDestroy(encoderState *EncoderState) error {
	return statusToGoError(C.nvjpegEncoderStateDestroy(encoderState.es))
}

type EncoderParams struct{ ep C.nvjpegEncoderParams_t }

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsCreate(nvjpegHandle_t handle, nvjpegEncoderParams_t *encoder_params, cudaStream_t stream);
func EncoderParamsCreate(handle *Handle) (*EncoderParams, error) {
	var encoderParams C.nvjpegEncoderParams_t
	if err := statusToGoError(C.nvjpegEncoderParamsCreate(handle.h, &encoderParams, nil)); err != nil {
		return nil, err
	}
	return &EncoderParams{ep: encoderParams}, nil
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsDestroy(nvjpegEncoderParams_t encoder_params);
func EncoderParamsDestroy(encoderParams *EncoderParams) error {
	return statusToGoError(C.nvjpegEncoderParamsDestroy(encoderParams.ep))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsSetQuality(nvjpegEncoderParams_t encoder_params, const int quality, cudaStream_t stream);
func EncoderParamsSetQuality(encoderParams *EncoderParams, quality int) error {
	return statusToGoError(C.nvjpegEncoderParamsSetQuality(encoderParams.ep, C.int(quality), nil))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsSetEncoding(nvjpegEncoderParams_t encoder_params, nvjpegJpegEncoding_t etype, cudaStream_t stream);
func EncoderParamsSetEncoding(encoderParams *EncoderParams, etype JpegEncoding) error {
	return statusToGoError(C.nvjpegEncoderParamsSetEncoding(encoderParams.ep, C.nvjpegJpegEncoding_t(etype), nil))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsSetOptimizedHuffman(nvjpegEncoderParams_t encoder_params, const int optimized, cudaStream_t stream);
func EncoderParamsSetOptimizedHuffman(encoderParams *EncoderParams, optimized int) error {
	return statusToGoError(C.nvjpegEncoderParamsSetOptimizedHuffman(encoderParams.ep, C.int(optimized), nil))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsSetSamplingFactors(nvjpegEncoderParams_t encoder_params, const nvjpegChromaSubsampling_t chroma_subsampling, cudaStream_t stream);
func EncoderParamsSetSamplingFactors(encoderParams *EncoderParams, chromaSubsampling ChromaSubsampling) error {
	return statusToGoError(C.nvjpegEncoderParamsSetSamplingFactors(encoderParams.ep, C.nvjpegChromaSubsampling_t(chromaSubsampling), nil))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncodeGetBufferSize(nvjpegHandle_t handle, const nvjpegEncoderParams_t encoder_params, int image_width, int image_height, size_t *max_stream_length);
func EncodeGetBufferSize(handle *Handle, encoderParams *EncoderParams, imageWidth, imageHeight int) (int, error) {
	var maxStreamLength C.size_t
	if err := statusToGoError(C.nvjpegEncodeGetBufferSize(handle.h, encoderParams.ep, C.int(imageWidth), C.int(imageHeight), &maxStreamLength)); err != nil {
		return 0, err
	}
	return int(maxStreamLength), nil
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncodeYUV(nvjpegHandle_t handle, nvjpegEncoderState_t encoder_state, const nvjpegEncoderParams_t encoder_params, const nvjpegImage_t *source, nvjpegChromaSubsampling_t chroma_subsampling, int image_width, int image_height, cudaStream_t stream);
func EncodeYUV(handle *Handle, encoderState *EncoderState, encoderParams *EncoderParams, source *Image, chromaSubsampling ChromaSubsampling, imageWidth, imageHeight int) error {
	return statusToGoError(C.nvjpegEncodeYUV(handle.h, encoderState.es, encoderParams.ep, source.asC(), C.nvjpegChromaSubsampling_t(chromaSubsampling), C.int(imageWidth), C.int(imageHeight), nil))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncodeImage(nvjpegHandle_t handle, nvjpegEncoderState_t encoder_state, const nvjpegEncoderParams_t encoder_params, const nvjpegImage_t *source, nvjpegInputFormat_t input_format, int image_width, int image_height, cudaStream_t stream);
func EncodeImage(handle *Handle, encoderState *EncoderState, encoderParams *EncoderParams, source *Image, inputFormat InputFormat, imageWidth, imageHeight int) error {
	return statusToGoError(C.nvjpegEncodeImage(handle.h, encoderState.es, encoderParams.ep, source.asC(), C.nvjpegInputFormat_t(inputFormat), C.int(imageWidth), C.int(imageHeight), nil))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncodeRetrieveBitstreamDevice(nvjpegHandle_t handle, nvjpegEncoderState_t encoder_state, unsigned char *data, size_t *length, cudaStream_t stream);
func EncodeRetrieveBitstreamDevice(handle *Handle, encoderState *EncoderState, data []byte) (uint, error) {
	cLength := C.size_t(len(data))

	var cData *C.uchar
	if data != nil {
		cData = (*C.uchar)(&data[0])
	}

	err := statusToGoError(C.nvjpegEncodeRetrieveBitstreamDevice(handle.h, encoderState.es, cData, &cLength, nil))

	return uint(cLength), err
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncodeRetrieveBitstream(nvjpegHandle_t handle, nvjpegEncoderState_t encoder_state, unsigned char *data, size_t *length, cudaStream_t stream);
func EncodeRetrieveBitstream(handle *Handle, encoderState *EncoderState, data []byte) (uint, error) {
	cLength := C.size_t(len(data))

	var cData *C.uchar
	if data != nil {
		cData = (*C.uchar)(&data[0])
	}

	err := statusToGoError(C.nvjpegEncodeRetrieveBitstream(handle.h, encoderState.es, cData, &cLength, nil))

	return uint(cLength), err
}
