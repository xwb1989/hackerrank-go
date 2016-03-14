package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100001)
	text, _ := reader.ReadString('\n')

	seen := map[rune]bool{}
	for _, r := range text {
		if seen[r] {
			delete(seen, r)
		} else {
			seen[r] = true
		}
	}
	if len(seen) <= 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
