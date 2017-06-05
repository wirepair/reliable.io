package reliable

// memsets the in slice to all values of value.
// ref: https://golang.org/src/bytes/bytes.go?s=10462:10501#L395
func MemsetUint32(in []uint32, value uint32) {
	if len(in) == 0 {
		return
	}

	in[0] = value
	for bp := 1; bp < len(in); bp *= 2 {
		copy(in[bp:], in[:bp])
	}
}
