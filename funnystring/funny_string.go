package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	scan.Scan()
	t, _ := strconv.Atoi(scan.Text())

	for ; t > 0; t-- {
		scan.Scan()
		s := scan.Text()
		funny := true
		for i := range s {
			j := len(s) - 1 - i
			if i == len(s)-1 {
				break
			} else {
				var a, b int
				if s[i] > s[i+1] {
					a = int(s[i] - s[i+1])
				} else {
					a = int(s[i+1] - s[i])
				}
				if s[j] > s[j-1] {
					b = int(s[j] - s[j-1])
				} else {
					b = int(s[j-1] - s[j])
				}
				if a == b {
					continue
				} else {
					funny = false
					break
				}
			}
		}
		if funny {
			fmt.Println("Funny")
		} else {
			fmt.Println("Not Funny")
		}
	}

}
