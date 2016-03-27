package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/insertionsort2

func main() {
	next := func() func() int {
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanWords)

		return func() int {
			scan.Scan()
			i, _ := strconv.Atoi(scan.Text())
			return i
		}
	}()

	l := next()

	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = next()
	}
	if l == 1 {
		return
	}

	for i := 1; i < l; i++ {
		picked := arr[i]
		for j := i - 1; j >= 0; j-- {
			curr := arr[j]
			if curr > picked {
				arr[j+1] = curr
				if j == 0 {
					arr[j] = picked
				}
			} else {
				arr[j+1] = picked
				break
			}
		}
		printSlice(arr)
	}
}
func printSlice(s []int) {
	for _, i := range s {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
