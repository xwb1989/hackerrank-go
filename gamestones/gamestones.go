package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	n, _ := strconv.Atoi(scan.Text())

	gem := map[rune]int{}
	for i := 0; i < n; i++ {
		scan.Scan()
		line := scan.Text()

		for _, r := range line {
			if gem[r] != i {
				continue
			} else {
				gem[r] = i + 1
			}
		}
	}
	cnt := 0
	for _, v := range gem {
		if v >= n {
			cnt++
		}
	}
	fmt.Println(cnt)
}
