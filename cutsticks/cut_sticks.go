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

	n := scan()

	sticks := make([]int, n)

	for i := range sticks {
		sticks[i] = scan()
	}

	cut := 0
	for {
		var cnt int
		cnt, cut = cutBy(sticks, cut)
		if cnt == 0 {
			break
		}
		fmt.Println(cnt)
	}
}

func cutBy(s []int, n int) (int, int) {
	cnt := 0
	min := math.MaxInt32
	for _, v := range s {
		if v > n {
			cnt++
			if v < min {
				min = v
			}
		}
	}
	return cnt, min
}
