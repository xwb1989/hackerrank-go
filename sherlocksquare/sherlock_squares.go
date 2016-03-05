package main

import (
	"bufio"
	"fmt"
	"math"
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

	t := scan()
	for ; t > 0; t-- {
		a := math.Sqrt(float64(scan()))
		b := math.Sqrt(float64(scan()))
		cnt := int(b) - int(math.Ceil(a)) + 1
		fmt.Println(cnt)
	}
}
