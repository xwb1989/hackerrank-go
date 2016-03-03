package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := scanner()
	cnt := scan()

	for i := 0; i < cnt; i++ {
		n := scan()
		// 5a + 3b = n
		a := 0
		found := false
		for a <= n/5 && !found {
			b := 0
			for b <= n/3 {
				sum := a*5 + b*3
				if sum < n {
					b++
				} else if sum == n {
					fmt.Printf("%s%s\n", strings.Repeat("5", b*3), strings.Repeat("3", a*5))
					found = true
					break
				} else {
					break
				}
			}
			a++
		}
		if !found {
			fmt.Println(-1)
		}
	}
}

func scanner() func() int {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	return func() int {
		scan.Scan()
		i, _ := strconv.Atoi(scan.Text())
		return i
	}
}
