package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	next := func() func() string {
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanLines)

		return func() string {
			scan.Scan()
			return scan.Text()
		}
	}()

	next()

	s := []byte(next())
	ki, _ := strconv.Atoi(next())
	k := byte(ki % 26)

	for i, r := range s {
		if r >= 'a' && r <= 'z' {
			r = r + k
			if r > 'z' {
				r -= byte(26)
			}
		} else if r >= 'A' && r <= 'Z' {
			r = r + k
			if r > 'Z' {
				r -= byte(26)
			}
		}
		s[i] = r
	}
	fmt.Println(string(s))
}
