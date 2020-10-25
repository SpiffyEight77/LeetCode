package leetcode

func reverseBits(num uint32) uint32 {
	var res uint32
	pow := 31
	for num != 0 {
		res += (num & 1) << pow
		num >>= 1
		pow--
	}
	return res
}
