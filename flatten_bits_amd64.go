//+build !noasm
//+build !appengine

package simdjson

import (
	"unsafe"
)

//go:noescape
func _flatten_bits(base_ptr, pbase unsafe.Pointer, idx uint64 /* will be downconverted to uint32 in assembly */, bits uint64)

func flatten_bits(base *[INDEX_SIZE]uint32, base_index *int, idx uint64, bits uint64) (size uint32) {

	_flatten_bits(unsafe.Pointer(&(*base)[0]), unsafe.Pointer(base_index), idx, bits)
	return
}
