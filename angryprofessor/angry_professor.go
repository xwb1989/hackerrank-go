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
	cnt := nextInt(scan)

	for i := 0; i < cnt; i++ {
		tot := nextInt(scan)
		thr := nextInt(scan)
		onTime := 0
		for j := 0; j < tot; j++ {
			t := nextInt(scan)
			if t <= 0 {
				onTime++
			}
		}
		if onTime < thr {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func nextInt(scan *bufio.Scanner) int {
	scan.Scan()
	i, _ := strconv.Atoi(scan.Text())
	return i
}
