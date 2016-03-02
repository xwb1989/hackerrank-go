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
	primarySum := 0
	secondarySum := 0
	level := 0
	i := 0
	j := 0
	for j < cnt*cnt {
		scan.Scan()
		num, _ := strconv.Atoi(scan.Text())
		if i == level {
			primarySum += num
		}
		if i == cnt-1-level {
			secondarySum += num
		}
		i++
		if i == cnt {
			level++
			i = 0
		}
		j++
	}
	var res int
	if primarySum > secondarySum {
		res = primarySum - secondarySum
	} else {
		res = secondarySum - primarySum
	}
	fmt.Println(res)
}
