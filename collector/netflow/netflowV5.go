package netflow

type HeaderV5 struct {
	version     		uint16 // NetFlow export format version number
	count				uint16 // Number of flows exported in this packet (1-30)
	sysUptime			uint16 // Current time in milliseconds since the export device booted
	unixSecs			uint16 // Current count of seconds since 0000 UTC 1970
	unixNsecs			uint16 // Residual nanoseconds since 0000 UTC 1970
	flowSequence		uint16 // Sequence counter of total flows seen
	engineType			uint8 // Type of flow-switching engine
	engineId			uint8 // Slot number of the flow-switching engine
	samplingInterval	uint16 // First two bits hold the sampling mode; remaining 14 bits hold value of sampling interval
}

type FlowRecordV5 struct {
	srcAddr		uint16 // Source IP address
	dstAddr		uint16 // Destination IP address
	nextHop		uint16 // IP address of next hop router
	input		uint16 // SNMP index of input interface
	output		uint16 // SNMP index of output interface
	dPkts		uint16 // Packets in the flow
	dOctets		uint16 // Total number of Layer 3 bytes in the packets of the flow
	first		uint16 // SysUptime at start of flow
	last		uint16 // SysUptime at the time the last packet of the flow was received
	srcPort		uint16 // TCP/UDP source port number or equivalent
	dstPort		uint16 // TCP/UDP destination port number or equivalent
	pad1		uint16 // Unused (zero) bytes
	tcpFlags	uint16 // Cumulative OR of TCP flags
	prot		uint16 // IP protocol type (for example, TCP = 6; UDP = 17)
	tos			uint16 // IP type of service (ToS)
	srcAs		uint16 // Autonomous system number of the source, either origin or peer
	dstAs		uint16 // Autonomous system number of the destination, either origin or peer
	srcMask		uint16 // Source address prefix mask bits
	dstMask		uint16 // Destination address prefix mask bits
	pad2		uint16 // 	Unused (zero) bytes
}
