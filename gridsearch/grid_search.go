package main

import (
	"bufio"
	"bytes"
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

	t, _ := strconv.Atoi(next())
	for ; t > 0; t-- {
		r, _ := strconv.Atoi(next())
		c, _ := strconv.Atoi(next())
		g := make([][]byte, r)
		for i := 0; i < r; i++ {
			g[i] = []byte(next())
		}

		pr, _ := strconv.Atoi(next())
		pc, _ := strconv.Atoi(next())
		p := make([][]byte, pr)
		for i := 0; i < pr; i++ {
			p[i] = []byte(next())
		}

		found := false
		for rn := 0; rn < r-pr+1; rn++ {
			for cn := 0; cn < c-pc+1; cn++ {
				if search(g, rn, cn, p) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func search(g [][]byte, r int, c int, p [][]byte) bool {
	for i := 0; i < len(p); i++ {
		row := g[r+i]
		val := row[c : c+len(p[i])]
		if !bytes.Equal(val, p[i]) {
			return false
		}
	}
	return true
}
