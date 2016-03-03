package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	cnt, _ := strconv.Atoi(scan.Text())
	s := cnt - 1
	for i := 0; i < cnt; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", s), strings.Repeat("#", cnt-s))
		s--
	}
}
