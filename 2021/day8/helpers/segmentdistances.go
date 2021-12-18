package helpers

import (
	"math/bits"
)

/*

Reference:

 aaaa
b    c
b    c
 dddd
e    f
e    f
 gggg

0: abc efg
1:   c f
2: a cde g
3: a cd fg
4:  bcd f
5: ab d fg
6: ab defg
7: a c  f
8: abcdefg
9: abcd fg

*/

func hamming(a, b uint) int {
	return bits.OnesCount(a ^ b)
}

func getDistances(center uint) []int {
	binarySegments := []uint{
		uint(0b1110111), // 0
		uint(0b0010100), // 1
		uint(0b1011101), // 2
		uint(0b1011011), // 3
		uint(0b0111010), // 4
		uint(0b1101011), // 5
		uint(0b1101111), // 6
		uint(0b1010010), // 7
		uint(0b1111111), // 8
		uint(0b1111011), // 9
	}
	distances := make([]int, len(binarySegments))
	for i, _ := range distances {
		distances[i] = hamming(binarySegments[center], binarySegments[i])
	}
	return distances
}

func GetAllDistances() [][]int{
	allDistances := make([][]int, 10)
	for i := 0; i < len(allDistances); i++ {
		allDistances[i] = getDistances(uint(i))
	}
	return allDistances
}
