package util

func BitCount64(num uint64) int {

	var count int
	for 0 < num {
		num &= (num - 1)
		count++
	}
	return count
}

func BitCount32(num uint32) int {

	var count int
	for 0 < num {
		num &= (num - 1)
		count++
	}
	return count
}

func HasBit(num uint32, bit uint32) bool {
	return num&bit == bit
}

func SplitBits64(num uint64) []uint64 {
	var sets []uint64
	for i := uint32(0); i < 64; i++ {
		if f := uint64(0x01 << i); num&f > 0 {
			sets = append(sets, f)
		}
	}
	return sets
}
