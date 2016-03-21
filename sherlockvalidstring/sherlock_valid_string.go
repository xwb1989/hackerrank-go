package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.hackerrank.com/challenges/sherlock-and-valid-string
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100001)
	l, _, _ := reader.ReadLine()

	rf := make([]int, 26)
	for _, i := range l {
		rf[i-'a']++
	}

	a := -1
	ac := 0
	b := -1
	bc := 0
	fail := false
	for _, f := range rf {
		if f == 0 {
			continue
		}
		if a == -1 {
			a = f
			ac++
		} else if a == f {
			ac++
		} else {
			if b == -1 {
				b = f
				bc++
			} else if f == b {
				bc++
			} else {
				fail = true
				break
			}
		}
	}
	if fail {
		fmt.Println("NO")
	} else {
		if ac < bc {
			a, b = b, a
			ac, bc = bc, ac
		}
		if b == -1 {
			fmt.Println("YES")
		} else if b == 1 && bc == 1 {
			fmt.Println("YES")
		} else if (b-a) == 1 && bc == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
	fmt.Println()
}
