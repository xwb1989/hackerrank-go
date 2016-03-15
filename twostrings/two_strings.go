package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100001)
	l, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(l))

	for ; t > 0; t-- {
		a, _ := reader.ReadString('\n')
		b, _ := reader.ReadString('\n')
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)

		ar := map[rune]bool{}
		for _, r := range a {
			ar[r] = true
		}

		found := false
		for _, r := range b {
			if ar[r] {
				found = true
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
