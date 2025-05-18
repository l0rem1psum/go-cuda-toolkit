package cudart

/*
#include <cuda_runtime_api.h>
*/
import "C"
import "unsafe"

func CUDAGetDevice() (int, error) {
	var device C.int
	err := cudaErrorToGoError(C.cudaGetDevice(&device))
	return int(device), err
}

func CUDAGetDeviceCount() (int, error) {
	var count C.int
	err := cudaErrorToGoError(C.cudaGetDeviceCount(&count))
	return int(count), err
}

type CUDADeviceProp struct {
	Name                                   string
	TotalGlobalMem                         uint64
	SharedMemPerBlock                      uint64
	RegsPerBlock                           int
	WarpSize                               int
	MemPitch                               uint64
	MaxThreadsPerBlock                     int
	MaxThreadsDim                          [3]int
	MaxGridSize                            [3]int
	ClockRate                              int
	TotalConstMem                          uint64
	Major                                  int
	Minor                                  int
	TextureAlignment                       uint64
	TexturePitchAlignment                  uint64
	DeviceOverlap                          int
	MultiProcessorCount                    int
	KernelExecTimeoutEnabled               int
	Integrated                             int
	CanMapHostMemory                       int
	ComputeMode                            int
	MaxTexture1D                           int
	MaxTexture1DMipmap                     int
	MaxTexture1DLinear                     int
	MaxTexture2D                           [2]int
	MaxTexture2DMipmap                     [2]int
	MaxTexture2DLinear                     [3]int
	MaxTexture2DGather                     [2]int
	MaxTexture3D                           [3]int
	MaxTexture3DAlt                        [3]int
	MaxTextureCubemap                      int
	MaxTexture1DLayered                    [2]int
	MaxTexture2DLayered                    [3]int
	MaxTextureCubemapLayered               [2]int
	MaxSurface1D                           int
	MaxSurface2D                           [2]int
	MaxSurface3D                           [3]int
	MaxSurface1DLayered                    [2]int
	MaxSurface2DLayered                    [3]int
	MaxSurfaceCubemap                      int
	MaxSurfaceCubemapLayered               [2]int
	SurfaceAlignment                       uint64
	ConcurrentKernels                      int
	ECCEnabled                             int
	PCIBusID                               int
	PCIDeviceID                            int
	PCIDomainID                            int
	TCCDriver                              int
	AsyncEngineCount                       int
	UnifiedAddressing                      int
	MemoryClockRate                        int
	MemoryBusWidth                         int
	L2CacheSize                            int
	PersistingL2CacheMaxSize               int
	MaxThreadsPerMultiProcessor            int
	StreamPrioritiesSupported              int
	GlobalL1CacheSupported                 int
	LocalL1CacheSupported                  int
	SharedMemPerMultiprocessor             uint64
	RegsPerMultiprocessor                  int
	ManagedMemory                          int
	IsMultiGPUBoard                        int
	MultiGPUBoardGroupID                   int
	SingleToDoublePrecisionPerfRatio       int
	PageableMemoryAccess                   int
	ConcurrentManagedAccess                int
	ComputePreemptionSupported             int
	CanUseHostPointerForRegisteredMem      int
	CooperativeLaunch                      int
	CooperativeMultiDeviceLaunch           int
	PageableMemoryAccessUsesHostPageTables int
	DirectManagedMemAccessFromHost         int
	AccessPolicyMaxWindowSize              int
}

