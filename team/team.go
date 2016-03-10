package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	next := func() func() string {
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanWords)
		return func() string {
			scan.Scan()
			return scan.Text()
		}
	}()
	n, _ := strconv.Atoi(next())
	m, _ := strconv.Atoi(next())

	p := make([]string, n)
	for i := 0; i < n; i++ {
		p[i] = next()
	}

	max := 0
	cnt := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			bit := 0
			for k := 0; k < m; k++ {
				if p[i][k] == '1' || p[j][k] == '1' {
					bit++
				}
			}
			if bit > max {
				max = bit
				cnt = 1
			} else if bit == max {
				cnt++
			}
		}
	}
	fmt.Println(max)
	fmt.Println(cnt)
}
