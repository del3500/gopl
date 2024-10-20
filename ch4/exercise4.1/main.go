package popcount

//											 Efficient Bit Counting Techniques in Go
//
// Counting the number of 1s in the binary representation of a number can be done in multiple
// ways. A naive approach would involve checking each bit of the number one at a time, which can
// be inefficient, especially for larger numbers or when this approach needs to be performed frequently.
//

//			 									Precomputation for efficiency
//
// The pc array (var pc [256]byte) is used to store the population count for all possible values of
// an 8-bit byte (0 to 255).
// 	Efficieny:
//								The maximum value of an 8-bit byte is 255, which means there are only 256 possible
//						 		possible byte values. By precomputing the population counts for these values, you
//								can quickly retrieve the count of 1s in any byte by simply indexing into the pc array.
//								this reduces time complexity of counting bits significantly.
//
//	Constant Time Lookup:
// 								Accessing an element in an array is a constant-time operating (0(1))
//								meaning that regardless of the number of bytes, you can get the population count
//								for a byte almost instantly.

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Popcount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
