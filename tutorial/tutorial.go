package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nextLine := func() func() []int {
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanLines)
		return func() []int {
			scan.Scan()
			a := strings.Split(scan.Text(), " ")
			res := make([]int, len(a))
			for idx, e := range a {
				i, _ := strconv.Atoi(e)
				res[idx] = i
			}
			return res
		}
	}()

	v := nextLine()[0]
	nextLine()
	a := nextLine()

	for idx, e := range a {
		if v == e {
			fmt.Println(idx)
			break
		}
	}

}