func CUDAGetDeviceProperties(device int) (*CUDADeviceProp, error) {
	var prop C.struct_cudaDeviceProp
	err := cudaErrorToGoError(C.cudaGetDeviceProperties(&prop, C.int(device)))
	return &CUDADeviceProp{
		Name:               C.GoString(&prop.name[0]),
		TotalGlobalMem:     uint64(prop.totalGlobalMem),
		SharedMemPerBlock:  uint64(prop.sharedMemPerBlock),
		RegsPerBlock:       int(prop.regsPerBlock),
		WarpSize:           int(prop.warpSize),
		MemPitch:           uint64(prop.memPitch),
		MaxThreadsPerBlock: int(prop.maxThreadsPerBlock),
		MaxThreadsDim: [3]int{
			int(prop.maxThreadsDim[0]),
			int(prop.maxThreadsDim[1]),
			int(prop.maxThreadsDim[2]),
		},
		MaxGridSize: [3]int{
			int(prop.maxGridSize[0]),
			int(prop.maxGridSize[1]),
			int(prop.maxGridSize[2]),
		},
		ClockRate:                int(prop.clockRate),
		TotalConstMem:            uint64(prop.totalConstMem),
		Major:                    int(prop.major),
		Minor:                    int(prop.minor),
		TextureAlignment:         uint64(prop.textureAlignment),
		TexturePitchAlignment:    uint64(prop.texturePitchAlignment),
		DeviceOverlap:            int(prop.deviceOverlap),
		MultiProcessorCount:      int(prop.multiProcessorCount),
		KernelExecTimeoutEnabled: int(prop.kernelExecTimeoutEnabled),
		Integrated:               int(prop.integrated),
		CanMapHostMemory:         int(prop.canMapHostMemory),
		ComputeMode:              int(prop.computeMode),
		MaxTexture1D:             int(prop.maxTexture1D),
		MaxTexture1DMipmap:       int(prop.maxTexture1DMipmap),
		MaxTexture1DLinear:       int(prop.maxTexture1DLinear),
		MaxTexture2D: [2]int{
			int(prop.maxTexture2D[0]),
			int(prop.maxTexture2D[1]),
		},
		MaxTexture2DMipmap: [2]int{
			int(prop.maxTexture2DMipmap[0]),
			int(prop.maxTexture2DMipmap[1]),
		},
		MaxTexture2DLinear: [3]int{
			int(prop.maxTexture2DLinear[0]),
			int(prop.maxTexture2DLinear[1]),
			int(prop.maxTexture2DLinear[2]),
		},
		MaxTexture2DGather: [2]int{
			int(prop.maxTexture2DGather[0]),
			int(prop.maxTexture2DGather[1]),
		},
		MaxTexture3D: [3]int{
			int(prop.maxTexture3D[0]),
			int(prop.maxTexture3D[1]),
			int(prop.maxTexture3D[2]),
		},
		MaxTexture3DAlt: [3]int{
			int(prop.maxTexture3DAlt[0]),
			int(prop.maxTexture3DAlt[1]),
			int(prop.maxTexture3DAlt[2]),
		},
		MaxTextureCubemap: int(prop.maxTextureCubemap),
		MaxTexture1DLayered: [2]int{
			int(prop.maxTexture1DLayered[0]),
			int(prop.maxTexture1DLayered[1]),
		},
		MaxTexture2DLayered: [3]int{
			int(prop.maxTexture2DLayered[0]),
			int(prop.maxTexture2DLayered[1]),
			int(prop.maxTexture2DLayered[2]),
		},
		MaxTextureCubemapLayered: [2]int{
			int(prop.maxTextureCubemapLayered[0]),
			int(prop.maxTextureCubemapLayered[1]),
		},
		MaxSurface1D: int(prop.maxSurface1D),
		MaxSurface2D: [2]int{
			int(prop.maxSurface2D[0]),
			int(prop.maxSurface2D[1]),
		},
		MaxSurface3D: [3]int{
			int(prop.maxSurface3D[0]),
			int(prop.maxSurface3D[1]),
			int(prop.maxSurface3D[2]),
		},
		MaxSurface1DLayered: [2]int{
			int(prop.maxSurface1DLayered[0]),
			int(prop.maxSurface1DLayered[1]),
		},
		MaxSurface2DLayered: [3]int{
			int(prop.maxSurface2DLayered[0]),
			int(prop.maxSurface2DLayered[1]),
			int(prop.maxSurface2DLayered[2]),
		},
		MaxSurfaceCubemap: int(prop.maxSurfaceCubemap),
		MaxSurfaceCubemapLayered: [2]int{
			int(prop.maxSurfaceCubemapLayered[0]),
			int(prop.maxSurfaceCubemapLayered[1]),
		},
		SurfaceAlignment:                       uint64(prop.surfaceAlignment),
		ConcurrentKernels:                      int(prop.concurrentKernels),
		ECCEnabled:                             int(prop.ECCEnabled),
		PCIBusID:                               int(prop.pciBusID),
		PCIDeviceID:                            int(prop.pciDeviceID),
		PCIDomainID:                            int(prop.pciDomainID),
		TCCDriver:                              int(prop.tccDriver),
		AsyncEngineCount:                       int(prop.asyncEngineCount),
		UnifiedAddressing:                      int(prop.unifiedAddressing),
		MemoryClockRate:                        int(prop.memoryClockRate),
		MemoryBusWidth:                         int(prop.memoryBusWidth),
		L2CacheSize:                            int(prop.l2CacheSize),
		PersistingL2CacheMaxSize:               int(prop.persistingL2CacheMaxSize),
		MaxThreadsPerMultiProcessor:            int(prop.maxThreadsPerMultiProcessor),
		StreamPrioritiesSupported:              int(prop.streamPrioritiesSupported),
		GlobalL1CacheSupported:                 int(prop.globalL1CacheSupported),
		LocalL1CacheSupported:                  int(prop.localL1CacheSupported),
		SharedMemPerMultiprocessor:             uint64(prop.sharedMemPerMultiprocessor),
		RegsPerMultiprocessor:                  int(prop.regsPerMultiprocessor),
		ManagedMemory:                          int(prop.managedMemory),
		IsMultiGPUBoard:                        int(prop.isMultiGpuBoard),
		MultiGPUBoardGroupID:                   int(prop.multiGpuBoardGroupID),
		SingleToDoublePrecisionPerfRatio:       int(prop.singleToDoublePrecisionPerfRatio),
		PageableMemoryAccess:                   int(prop.pageableMemoryAccess),
		ConcurrentManagedAccess:                int(prop.concurrentManagedAccess),
		ComputePreemptionSupported:             int(prop.computePreemptionSupported),
		CanUseHostPointerForRegisteredMem:      int(prop.canUseHostPointerForRegisteredMem),
		CooperativeLaunch:                      int(prop.cooperativeLaunch),
		CooperativeMultiDeviceLaunch:           int(prop.cooperativeMultiDeviceLaunch),
		PageableMemoryAccessUsesHostPageTables: int(prop.pageableMemoryAccessUsesHostPageTables),
		DirectManagedMemAccessFromHost:         int(prop.directManagedMemAccessFromHost),
		AccessPolicyMaxWindowSize:              int(prop.accessPolicyMaxWindowSize),
	}, err
}

func CUDASetDevice(device int) error {
	return cudaErrorToGoError(C.cudaSetDevice(C.int(device)))
}

func CUDADeviceReset() error {
	return cudaErrorToGoError(C.cudaDeviceReset())
}

func CUDADeviceSynchronize() error {
	return cudaErrorToGoError(C.cudaDeviceSynchronize())
}

type CUDAIPCMemHandle struct {
	handle C.cudaIpcMemHandle_t
}

func CUDAIPCGetMemHandle(devicePtr unsafe.Pointer) (*CUDAIPCMemHandle, error) {
	var handle C.cudaIpcMemHandle_t
	err := cudaErrorToGoError(C.cudaIpcGetMemHandle(&handle, devicePtr))
	return &CUDAIPCMemHandle{
		handle: handle,
	}, err
}
