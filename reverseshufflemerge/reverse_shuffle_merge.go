package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	s := scan.Text()

	rf := make([]int, 26)
	for _, r := range s {
		rf[r-'a']++
	}

	for i, f := range rf {
		rf[i] = f / 2
	}

	var found string = string(greedy(rf, []byte(s)))
	fmt.Println(found)
}

func leftMost(rf []int) int {
	for i := 0; i < len(rf); i++ {
		if rf[i] > 0 {
			return i
		}
	}
	return 25
}

func greedy(rf []int, origin []byte) []byte {
	res := make([]byte, len(origin)/2)
	canSkip := make([]int, 26)
	copy(canSkip, rf)
	pos := 0

	var best byte = 'z' + 1
	var bestIdx = -1
	for i := len(origin) - 1; i >= 0; i-- {
		minIdx := leftMost(rf)
		min := byte(minIdx) + 'a'
		curr := origin[i]
		currIdx := int(curr - 'a')
		var b byte
		var bIdx int
		if rf[currIdx] <= 0 {
			continue
		}
		if curr > min {
			if canSkip[currIdx] == 0 {
				if curr >= best {
					for i < bestIdx {
						i++
						curr = origin[i]
						currIdx = int(curr - 'a')
						canSkip[currIdx]++
					}
				}
				b = curr
				bIdx = currIdx
			} else {
				canSkip[currIdx]--
				if curr < best {
					best = curr
					bestIdx = i
				}
				continue
			}
		} else {
			b = curr
			bIdx = currIdx
		}
		rf[bIdx]--
		res[pos] = b
		pos++
		best = 'z' + 1
		if pos == len(res) {
			break
		}
	}
	return res
}
