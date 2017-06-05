package reliable

type Config struct {
	Name                         string // name of the config
	Index                        int
	MaxPacketSize                int
	FragmentAbove                int
	MaxFragments                 int
	FragmentSize                 int
	AckBufferSize                int
	SentPacketsBufferSize        int
	ReceivedPacketsBufferSize    int
	FragmentReassemblyBufferSize int
}
