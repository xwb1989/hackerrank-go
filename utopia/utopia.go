package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := scanner()
	cnt := scan()
	for i := 0; i < cnt; i++ {
		cyc := scan()
		var height int64 = 1
		for j := 0; j < cyc; j++ {
			if j%2 == 0 {
				height *= 2
			} else {
				height++
			}
		}
		fmt.Println(height)

	}
}

func scanner() func() int {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	return func() int {
		scan.Scan()
		i, _ := strconv.Atoi(scan.Text())
		return i
	}
}
