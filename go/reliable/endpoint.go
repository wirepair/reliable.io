package reliable

type Endpoint struct {
	endpointConfig     Config
	numAcks            int
	acks               uint16
	sequence           uint16
	sentPackets        *SequenceBuffer
	recvPackets        *SequenceBuffer
	fragmentReassembly *SequenceBuffer
	//uint64             Counters // ?
}
