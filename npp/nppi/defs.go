package nppi

/*
#include <nppdefs.h>
*/
import "C"
import "unsafe"

type Size struct {
	Width, Height int
}

func (s Size) asC() C.NppiSize {
	return C.NppiSize{width: C.int(s.Width), height: C.int(s.Height)}
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
