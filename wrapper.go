package cutil

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/dereklstinson/half"
)

//Wrapper is a wrapper around some golang memory
type Wrapper struct {
	ptr       unsafe.Pointer
	unitlen   uint
	unitbytes uint
	//	typeflag  int
}

//CreateUnsafeWrapper creates Wrapper from an unsafe.Pointer
func CreateUnsafeWrapper(p unsafe.Pointer, sib uint) *Wrapper {
	return &Wrapper{ptr: p, unitbytes: 1, unitlen: sib}
}

//Ptr satisfies Pointer interface
func (g *Wrapper) Ptr() unsafe.Pointer { return g.ptr }

//DPtr satisfies DPointer interface
func (g *Wrapper) DPtr() *unsafe.Pointer { return &g.ptr }

//SIB returns the size in bytes the wrapper has.
func (g *Wrapper) SIB() (sib uint) {
	return g.unitlen * g.unitbytes

}

/*
//OffSet returns a new GoMem
func (g *Wrapper) OffSet(byunits uint) *Wrapper {

	offset := unsafe.Pointer(uintptr(g.ptr) + uintptr(byunits*g.unitbytes))

	return &Wrapper{
		ptr:       offset,
		unitlen:   g.unitlen - byunits,
		unitbytes: g.unitbytes,
	}
}

//TotalBytes returns the total bytes this has
func (g *Wrapper) TotalBytes() uint {
	return g.unitlen * g.unitbytes

}
*/

//WrapGoMem returns a GoMem considering the input type.
//Will only support slices and pointers to go types
func WrapGoMem(input interface{}) (*Wrapper, error) {
	ptr := new(Wrapper)
	switch val := input.(type) {
	case []int:
		ptr.ptr = unsafe.Pointer(&val[0])
		ptr.unitlen = (uint)(len(val))
		ptr.unitbytes = (uint)(unsafe.Sizeof(val[0]))
		return ptr, nil
	case []int8:
		ptr.ptr = unsafe.Pointer(&val[0])
		ptr.unitlen = (uint)(len(val))
		ptr.unitbytes = (uint)(unsafe.Sizeof(val[0]))
		return ptr, nil
	case []byte:
		ptr.ptr = unsafe.Pointer(&val[0])
		ptr.unitlen = (uint)(len(val))
		ptr.unitbytes = (uint)(unsafe.Sizeof(val[0]))
		return ptr, nil
	case []float64:
		ptr.ptr = unsafe.Pointer(&val[0])
		ptr.unitlen = (uint)(len(val))
		ptr.unitbytes = (uint)(unsafe.Sizeof(val[0]))
		return ptr, nil
	case []uint32:
		ptr.ptr = unsafe.Pointer(&val[0])
		ptr.unitlen = (uint)(len(val))
		ptr.unitbytes = (uint)(unsafe.Sizeof(val[0]))
		return ptr, nil
	case []float32:
		ptr.ptr = unsafe.Pointer(&val[0])
		ptr.unitlen = (uint)(len(val))
		ptr.unitbytes = (uint)(unsafe.Sizeof(val[0]))
		return ptr, nil
	case []int32:
		ptr.ptr = unsafe.Pointer(&val[0])
		ptr.unitlen = (uint)(len(val))
		ptr.unitbytes = (uint)(unsafe.Sizeof(val[0]))
		return ptr, nil
	case []half.Float16:
		ptr.ptr = unsafe.Pointer(&val[0])
		ptr.unitlen = (uint)(len(val))
		ptr.unitbytes = (uint)(unsafe.Sizeof(val[0]))
		return ptr, nil
	case *int:
		ptr.ptr = unsafe.Pointer(val)
		ptr.unitlen = 1
		ptr.unitbytes = (uint)(unsafe.Sizeof(val))
		return ptr, nil
	case *int8:
		ptr.ptr = unsafe.Pointer(val)
		ptr.unitlen = 1
		ptr.unitbytes = (uint)(unsafe.Sizeof(val))
		return ptr, nil
	case *byte:
		ptr.ptr = unsafe.Pointer(val)
		ptr.unitlen = 1
		ptr.unitbytes = (uint)(unsafe.Sizeof(val))
		return ptr, nil
	case *float64:
		ptr.ptr = unsafe.Pointer(val)
		ptr.unitlen = 1
		ptr.unitbytes = (uint)(unsafe.Sizeof(val))
		return ptr, nil
	case *float32:
		ptr.ptr = unsafe.Pointer(val)
		ptr.unitlen = 1
		ptr.unitbytes = (uint)(unsafe.Sizeof(val))
		return ptr, nil
	case *half.Float16:
		ptr.ptr = unsafe.Pointer(val)
		ptr.unitlen = 1
		ptr.unitbytes = (uint)(unsafe.Sizeof(val))
		return ptr, nil
	case *int32:
		ptr.ptr = unsafe.Pointer(val)
		ptr.unitlen = 1
		ptr.unitbytes = (uint)(unsafe.Sizeof(val))
		return ptr, nil
	case *uint32:
		ptr.ptr = unsafe.Pointer(val)
		ptr.unitlen = 1
		ptr.unitbytes = (uint)(unsafe.Sizeof(val))
		return ptr, nil
	default:
		thetype := fmt.Errorf("Type %T", val)
		return nil, errors.New("MakeGoPointer: Unsupported Type -- Type: " + thetype.Error())
	}
}

/*
case *CInt:
	ptr.ptr = unsafe.Pointer(val)
	ptr.unitlen = 1
	ptr.unitbytes = (uint)(unsafe.Sizeof(val))
	return ptr, nil
case *CDouble:
	ptr.ptr = unsafe.Pointer(val)
	ptr.unitlen = 1
	ptr.unitbytes = (uint)(unsafe.Sizeof(val))
	return ptr, nil
case *CFloat:

	ptr.ptr = unsafe.Pointer(val)
	ptr.unitlen = 1
	ptr.unitbytes = (uint)(unsafe.Sizeof(val))
	return ptr, nil
case *CUInt:

	ptr.ptr = val.CPtr()
	ptr.typevalue = "CUInt"
	ptr.size = SizeT(val.Bytes())
	return &ptr, nil
case *CHalf:
	ptr.ptr = val.CPtr()
	ptr.typevalue = "CUInt"
	ptr.size = SizeT(val.Bytes())
	return &ptr, nil

*/
