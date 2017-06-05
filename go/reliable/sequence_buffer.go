package reliable

type Entry struct {
	data []byte
}

type SequenceBuffer struct {
	sequence       uint16
	numEntries     int
	entryStride    int
	entrySequences []uint32
	entryData      []*Entry
}

func NewSequenceBuffer(numEntries, entryStride int) *SequenceBuffer {
	s := &SequenceBuffer{}
	s.numEntries = numEntries
	s.entryStride = entryStride
	s.Reset()
	return s
}

func (s *SequenceBuffer) Reset() {
	s.entrySequences = make([]uint32, s.numEntries)
	MemsetUint32(s.entrySequences, 0xFFFFFFFF)
	s.entryData = make([]*Entry, s.numEntries*s.entryStride)
}

func (s *SequenceBuffer) RemoveEntries(start, finish int) {

}

func (s *SequenceBuffer) Insert(sequence uint16) {

}

func (s *SequenceBuffer) InsertWithCleanUp(sequence uint16) {

}

func (s *SequenceBuffer) Remove(sequence uint16) {

}

func (s *SequenceBuffer) RemoveWithCleanUp(sequence uint16) {

}

func (s *SequenceBuffer) Available(sequence uint16) bool {
	return false
}

func (s *SequenceBuffer) Find(sequence uint16) *Entry {
	return nil
}

func (s *SequenceBuffer) AtIndex(index int) *Entry {
	return nil
}

func (s *SequenceBuffer) GenerateAckBits() (uint32, uint32) {
	return 0, 0
}

func isSequenceGreater(sequence1, sequence2 uint16) bool {
	return ((sequence1 > sequence2) && (sequence1-sequence2 <= 32768)) ||
		((sequence1 < sequence2) && (sequence2-sequence1 > 32768))
}

func isSequenceLess(sequence1, sequence2 uint16) bool {
	return !isSequenceGreater(sequence1, sequence2)
}
