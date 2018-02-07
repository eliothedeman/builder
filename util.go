package builder

import (
	"reflect"
	"unsafe"
)

func sToB(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var b []byte
	byteHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	byteHeader.Data = strHeader.Data

	l := len(s)
	byteHeader.Len = l
	byteHeader.Cap = l
	return b
}
