//go:build cuda12

package nvjpeg

const (
	ErrIncompleteBitstream = Status(C.NVJPEG_STATUS_INCOMPLETE_BITSTREAM)
)

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

const (
	ChromaSubsampling410V = ChromaSubsampling(C.NVJPEG_CSS_410V)
)

const (
	OutputFormatUnchangedIU16 = OutputFormat(C.NVJPEG_OUTPUT_UNCHANGEDI_U16)
)

const (
	BackendGPUHybridDevice = Backend(C.NVJPEG_BACKEND_GPU_HYBRID_DEVICE)
	BackendHardwareDevice  = Backend(C.NVJPEG_BACKEND_HARDWARE_DEVICE)
	BackendLosslessJPEG    = Backend(C.NVJPEG_BACKEND_LOSSLESS_JPEG)
)

const (
	JpegEncodingLosslessHuffman = JpegEncoding(C.NVJPEG_ENCODING_LOSSLESS_HUFFMAN)
)
