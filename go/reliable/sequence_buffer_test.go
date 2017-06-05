package reliable

import (
	"testing"
)

func TestIsSequenceGreater(t *testing.T) {
	sequence1 := uint16(32)
	sequence2 := uint16(24)
	if !isSequenceGreater(sequence1, sequence2) {
		t.Fatalf("sequence1 (%d) should be > sequence2 (%d)\n", sequence1, sequence2)
	}

	if isSequenceLess(sequence1, sequence2) {
		t.Fatalf("sequence1 (%d) should be < sequence2 (%d)\n", sequence1, sequence2)
	}
}
