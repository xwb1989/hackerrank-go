package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100009)
	l, _, _ := reader.ReadLine()
	t, _ := strconv.Atoi(string(l))

	for ; t > 0; t-- {
		s, _, _ := reader.ReadLine()

		removed := -1
		for i := 0; i < len(s)/2; i++ {
			a := s[i]
			b := s[len(s)-1-i]
			if a == b {
				continue
			} else {
				if checker(s, i) {
					removed = i
				} else {
					removed = len(s) - 1 - i
				}
				break
			}
		}
		fmt.Println(removed)
	}
}

func checker(s []byte, skip int) bool {
	off := 0
	for i := 0; i < len(s)/2; i++ {
		if i == skip {
			off = 1
			continue
		}
		a := s[i]
		b := s[len(s)-1-i+off]
		if a == b {
			continue
		} else {
			return false
		}
	}
	return true
}
