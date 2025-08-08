package homework10

import "unsafe"

func Defragment(memory []byte, pointers []unsafe.Pointer) {

	var idx = 0
	var ptr = uintptr(unsafe.Pointer(&memory[idx]))
	for i := 0; i < len(pointers); i++ {
		delta := uintptr(pointers[i]) - ptr
		memory[delta], memory[idx] = 0x00, memory[delta]
		pointers[i] = unsafe.Pointer(&memory[idx])
		idx++
	}
}
