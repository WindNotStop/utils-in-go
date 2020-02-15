package large_scale_concurrency

import (
	"reflect"
	"unsafe"
)

//[]byte转string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//string转[]byte
func S2B(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	b := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&b))
}
