package homework11

import "unsafe"

func Trace(stacks [][]uintptr) []uintptr {
	var visited map[uintptr]struct{} = make(map[uintptr]struct{})
	var result []uintptr
	for i := 0; i < len(stacks); i++ {
		for j := 0; j < len(stacks[i]); j++ {
			ptr := stacks[i][j]
			if ptr != uintptr(0x00) {
				mark(ptr, &result, visited)
			}
		}
	}
	return result
}

func mark(ptr uintptr, result *[]uintptr, visited map[uintptr]struct{}) {
	if _, ok := visited[ptr]; ok {
		return
	}
	visited[ptr] = struct{}{}
	pptr := (*uintptr)(unsafe.Pointer(ptr))
	if pptr != nil {
		*result = append(*result, ptr)
		mark(*pptr, result, visited)
	}
}
