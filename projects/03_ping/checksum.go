package main

// Checksum receives a byte slice and returns the uint16 ones's complement of
// the one's complement of the sum of all 16-bit words
func Checksum(data []byte) uint16 {
	// Append a byte of zeroes because we're dealing with uint16
	if len(data)%2 != 0 {
		data = append(data, 0x00)
	}

	// Convert byte slice into uint16 slice
	sixteen := make([]uint16, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		sixteen[i/2] = uint16(data[i])<<8 | uint16(data[i+1])
	}

	total := sixteen[0]
	// Get the 16-bit 1's complement
	for i := 1; i < len(sixteen); i++ {
		total = onesComplementAdd(total, sixteen[i])
	}

	// Get final 1's complement of result by XORing with 1s
	return uint16(^total)
}

func onesComplementAdd(x, y uint16) uint16 {
	sum32 := uint32(x) + uint32(y)
	return uint16(sum32) + uint16(sum32>>16)
}
