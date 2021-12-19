package helpers

import (
	"log"
	"math/bits"
	"sort"
	"strings"
	"strconv"
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
1:   c  f
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
	//fmt.Printf("%b %b\n", a, b)
	return bits.OnesCount(a ^ b)
}

func getDistances(center uint) []int {
	binarySegments := []uint{
		uint(0b1110111), // 0
		uint(0b0010010), // 1
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
	sort.Ints(distances)
	return distances
}

func GetAllDistances() [][]int {
	allDistances := make([][]int, 10)
	for i := 0; i < len(allDistances); i++ {
		allDistances[i] = getDistances(uint(i))
	}
	return allDistances
}

func alphaToBinary(s string) uint {
	possibleChars := []string{"a", "b", "c", "d", "e", "f", "g"}
	binaryString := ""
	for _, char := range possibleChars {
		if strings.Contains(s, char) {
			binaryString += "1"
		} else {
			binaryString += "0"
		}
	}
	binary, err := strconv.ParseUint(binaryString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return uint(binary)
}

func alphaToBinarySlice(segments []string) []uint {
	binaries := make([]uint, len(segments))
	for i, _ := range binaries{
		binaries[i] = alphaToBinary(segments[i])
	}
	return binaries
}

func GetAlphaDistances(center string, segments []string) []int {
	distances := make([]int, len(segments))
	centerBinary := alphaToBinary(center)
	segmentBinaries := alphaToBinarySlice(segments)
	for i, segmentBinary := range segmentBinaries{
		distances[i] = hamming(centerBinary, segmentBinary)
	}
	sort.Ints(distances)
	return distances
}

func GetAllAlphaDistances(segments []string) [][]int {
	allDistances := make([][]int, len(segments))
	for i := 0; i < len(allDistances); i++ {
		allDistances[i] = GetAlphaDistances(segments[i], segments)
	}
	return allDistances
}
