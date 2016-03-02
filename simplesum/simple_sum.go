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
	sum := 0
	for i := 0; i < cnt; i++ {
		scan.Scan()
		j, _ := strconv.Atoi(scan.Text())
		sum += j
	}
	fmt.Println(sum)
}
