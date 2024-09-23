package nppi

/*
#include <nppi_support_functions.h>
*/
import "C"
import "unsafe"

// Memory Management

// Npp8u  * nppiMalloc_8u_C1(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_8u_C1(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_8u_C1(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp8u  * nppiMalloc_8u_C2(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_8u_C2(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_8u_C2(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp8u  * nppiMalloc_8u_C3(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_8u_C3(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_8u_C3(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp8u  * nppiMalloc_8u_C4(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_8u_C4(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_8u_C4(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16u * nppiMalloc_16u_C1(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16u_C1(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16u_C1(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16u * nppiMalloc_16u_C2(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16u_C2(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16u_C2(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16u * nppiMalloc_16u_C3(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16u_C3(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16u_C3(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16u * nppiMalloc_16u_C4(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16u_C4(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16u_C4(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16s * nppiMalloc_16s_C1(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16s_C1(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16s_C1(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16s * nppiMalloc_16s_C2(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16s_C2(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16s_C2(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16s * nppiMalloc_16s_C4(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16s_C4(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16s_C4(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16sc * nppiMalloc_16sc_C1(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16sc_C1(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16sc_C1(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16sc * nppiMalloc_16sc_C2(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16sc_C2(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16sc_C2(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16sc * nppiMalloc_16sc_C3(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16sc_C3(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16sc_C3(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp16sc * nppiMalloc_16sc_C4(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_16sc_C4(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_16sc_C4(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32s * nppiMalloc_32s_C1(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32s_C1(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32s_C1(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32s * nppiMalloc_32s_C3(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32s_C3(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32s_C3(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32s * nppiMalloc_32s_C4(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32s_C4(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32s_C4(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32s * nppiMalloc_32sc_C1(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32sc_C1(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32sc_C1(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32s * nppiMalloc_32sc_C2(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32sc_C2(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32sc_C2(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32s * nppiMalloc_32sc_C3(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32sc_C3(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32sc_C3(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32s * nppiMalloc_32sc_C4(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32sc_C4(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32sc_C4(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32f * nppiMalloc_32f_C1(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32f_C1(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32f_C1(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32f * nppiMalloc_32f_C2(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32f_C2(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32f_C2(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32f * nppiMalloc_32f_C3(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32f_C3(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32f_C3(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32f * nppiMalloc_32f_C4(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32f_C4(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32f_C4(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32fc * nppiMalloc_32fc_C1(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32fc_C1(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32fc_C1(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32fc * nppiMalloc_32fc_C2(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32fc_C2(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32fc_C2(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32fc * nppiMalloc_32fc_C3(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32fc_C3(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32fc_C3(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// Npp32fc * nppiMalloc_32fc_C4(int nWidthPixels, int nHeightPixels, int * pStepBytes);
func Malloc_32fc_C4(nWidthPixels int, nHeightPixels int) (unsafe.Pointer, int, error) {
	var pStepBytes C.int
	return unsafe.Pointer(C.nppiMalloc_32fc_C4(C.int(nWidthPixels), C.int(nHeightPixels), &pStepBytes)), int(pStepBytes), nil
}

// void nppiFree(void * pData);
func Free(pData unsafe.Pointer) {
	C.nppiFree(pData)
}
