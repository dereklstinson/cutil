package cutil

import "unsafe"

//Mem is mem that is shared between cuda packages
type Mem interface {
	Pointer
	DPointer
}

//Pointer interface returns an unsafe.Pointer
type Pointer interface {
	Ptr() unsafe.Pointer
}

//DPointer interface returns an *unsafe.Pointer
type DPointer interface {
	DPtr() *unsafe.Pointer
}

//Offset returns a gocu.Mem with the unsafe.Pointer stored in it at the offset bybytes
func Offset(p Pointer, bybytes uint) Pointer {
	return wrapunsafe(unsafe.Pointer(uintptr(p.Ptr()) + uintptr(bybytes)))
}

func wrapunsafe(x unsafe.Pointer) *Wrapper {
	return &Wrapper{
		ptr: x,
	}
}
