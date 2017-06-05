package reliable

type Reassembly struct {
	sequence             uint16
	ack                  uint16
	ackBits              uint16
	numFragmentsReceived int
	numFragmentsTotal    int
	packetData           []byte
	packetBytes          int
	packetHeaderBytes    int
	fragmentReceived     []byte // [256]?
}
