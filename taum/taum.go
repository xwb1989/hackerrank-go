package main

import (
	"bufio"
	"fmt"
	"math"
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
		b, w, x, y, z := next(), next(), next(), next(), next()

		abs := int(math.Abs(float64(x - y)))
		if abs > z {
			if x > y {
				fmt.Println((b+w)*y + b*z)
			} else {
				fmt.Println((b+w)*x + w*z)
			}
		} else {
			fmt.Println(b*x + w*y)
		}
	}
}
