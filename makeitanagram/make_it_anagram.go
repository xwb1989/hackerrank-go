package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100001)

	a, _ := reader.ReadString('\n')
	b, _ := reader.ReadString('\n')

	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)
	ar := map[rune]int{}
	br := map[rune]int{}

	for _, r := range a {
		ar[r]++
	}

	for _, r := range b {
		br[r]++
	}

	fmt.Println("ar", ar)
	fmt.Println("br", br)
	tot := 0
	for k, v := range ar {
		bc := br[k]
		if v > bc {
			tot += v - bc
		}
	}
	for k, v := range br {
		ac := ar[k]
		if v > ac {
			tot += v - ac
		}
	}
	fmt.Println(tot)
}
