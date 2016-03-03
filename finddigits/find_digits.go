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
		num := scan()
		j := num
		c := 0
		for j > 0 {
			r := j % 10
			j /= 10
			if r != 0 && num%r == 0 {
				c++
			}
		}
		fmt.Println(c)
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
