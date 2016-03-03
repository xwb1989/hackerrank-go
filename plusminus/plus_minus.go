package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	cnt, _ := strconv.Atoi(scan.Text())

	pos := 0
	neg := 0
	zer := 0

	for i := 0; i < cnt; i++ {
		scan.Scan()
		num, _ := strconv.Atoi(scan.Text())
		switch {
		case num > 0:
			pos++
		case num == 0:
			zer++
		case num < 0:
			neg++
		}
	}
	divide := func(a int, b int) float64 {
		return float64(a) / float64(b)
	}
	fmt.Printf("%f\n", divide(pos, cnt))
	fmt.Printf("%f\n", divide(neg, cnt))
	fmt.Printf("%f\n", divide(zer, cnt))
}
