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
		if len(s)%2 == 0 {
			a := s[0 : len(s)/2]
			b := s[len(s)/2:]

			ar := map[rune]int{}
			br := map[rune]int{}

			for _, r := range a {
				ar[r]++
			}

			for _, r := range b {
				br[r]++
			}

			tot := 0
			for k, v := range ar {
				bc := br[k]
				if v > bc {
					tot += v - bc
				}
			}
			fmt.Println(tot)

		} else {
			fmt.Println(-1)
		}
	}
}
