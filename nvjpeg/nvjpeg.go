package nvjpeg

/*
#include <nvjpeg.h>
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/l0rem1psum/go-cuda-toolkit/cudart"
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

type ChromaSubsampling int

const (
	ChromaSubsampling444     = ChromaSubsampling(C.NVJPEG_CSS_444)
	ChromaSubsampling422     = ChromaSubsampling(C.NVJPEG_CSS_422)
	ChromaSubsampling420     = ChromaSubsampling(C.NVJPEG_CSS_420)
	ChromaSubsampling440     = ChromaSubsampling(C.NVJPEG_CSS_440)
	ChromaSubsampling411     = ChromaSubsampling(C.NVJPEG_CSS_411)
	ChromaSubsampling410     = ChromaSubsampling(C.NVJPEG_CSS_410)
	ChromaSubsamplingGray    = ChromaSubsampling(C.NVJPEG_CSS_GRAY)
	ChromaSubsamplingUnknown = ChromaSubsampling(C.NVJPEG_CSS_UNKNOWN)
)

type OutputFormat int

const (
	OutputFormatUnchanged = OutputFormat(C.NVJPEG_OUTPUT_UNCHANGED)
	OutputFormatYUV       = OutputFormat(C.NVJPEG_OUTPUT_YUV)
	OutputFormatY         = OutputFormat(C.NVJPEG_OUTPUT_Y)
	OutputFormatRGB       = OutputFormat(C.NVJPEG_OUTPUT_RGB)
	OutputFormatBGR       = OutputFormat(C.NVJPEG_OUTPUT_BGR)
	OutputFormatRGBI      = OutputFormat(C.NVJPEG_OUTPUT_RGBI)
	OutputFormatBGRI      = OutputFormat(C.NVJPEG_OUTPUT_BGRI)
	OutputFormatMax       = OutputFormat(C.NVJPEG_OUTPUT_FORMAT_MAX)
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
	BackendDefault   = Backend(C.NVJPEG_BACKEND_DEFAULT)
	BackendHybrid    = Backend(C.NVJPEG_BACKEND_HYBRID)
	BackendGPUHybrid = Backend(C.NVJPEG_BACKEND_GPU_HYBRID)
	BackendHardware  = Backend(C.NVJPEG_BACKEND_HARDWARE)
)

type JpegEncoding int

const (
	JpegEncodingUnknown               = JpegEncoding(C.NVJPEG_ENCODING_UNKNOWN)
	JpegEncodingBaselineDCT           = JpegEncoding(C.NVJPEG_ENCODING_BASELINE_DCT)
	JpegEncodingExtendedSequentialDCT = JpegEncoding(C.NVJPEG_ENCODING_EXTENDED_SEQUENTIAL_DCT_HUFFMAN)
	JpegEncodingProgressiveDCTHuffman = JpegEncoding(C.NVJPEG_ENCODING_PROGRESSIVE_DCT_HUFFMAN)
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

// nvjpegStatus_t NVJPEGAPI nvjpegJpegStateCreate(nvjpegHandle_t handle, nvjpegJpegState_t *jpeg_handle);
func JpegStateCreate(handle *Handle) (*JpegState, error) {
	var jpegState C.nvjpegJpegState_t
	if err := statusToGoError(C.nvjpegJpegStateCreate(handle.h, &jpegState)); err != nil {
		return nil, err
	}
	return &JpegState{s: jpegState}, nil
}

// nvjpegStatus_t NVJPEGAPI nvjpegJpegStateDestroy(nvjpegJpegState_t jpeg_handle);
func JpegStateDestroy(jpegState *JpegState) error {
	return statusToGoError(C.nvjpegJpegStateDestroy(jpegState.s))
}

// nvjpegStatus_t NVJPEGAPI nvjpegGetImageInfo(nvjpegHandle_t handle, const unsigned char *data, size_t length, int *nComponents, nvjpegChromaSubsampling_t *subsampling, int *widths, int *heights);
func GetImageInfo(handle *Handle, data []byte) (int, ChromaSubsampling, []int, []int, error) {
	var cData *C.uchar
	if data != nil {
		cData = (*C.uchar)(&data[0])
	}
	cLength := C.size_t(len(data))

	var nComponents C.int
	var subsampling C.nvjpegChromaSubsampling_t
	var widths [MAX_COMPONENT]C.int
	var heights [MAX_COMPONENT]C.int

	if err := statusToGoError(C.nvjpegGetImageInfo(handle.h, cData, cLength, &nComponents, &subsampling, &widths[0], &heights[0])); err != nil {
		return 0, 0, nil, nil, err
	}

	return int(nComponents), ChromaSubsampling(subsampling), []int{int(widths[0]), int(widths[1]), int(widths[2]), int(widths[3])}, []int{int(heights[0]), int(heights[1]), int(heights[2]), int(heights[3])}, nil
}

// nvjpegStatus_t NVJPEGAPI nvjpegDecode(nvjpegHandle_t handle, nvjpegJpegState_t jpeg_handle, const unsigned char *data, size_t length, nvjpegOutputFormat_t output_format, nvjpegImage_t *destination, cudaStream_t stream);
func Decode(handle *Handle, jpegState *JpegState, data []byte, outputFormat OutputFormat, destination *Image, cudaStream *cudart.CUDAStream) error {
	var cData *C.uchar
	if data != nil {
		cData = (*C.uchar)(&data[0])
	}
	cLength := C.size_t(len(data))

	return statusToGoError(C.nvjpegDecode(handle.h, jpegState.s, cData, cLength, C.nvjpegOutputFormat_t(outputFormat), destination.asC(), C.cudaStream_t(cudaStream.C())))
}

type EncoderState struct{ es C.nvjpegEncoderState_t }

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderStateCreate(nvjpegHandle_t handle, nvjpegEncoderState_t *encoder_state, cudaStream_t stream);
func EncoderStateCreate(handle *Handle, cudaStream *cudart.CUDAStream) (*EncoderState, error) {
	var encoderState C.nvjpegEncoderState_t
	if err := statusToGoError(C.nvjpegEncoderStateCreate(handle.h, &encoderState, C.cudaStream_t(cudaStream.C()))); err != nil {
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
func EncoderParamsCreate(handle *Handle, cudaStream *cudart.CUDAStream) (*EncoderParams, error) {
	var encoderParams C.nvjpegEncoderParams_t
	if err := statusToGoError(C.nvjpegEncoderParamsCreate(handle.h, &encoderParams, C.cudaStream_t(cudaStream.C()))); err != nil {
		return nil, err
	}
	return &EncoderParams{ep: encoderParams}, nil
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsDestroy(nvjpegEncoderParams_t encoder_params);
func EncoderParamsDestroy(encoderParams *EncoderParams) error {
	return statusToGoError(C.nvjpegEncoderParamsDestroy(encoderParams.ep))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsSetQuality(nvjpegEncoderParams_t encoder_params, const int quality, cudaStream_t stream);
func EncoderParamsSetQuality(encoderParams *EncoderParams, quality int, cudaStream *cudart.CUDAStream) error {
	return statusToGoError(C.nvjpegEncoderParamsSetQuality(encoderParams.ep, C.int(quality), C.cudaStream_t(cudaStream.C())))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsSetEncoding(nvjpegEncoderParams_t encoder_params, nvjpegJpegEncoding_t etype, cudaStream_t stream);
func EncoderParamsSetEncoding(encoderParams *EncoderParams, etype JpegEncoding, cudaStream *cudart.CUDAStream) error {
	return statusToGoError(C.nvjpegEncoderParamsSetEncoding(encoderParams.ep, C.nvjpegJpegEncoding_t(etype), C.cudaStream_t(cudaStream.C())))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsSetOptimizedHuffman(nvjpegEncoderParams_t encoder_params, const int optimized, cudaStream_t stream);
func EncoderParamsSetOptimizedHuffman(encoderParams *EncoderParams, optimized int, cudaStream *cudart.CUDAStream) error {
	return statusToGoError(C.nvjpegEncoderParamsSetOptimizedHuffman(encoderParams.ep, C.int(optimized), C.cudaStream_t(cudaStream.C())))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncoderParamsSetSamplingFactors(nvjpegEncoderParams_t encoder_params, const nvjpegChromaSubsampling_t chroma_subsampling, cudaStream_t stream);
func EncoderParamsSetSamplingFactors(encoderParams *EncoderParams, chromaSubsampling ChromaSubsampling, cudaStream *cudart.CUDAStream) error {
	return statusToGoError(C.nvjpegEncoderParamsSetSamplingFactors(encoderParams.ep, C.nvjpegChromaSubsampling_t(chromaSubsampling), C.cudaStream_t(cudaStream.C())))
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
func EncodeYUV(handle *Handle, encoderState *EncoderState, encoderParams *EncoderParams, source *Image, chromaSubsampling ChromaSubsampling, imageWidth, imageHeight int, cudaStream *cudart.CUDAStream) error {
	return statusToGoError(C.nvjpegEncodeYUV(handle.h, encoderState.es, encoderParams.ep, source.asC(), C.nvjpegChromaSubsampling_t(chromaSubsampling), C.int(imageWidth), C.int(imageHeight), C.cudaStream_t(cudaStream.C())))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncodeImage(nvjpegHandle_t handle, nvjpegEncoderState_t encoder_state, const nvjpegEncoderParams_t encoder_params, const nvjpegImage_t *source, nvjpegInputFormat_t input_format, int image_width, int image_height, cudaStream_t stream);
func EncodeImage(handle *Handle, encoderState *EncoderState, encoderParams *EncoderParams, source *Image, inputFormat InputFormat, imageWidth, imageHeight int, cudaStream *cudart.CUDAStream) error {
	return statusToGoError(C.nvjpegEncodeImage(handle.h, encoderState.es, encoderParams.ep, source.asC(), C.nvjpegInputFormat_t(inputFormat), C.int(imageWidth), C.int(imageHeight), C.cudaStream_t(cudaStream.C())))
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncodeRetrieveBitstreamDevice(nvjpegHandle_t handle, nvjpegEncoderState_t encoder_state, unsigned char *data, size_t *length, cudaStream_t stream);
func EncodeRetrieveBitstreamDevice(handle *Handle, encoderState *EncoderState, data []byte, cudaStream *cudart.CUDAStream) (uint, error) {
	cLength := C.size_t(len(data))

	var cData *C.uchar
	if data != nil {
		cData = (*C.uchar)(&data[0])
	}

	err := statusToGoError(C.nvjpegEncodeRetrieveBitstreamDevice(handle.h, encoderState.es, cData, &cLength, C.cudaStream_t(cudaStream.C())))

	return uint(cLength), err
}

// nvjpegStatus_t NVJPEGAPI nvjpegEncodeRetrieveBitstream(nvjpegHandle_t handle, nvjpegEncoderState_t encoder_state, unsigned char *data, size_t *length, cudaStream_t stream);
func EncodeRetrieveBitstream(handle *Handle, encoderState *EncoderState, data []byte, cudaStream *cudart.CUDAStream) (uint, error) {
	cLength := C.size_t(len(data))

	var cData *C.uchar
	if data != nil {
		cData = (*C.uchar)(&data[0])
	}

	err := statusToGoError(C.nvjpegEncodeRetrieveBitstream(handle.h, encoderState.es, cData, &cLength, C.cudaStream_t(cudaStream.C())))

	return uint(cLength), err
}
