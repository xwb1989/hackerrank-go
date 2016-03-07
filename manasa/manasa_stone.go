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
		a := next()
		b := next()
		if a > b {
			a, b = b, a
		}
		steps := n - 1
		reported := map[int]bool{}
		for i := 0; i <= steps; i++ {
			v := b*i + a*(steps-i)
			if _, ok := reported[v]; !ok {
				fmt.Print(v, " ")
				reported[v] = true
			}
		}
		fmt.Println()
	}
}
