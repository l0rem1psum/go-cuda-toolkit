package nppi

/*
#include <nppi_color_conversion.h>
*/
import "C"
import "unsafe"

func unsafeArr2ToNpp8uArr2(arr [2]unsafe.Pointer) **C.Npp8u {
	arr2 := [2]*C.Npp8u{
		(*C.Npp8u)(arr[0]),
		(*C.Npp8u)(arr[1]),
	}
	return &arr2[0]
}

func unsafeArr3ToNpp8uArr3(arr [3]unsafe.Pointer) **C.Npp8u {
	arr3 := [3]*C.Npp8u{
		(*C.Npp8u)(arr[0]),
		(*C.Npp8u)(arr[1]),
		(*C.Npp8u)(arr[2]),
	}
	return &arr3[0]
}

func unsafeArr4ToNpp8uArr4(arr [4]unsafe.Pointer) **C.Npp8u {
	arr4 := [4]*C.Npp8u{
		(*C.Npp8u)(arr[0]),
		(*C.Npp8u)(arr[1]),
		(*C.Npp8u)(arr[2]),
		(*C.Npp8u)(arr[3]),
	}
	return &arr4[0]
}
