package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	scan.Scan()
	text := scan.Text()

	alpha := map[int]bool{}
	for _, c := range text {
		var offset int
		if c >= 'a' && c <= 'z' {
			offset = int(c - 'a')
		} else if c >= 'A' && c <= 'Z' {
			offset = int(c - 'A')
		} else {
			continue
		}
		alpha[offset] = true
		if len(alpha) == 26 {
			break
		}
	}
	if len(alpha) == 26 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
