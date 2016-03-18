package main

// https://www.hackerrank.com/challenges/common-child
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	scan.Scan()
	a := scan.Text()
	scan.Scan()
	b := scan.Text()

	mem := make([][]int, len(a))
	for i := range mem {
		mem[i] = make([]int, len(b))
	}

	for i := range mem {
		for j := range mem[i] {
			var prev int
			if i == 0 && j == 0 {
				prev = 0
			} else if i == 0 {
				prev = mem[0][j-1]
			} else if j == 0 {
				prev = mem[i-1][0]
			}

			if a[i] != b[j] {
				if i > 0 && j > 0 {
					if mem[i-1][j] > mem[i][j-1] {
						prev = mem[i-1][j]
					} else {
						prev = mem[i][j-1]
					}
				}
				mem[i][j] = prev
			} else {
				if i > 0 && j > 0 {
					prev = mem[i-1][j-1]
				}
				mem[i][j] = prev + 1
			}
		}
	}
	fmt.Println(mem[len(a)-1][len(b)-1])
}
