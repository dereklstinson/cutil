# cutil

Helpful interfaces for interfacing go packages that use cgo.

## What this does

This is a small package that contains types, methods and interfaces I seemed to use over and over again to allow go packages that contained c libraries interface with each other easier.

This is not a fix all to go's unexported use of c types.  Each package that uses cgo will still have to do type casting where it is needed.  

## Example of use case from GoCudnn

A brief cudnnActivationForward takes a void pointer for alpha and beta, but they need to be the same type(for the most part) as the tensors x and y.  xD and yD are both tensor descriptors that contain information on the memory of x and y.  The function cscalarbydatatype takes a float64 (alpha and beta) and converts it to a CScalar (a1 and b) of the type that is compatable with cudnnActivationForward().  x and y where allocated by another package, but where able to be used by GoCudnn through the Mem interface.  

````go
func (a *ActivationD) Forward(handle *Handle,
                              alpha float64,
                              xD *TensorD, x cutil.Mem,
                              beta float64,
                              yD *TensorD, y cutil.Mem),
                              error {

                              a1 := cscalarbydatatype(yD.dtype, alpha)
                              b := cscalarbydatatype(yD.dtype, beta)

        return Status(
               C.cudnnActivationForward(handle.x,
                                       a.descriptor,
                                       a1.CPtr(),
                                       xD.descriptor,
                                       x.Ptr(),
                                       b.CPtr(),
                                       yD.descriptor,
                                       y.Ptr(),
                                       )).error("ActivationForward")
}

func cscalarbydatatype(dtype DataType, num float64) ctype.CScalar {
    var x DataType
    switch dtype {
    case x.Double():
        return cutil.CDouble(num)
    case x.Float():
        return cutil.CFloat(num)
    case x.Half():
        y := float32(num)
        return cutil.CFloat(y)
    default:
        return nil
    }

}

````

## Note

I will add to this library where it is needed.
