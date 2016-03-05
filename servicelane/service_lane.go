package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := func() func() int {
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanWords)
		return func() int {
			scan.Scan()
			i, _ := strconv.Atoi(scan.Text())
			return i
		}
	}()

	n := scan()
	t := scan()

	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = scan()
	}
	for ; t > 0; t-- {
		a := scan()
		b := scan()
		max := 3
		for ; a <= b; a++ {
			if l[a] < max {
				max = l[a]
			}
			if max == 1 {
				break
			}
		}
		fmt.Println(max)
	}

}
