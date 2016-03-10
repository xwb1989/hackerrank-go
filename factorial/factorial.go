package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	n, _ := strconv.ParseInt(scan.Text(), 10, 64)
	fmt.Println(fact(n).String())
}

func fact(n int64) *big.Int {
	res := big.NewInt(n)
	if n == 1 {
		return res
	} else {
		return res.Mul(res, fact(n-1))
	}
}
