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
	d0, m0, y0 := next(), next(), next()
	d1, m1, y1 := next(), next(), next()
	if y0 < y1 || (y0 == y1 && m0 < m1) || (y0 == y1 && m0 == m1 && d0 <= d1) {
		fmt.Println(0)
	} else if y0 == y1 && m0 == m1 {
		fmt.Println(15 * (d0 - d1))
	} else if y0 == y1 {
		fmt.Println(500 * (m0 - m1))
	} else {
		fmt.Println(10000)
	}
}
