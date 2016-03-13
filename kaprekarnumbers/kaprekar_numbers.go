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
	p, q := next(), next()

	found := false
	for ; p <= q; p++ {
		square := p * p
		d := 1
		for i := p; i > 0; i /= 10 {
			d *= 10
		}
		r := square % d
		l := square / d
		if r+l == p {
			fmt.Print(p, " ")
			found = true
		}
	}
	if !found {
		fmt.Println("INVALID RANGE")
	} else {
		fmt.Println()
	}
}
