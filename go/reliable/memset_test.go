package reliable

import (
	"testing"
)

func TestMemsetUint32(t *testing.T) {
	expected := uint32(0xFFFFFFFF)
	in := make([]uint32, 24)

	MemsetUint32(in, expected)
	for i, val := range in {
		if val != expected {
			t.Fatalf("index: %d was not 0xFFFFFFFF got: %x\n", i, val)
		}
	}
}
