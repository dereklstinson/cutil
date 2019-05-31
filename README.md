# ctype

Helpful interface for go packages that use cgo.

## What this does

This is a small package that converts go types to a ctype.
A ctype contains a method to return an unsafe.Pointer.
In order to use it with cgo. It still needs to be type cast.
