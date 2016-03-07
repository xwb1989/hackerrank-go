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
	g := make([]string, n)
	for i := 0; i < n; i++ {
		g[i] = next()
	}

	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c := g[i][j]
			res[i][j] = c
			if i > 0 && i < n-1 && j > 0 && j < n-1 {
				u := g[i-1][j]
				r := g[i][j+1]
				d := g[i+1][j]
				l := g[i][j-1]
				if c > u && c > r && c > d && c > l {
					res[i][j] = 'X'
				}
			}
		}
	}
	for _, r := range res {
		fmt.Println(string(r))
	}
}
