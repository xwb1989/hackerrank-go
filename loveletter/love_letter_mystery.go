package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	t, _ := strconv.Atoi(scan.Text())
	for ; t > 0; t-- {
		scan.Scan()
		s := scan.Text()

		tot := 0
		for i := 0; i < len(s)/2; i++ {
			j := len(s) - 1 - i
			a := s[i]
			b := s[j]
			if a > b {
				tot += int(a - b)
			} else {
				tot += int(b - a)
			}
		}
		fmt.Println(tot)

	}
}
