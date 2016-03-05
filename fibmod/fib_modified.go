package main

import (
	"bufio"
	"fmt"
	"math/big"
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

	a := big.NewInt(int64(scan()))
	b := big.NewInt(int64(scan()))
	n := scan()

	for ; n >= 3; n-- {
		c := big.NewInt(0)
		a.Add(c.Mul(b, b), a)
		a, b = b, a
	}
	fmt.Println(b)
}
