package homework3

import "unsafe"

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	var refs int = 1
	return COWBuffer{
		data: data,
		refs: &refs,
	}
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++
	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

func (b *COWBuffer) Close() {
	if b.refs != nil && *b.refs > 0 {
		*b.refs--
	}
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}
	if *b.refs > 1 {
		*b.refs--
		var newRefs int = 1
		var newData = make([]byte, len(b.data))
		copy(newData, b.data)
		b.refs = &newRefs
		b.data = newData
	}
	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
