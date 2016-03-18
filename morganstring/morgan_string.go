package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/morgan-and-a-string
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100005)
	l, _, _ := reader.ReadLine()
	t, _ := strconv.Atoi(string(l))

	for ; t > 0; t-- {
		ab, _, _ := reader.ReadLine()
		a := string(ab)
		bb, _, _ := reader.ReadLine()
		b := string(bb)
		var i, j, k int
		var pickA bool
		var whenSmaller byte
		res := make([]byte, len(a)+len(b))
		for i < len(a) && j < len(b) {
			var pick byte
			if a[i] < b[j] {
				pick = a[i]
				i++
			} else if a[i] > b[j] {
				pick = b[j]
				j++
			} else {
				if a[i] >= whenSmaller {
					ii, jj := i, j
					for ii < len(a) && jj < len(b) && a[ii] == b[jj] {
						ii++
						jj++
					}
					if ii == len(a) && jj == len(b) {
						pickA = true
						whenSmaller = a[ii-1]
					} else if ii == len(a) {
						pickA = false
						whenSmaller = b[jj]
					} else if jj == len(b) {
						pickA = true
						whenSmaller = a[ii]
					} else {
						pickA = a[ii] < b[jj]
						if pickA {
							whenSmaller = a[ii]
						} else {
							whenSmaller = b[jj]
						}
					}
				}
				if pickA {
					pick = a[i]
					i++
				} else {
					pick = b[j]
					j++
				}
				if pick >= whenSmaller {
					whenSmaller = 'A' - 1
				}
			}
			res[k] = pick
			k++
		}
		if i < len(a) {
			copy(res[k:], a[i:])
		} else if j < len(b) {
			copy(res[k:], b[j:])
		}
		fmt.Println(string(res))
	}
}
