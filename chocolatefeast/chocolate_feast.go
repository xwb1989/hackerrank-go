package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	next := func() func() int {
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanWords)

		return func() int {
			scan.Scan()
			i, _ := strconv.Atoi(scan.Text())
			return i
		}

	}()

	t := next()
	for ; t > 0; t-- {
		n := next()
		c := next()
		m := next()
		eat := 0
		w := 0
		for n >= c || w >= m {
			var i int
			n, i, w = exchange(n, c, w, m)
			eat += i
		}
		fmt.Println(eat)
	}
}

func exchange(n, c, w, m int) (int, int, int) {
	i := n / c
	nr := n % c
	i += w / m
	wr := w % m
	wr += i
	return nr, i, wr
}
