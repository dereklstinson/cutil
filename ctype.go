package cutil

//#include <stdbool.h>
//#include <stdio.h>
import "C"
import (
	"unsafe"

	"github.com/dereklstinson/half"
)

//CScalar is used for scalar multiplications with cudnn.  They have to be Ctypes. It could have easily been called voider
type CScalar interface {
	CPtr() unsafe.Pointer
	SIB() uint
}

//CScalartoFloat64 changes a CScalar to a float64 value so it could be read or debugging.
func CScalartoFloat64(x CScalar) float64 {
	switch y := x.(type) {
	case CDouble:
		return float64(y)
	case CFloat:
		return float64(y)
	case CInt:
		return float64(y)
	case CUInt:
		return float64(y)
	case CHalf:
		h := (half.Float16)(y)
		return float64(h)
	case CChar:
		return float64(y)
	case CUChar:
		return float64(y)

	}
	panic("Unsupported val for CScalartoFloat64")

}

//CScalarConversion takes a go type and converts it to a CScalar interface. golang type int and int32 will both be converted to a CInt type.
//If a go type is not supported then it will return a nil.
//Current support is float64,float32,int, int32, int8,uint32, uint, uint8(byte)
func CScalarConversion(gotype interface{}) CScalar {
	switch x := gotype.(type) {
	case float64:
		return CDouble(x)
	case float32:
		return CFloat(x)
	case int:
		return CInt(x)
	case int32:
		return CInt(x)
	case int8:
		return CChar(x)
	case uint8:
		return CUChar(x)
	case uint32:
		return CUInt(x)
	case uint:
		return CUInt(x)
	case half.Float16:
		return CHalf(x)
	case bool:
		return (CBool)(x)
	case CScalar:
		return x
	default:
		return nil
	}
}

//CHalf is a half precision
type CHalf C.ushort

//CPtr returns an unsafe pointer of the half
func (f CHalf) CPtr() unsafe.Pointer { return unsafe.Pointer(&f) }

//SIB returns the number of bytes the CScalar has as an sizeT
func (f CHalf) SIB() uint { return (2) }

//CFloat is a float in C
type CFloat C.float

//CPtr returns an unsafe pointer of the float
func (f CFloat) CPtr() unsafe.Pointer { return unsafe.Pointer(&f) }

//SIB returns the number of bytes the CScalar has
func (f CFloat) SIB() uint { return 4 }

//CDouble is a double in C
type CDouble C.double

//CPtr returns an unsafe pointer of the double
func (d CDouble) CPtr() unsafe.Pointer { return unsafe.Pointer(&d) }

//SIB returns the number of bytes the CScalar has
func (d CDouble) SIB() uint { return 8 }

//CInt is a int in C
type CInt C.int

//CPtr returns an unsafe pointer of the int
func (i CInt) CPtr() unsafe.Pointer { return unsafe.Pointer(&i) }

//SIB returns the number of bytes the CScalar has
func (i CInt) SIB() uint { return 4 }

//CUInt is an unsigned int in C
type CUInt C.uint

//CPtr returns an unsafe pointer of the Unsigned Int
func (i CUInt) CPtr() unsafe.Pointer { return unsafe.Pointer(&i) }

//SIB returns the number of bytes the CScalar has
func (i CUInt) SIB() uint { return 4 }

//CChar is a signed char
type CChar C.char

//CPtr retunrs an unsafe pointer for CInt8
func (c CChar) CPtr() unsafe.Pointer { return unsafe.Pointer(&c) }

//SIB returns the number of bytes the CScalar has
func (c CChar) SIB() uint { return 1 }

//CUChar is a C.uchar
type CUChar C.uchar

//SIB returns the number of bytes the CScalar has
func (c CUChar) SIB() uint { return 1 }

//CPtr retunrs an unsafe pointer for CUInt8
func (c CUChar) CPtr() unsafe.Pointer { return unsafe.Pointer(&c) }

//CBool is a wrapper for C.bool it is in the stdbool.h header
type CBool C.bool

//SIB returns the number of bytes the CScalar has
func (c CBool) SIB() uint { return (uint)(C.sizeof_bool) }

//CPtr retunrs an unsafe pointer for CBool
func (c CBool) CPtr() unsafe.Pointer { return (unsafe.Pointer)(&c) }

//CSizet is a wrapper for C.size_t
type CSizet C.size_t

//SIB returns the number of bytes the CScalar has
func (c CSizet) SIB() uint { return (uint)(C.sizeof_size_t) }

//CPtr retunrs an unsafe pointer for CSizet
func (c CSizet) CPtr() unsafe.Pointer { return (unsafe.Pointer)(&c) }
